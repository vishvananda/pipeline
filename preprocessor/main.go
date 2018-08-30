package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/bamiaux/rez"
	gp "github.com/oracle/graphpipe-go"
)

var (
	client *http.Client
)

const (
	nextModelUri = "http://127.0.0.1:9000"
	defaultPort  = "4000"
)

func main() {
	port := defaultPort
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	client = newClient()

	if err := gp.Serve("0.0.0.0:"+port, false, apply, nil, nil); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func apply(_requestContext *gp.RequestContext, _config string, images []string) ([][]float32, error) {
	request, err := buildRequest(images)
	if err != nil {
		return nil, err
	}

	req := []interface{}{request}
	results, err := gp.MultiRemote(client, nextModelUri, "", req, nil, nil)
	if err != nil {
		return nil, err
	}
	return results[0].([][]float32), nil
}

func newClient() *http.Client {
	// set some timeouts for the client
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		IdleConnTimeout:     5 * time.Second,
		TLSHandshakeTimeout: 5 * time.Second,
		MaxIdleConns:        1000,
		MaxIdleConnsPerHost: 200,
		Proxy:               http.ProxyFromEnvironment,
	}
	return &http.Client{
		Timeout:   time.Second * 120,
		Transport: transport,
	}
}

func decode(jpg string) ([][][]float32, error) {
	// decode jpeg into image.Image
	img, err := jpeg.Decode(strings.NewReader(jpg))
	if err != nil {
		return nil, err
	}

	m := image.NewYCbCr(image.Rect(0, 0, 224, 224), 0)
	if err := rez.Convert(m, img, rez.NewBicubicFilter()); err != nil {
		return nil, err
	}

	result := make([][][]float32, 224)
	for row := range result {
		result[row] = make([][]float32, 224)
		for col := range result[row] {
			c := m.At(col, row).(color.YCbCr)
			r, g, b := color.YCbCrToRGB(c.Y, c.Cb, c.Cr)
			// rgb -> bgr and subtract mean values
			result[row][col] = []float32{float32(b) - 103.939, float32(g) - 116.779, float32(r) - 123.68}
		}
	}
	return result, nil
}

type ErrorResult struct {
	index int
	err   error
}

// buildRequest builds the request in parallel
func buildRequest(images []string) ([][][][]float32, error) {
	n := len(images)
	req := make([][][][]float32, n)

	errChan := make(chan ErrorResult, n)
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(n)
	for i, jpg := range images {
		go func(i int, jpg string) {
			defer wg.Done()
			var err error
			req[i], err = decode(jpg)
			if err != nil {
				errChan <- ErrorResult{i, err}
			}
		}(i, jpg)
	}
	wg.Wait()
	select {
	case res, ok := <-errChan:
		if ok {
			// just return the first error
			return nil, fmt.Errorf("Failed to process image at index %d: %v", res.index, res.err)
		}
		// channel closed
	default:
		// no error
	}
	return req, nil
}
