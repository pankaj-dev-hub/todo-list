package main

import (
	"log"
	"net"

	pb "pankaj-dev-hub/todo-list/internal/reminder/grpc"
	"pankaj-dev-hub/todo-list/internal/reminder/service"

	"google.golang.org/grpc"
)

func main() {

	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	println("gRPC server started....")

	s := grpc.NewServer()
	println("called new server started....")
	pb.RegisterTodoReminderServer(s, &service.TodoServer{})
	println("called register server started....")

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
