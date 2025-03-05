// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.
//go:build requires_docker

package generativeai_test

import (
	"context"
	"flag"
	"log"
	"os"
	"testing"

	generativelanguage "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1"
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

	conn, err := grpc.NewClient("localhost:9095", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = conn.Close() }()

	opt := option.WithGRPCConn(conn)
	ctx := context.Background()

	generativelanguageGRPC, err = generativelanguage.NewGenerativeClient(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = generativelanguageGRPC.Close() }()

	generativelanguageREST, err = generativelanguage.NewGenerativeRESTClient(ctx, restClientOpts...)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = generativelanguageREST.Close() }()

	modelGRPC, err = generativelanguage.NewModelClient(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = modelGRPC.Close() }()

	modelREST, err = generativelanguage.NewModelRESTClient(ctx, restClientOpts...)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = modelREST.Close() }()

	os.Exit(m.Run())
}
