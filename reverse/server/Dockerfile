FROM golang
ADD . /go/src/github.com/emillium/k8s-go-sandbox/reverse/server
RUN go install github.com/emillium/k8s-go-sandbox/reverse/server
ENTRYPOINT ["/go/bin/server"]
EXPOSE 5300
