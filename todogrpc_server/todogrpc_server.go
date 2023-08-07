package main

import (
	"context"
	"log"
	"net"

	"bxcodec-clean-arch/domain"
	pb "bxcodec-clean-arch/todogrpc"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type TodoGrpcServer struct {
	pb.UnimplementedTodoCRUDServer
	todoUsecase domain.TodoUsecase
}

func (s *TodoGrpcServer) CreateTodo(ctx context.Context, in *pb.NewTodo) (*pb.Error, error) {
	log.Printf("Received: %v", in.GetName())
	todo := domain.Todo{
		Name:   in.GetName(),
		Status: in.GetStatus(),
	}
	err := s.todoUsecase.CreateTodo(&todo)
	if err == nil {
		return &pb.Error{Err: "No error"}, nil
	} else {
		log.Println(err)
		return &pb.Error{Err: err.Error()}, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodoCRUDServer(s, &TodoGrpcServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
