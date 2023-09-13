// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/qclaogui/golang-api-server/api/todo/v1/todopb"
	"github.com/qclaogui/golang-api-server/cmd/client-grpc/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:9095", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() { _ = conn.Close() }()

	client := pb.NewToDoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Call createTodo
	id := createTodo(ctx, client)

	// Call getTodo
	todo := getTodo(ctx, client, id)

	// Call updateTodo
	updateTodo(ctx, client, todo)

	// Call listTodo
	listTodo(ctx, client)

	// Call deleteTodo
	deleteTodo(ctx, client, id)

}

func createTodo(ctx context.Context, client pb.ToDoServiceClient) string {

	resp, err := client.Create(ctx, &pb.CreateRequest{
		Api: apiVersion,
		Item: &pb.ToDo{
			Title:       "title",
			Description: "description",
			CreatedAt:   timestamppb.New(time.Now().UTC().Add(time.Minute)),
		},
	})
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}

	log.Printf("Create resp: <%v>\n\n", resp)
	return resp.Id
}

func getTodo(ctx context.Context, client pb.ToDoServiceClient, id string) *pb.ToDo {

	resp, err := client.Get(ctx, &pb.GetRequest{
		Api: apiVersion,
		Id:  id,
	})
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}
	log.Printf("Get resp: <%v>\n\n", resp)
	return resp.Item

}

func updateTodo(ctx context.Context, client pb.ToDoServiceClient, todo *pb.ToDo) {

	resp, err := client.Update(ctx, &pb.UpdateRequest{
		Api: apiVersion,
		Item: &pb.ToDo{
			Id:          todo.Id,
			Title:       todo.Title,
			Description: todo.Description + " updated",
			CreatedAt:   todo.CreatedAt,
		},
	})
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update resp: <%v>\n\n", resp)
}

func listTodo(ctx context.Context, client pb.ToDoServiceClient) {

	resp, err := client.List(ctx, &pb.ListRequest{
		Api: apiVersion,
	})
	if err != nil {
		log.Fatalf("List failed: %v", err)
	}
	log.Printf("List resp: <%v>\n\n", resp)
}

func deleteTodo(ctx context.Context, client pb.ToDoServiceClient, id string) {

	resp, err := client.Delete(ctx, &pb.DeleteRequest{
		Api: apiVersion,
		Id:  id,
	})
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete resp: <%v>\n\n", resp)
}
