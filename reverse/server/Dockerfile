FROM golang
ADD . /go/src/github.com/emillium/k8s-go-sandbox/reverse/server
WORKDIR /go/src/github.com/emillium/k8s-go-sandbox/reverse/server
RUN go get .
RUN go install server.go
ENTRYPOINT ["/go/bin/server"]
EXPOSE 5300
