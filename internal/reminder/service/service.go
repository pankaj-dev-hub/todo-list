package service

import (
	"context"
	pb "pankaj-katyare/todo-list/internal/reminder/grpc"
)

type TodoServer struct {
	pb.UnimplementedTodoReminderServer
}

// We implement the CreateReminder method of the server interface.
func (s *TodoServer) CreateReminder(ctx context.Context, in *pb.CreateReminderRequest) (*pb.CreateReminderResponse, error) {
	return &pb.CreateReminderResponse{Status: "Completed", Result: true}, nil
}

// We implement the GetReminder method of the server interface.
func (s *TodoServer) GetReminder(ctx context.Context, in *pb.GetReminderRequest) (*pb.GetReminderResponse, error) {
	return &pb.GetReminderResponse{Status: "true"}, nil
}
