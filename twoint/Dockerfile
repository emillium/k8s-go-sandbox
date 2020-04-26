FROM golang
ADD ./cmd/api /go/src/github.com/emillium/k8s-go-sandbox/two-int/cmd/api
ADD ./pkg/api/v1 /go/src/github.com/emillium/k8s-go-sandbox/twoint/pkg/api/v1
WORKDIR /go/src/github.com/emillium/k8s-go-sandbox/two-int/cmd/api
RUN go get .
ENTRYPOINT go run main.go
EXPOSE 5300
