FROM golang

COPY . /go/src/service
WORKDIR /go/src/service
RUN go get .
ENTRYPOINT go run main.go
EXPOSE 3000
