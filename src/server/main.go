package main

import (
	"api"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, in *api.Request) (*api.Response, error) {
	return &api.Response{ResponseApi: "Welcome to gRPC " + in.RequestApi}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":8080")

	c := grpc.NewServer()

	api.RegisterGreetingServer(c, &Server{})

	reflection.Register(c)

	fmt.Println("Server started at port :8080")
	c.Serve(lis)
}
