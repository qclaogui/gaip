// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.
//go:build requires_docker

package todo_test

import (
	"context"
	"flag"
	"log"
	"os"
	"testing"

	todo "github.com/qclaogui/gaip/genproto/todo/apiv1"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var restClientOpts = []option.ClientOption{
	option.WithEndpoint("http://localhost:7469"),
	option.WithoutAuthentication(),
}

func TestMain(m *testing.M) {
	flag.Parse()

	conn, err := grpc.Dial("localhost:9095", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = conn.Close() }()

	opt := option.WithGRPCConn(conn)
	ctx := context.Background()

	todoGRPC, err = todo.NewClient(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = todoGRPC.Close() }()

	todoREST, err = todo.NewClient(ctx, restClientOpts...)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = todoREST.Close() }()

	os.Exit(m.Run())
}
