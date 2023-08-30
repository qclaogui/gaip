// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package main

import (
	"context"
	"flag"
	"log"
	"time"

	todopbv1 "github.com/qclaogui/golang-api-server/api/gen/proto/todo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {
	// get configuration
	address := flag.String("server", "localhost:9095", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() { _ = conn.Close() }()

	c := todopbv1.NewToDoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reminder := timestamppb.New(time.Now().UTC().Add(time.Minute))

	// Create
	req1 := todopbv1.CreateRequest{
		Api: apiVersion,
		ToDo: &todopbv1.ToDo{
			Title:       "title",
			Description: "description",
			Reminder:    reminder,
		},
	}

	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}

	log.Printf("Create result: <%+v>\n\n", res1)

	//	Read
	id := res1.Id
	req2 := todopbv1.ReadRequest{Api: apiVersion, Id: id}

	res2, err := c.Read(ctx, &req2)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Printf("Read result: <%+v>\n\n", res2)

	// Update
	req3 := todopbv1.UpdateRequest{
		Api: apiVersion,
		ToDo: &todopbv1.ToDo{
			Id:          res2.ToDo.Id,
			Title:       res2.ToDo.Title,
			Description: res2.ToDo.Description + " + updated",
			Reminder:    res2.ToDo.Reminder,
		},
	}
	res3, err := c.Update(ctx, &req3)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update result: <%+v>\n\n", res3)

	// ReadAll
	req4 := todopbv1.ReadAllRequest{
		Api: apiVersion,
	}
	res4, err := c.ReadAll(ctx, &req4)
	if err != nil {
		log.Fatalf("ReadAll failed: %v", err)
	}
	log.Printf("ReadAll result: <%+v>\n\n", res4)

	// Delete
	req5 := todopbv1.DeleteRequest{
		Api: apiVersion,
		Id:  id,
	}
	res5, err := c.Delete(ctx, &req5)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", res5)

}
