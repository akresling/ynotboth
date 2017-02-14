package service

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/akresling/ynotboth/pb"
)

// Example is an implementation of the ExampleServer interface definition
type Example struct{}

// Hello will take a name and favourite colour and return a greeting
func (h Example) Hello(ctx context.Context, hr *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Greeting: fmt.Sprintf("Hello %s I like the colour %s too!", hr.GetName(), hr.GetColor())}, nil
}
