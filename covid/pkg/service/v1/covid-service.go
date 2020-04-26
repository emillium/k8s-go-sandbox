package v1

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/emillium/k8s-go-sandbox/covid/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type covidServiceServer struct{}

// NewToDoServiceServer creates ToDo service
func NewCovidServiceServer() v1.CovidServiceServer {
	return &covidServiceServer{}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *covidServiceServer) checkAPI(api string) error {
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
func (s *covidServiceServer) Add(ctx context.Context, req *v1.AddRequest) (*v1.AddResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Do some database shit
	result := &v1.AddResponse{
		Api:    apiVersion,
		Result: req.Covid.A + req.Covid.B,
	}

	// Log result
	logMessage := fmt.Sprintf("A: %d   B: %d     sum: %d", req.Covid.A, req.Covid.B, result.Result)
	log.Println(logMessage)

	return result, nil
}
