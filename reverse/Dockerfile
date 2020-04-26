FROM golang
ADD ./api /go/src/github.com/emillium/k8s-go-sandbox/reverse/api
ADD ./server /go/src/github.com/emillium/k8s-go-sandbox/reverse/server
WORKDIR /go/src/github.com/emillium/k8s-go-sandbox/reverse/server
RUN go get .
ENTRYPOINT go run server.go
EXPOSE 5300
