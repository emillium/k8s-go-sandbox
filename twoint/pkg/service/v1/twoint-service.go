package v1

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/emillium/k8s-go-sandbox/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type twoIntServiceServer struct{}

// NewToDoServiceServer creates ToDo service
func NewTwoIntServiceServer() v1.TwoIntServiceServer {
	return &twoIntServiceServer{}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *twoIntServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// Create new todo task
func (s *twoIntServiceServer) Add(ctx context.Context, req *v1.AddRequest) (*v1.AddResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Do some database shit
	result := &v1.AddResponse{
		Api:    apiVersion,
		Result: req.TwoInt.A + req.TwoInt.B,
	}

	// Log result
	logMessage := fmt.Sprintf("A: %d   B: %d     sum: %d", req.TwoInt.A, req.TwoInt.B, result.Result)
	log.Println(logMessage)

	return result, nil
}
