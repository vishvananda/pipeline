from sleepsonthefloor/graphpipe-tf:cpu

ADD squeezenet.pb /squeezenet.pb

ENTRYPOINT ["/graphpipe-tf", "-m", "/squeezenet.pb"]
