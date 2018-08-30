package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	gp "github.com/oracle/graphpipe-go"
)

var (
	client *http.Client
)

const (
	defaultPort = "3000"
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

func apply(_requestContext *gp.RequestContext, config string, images []string) ([][]float32, error) {
	request, err := buildRequest(images)
	if err != nil {
		return nil, err
	}

	req := []interface{}{request}
	results, err := gp.MultiRemote(client, config, "", req, nil, nil)
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

func download(url string) (string, error) {
	// decode jpeg into image.Image
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	result := string(b)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Url returned status code %d: %s", resp.StatusCode, result)
	}
	return result, nil
}

type ErrorResult struct {
	index int
	err   error
}

// buildRequest builds the request in parallel
func buildRequest(urls []string) ([]string, error) {
	n := len(urls)
	req := make([]string, n)

	errChan := make(chan ErrorResult, n)
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(n)
	for i, url := range urls {
		go func(i int, url string) {
			defer wg.Done()
			var err error
			req[i], err = download(url)
			if err != nil {
				errChan <- ErrorResult{i, err}
			}
		}(i, url)
	}
	wg.Wait()
	select {
	case res, ok := <-errChan:
		if ok {
			// just return the first error
			return nil, fmt.Errorf("Failed to process url at index %d: %v", res.index, res.err)
		}
		// channel closed
	default:
		// no error
	}
	return req, nil
}
