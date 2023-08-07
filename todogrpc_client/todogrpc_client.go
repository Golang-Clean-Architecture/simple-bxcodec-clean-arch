package main

import (
	"context"
	"log"
	"time"

	pb "bxcodec-clean-arch/todogrpc"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTodoCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var new_users = make(map[string]string)
	new_users["Task 2"] = "Not Done"
	for name, status := range new_users {
		_, err := c.CreateTodo(ctx, &pb.NewTodo{Name: name, Status: status})
		if err != nil {
			log.Fatalf("could not create todo: %v", err)
		}
	}
}
