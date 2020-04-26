package cmd

import (
	"context"

	"github.com/emillium/k8s-go-sandbox/covid/pkg/protocol/grpc"
	v1 "github.com/emillium/k8s-go-sandbox/covid/pkg/service/v1"
)

// Config is configuration for Server
type Config struct{}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	v1API := v1.NewCovidServiceServer()

	return grpc.RunServer(ctx, v1API, "3000")
}
