// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.
//go:build requires_docker

package routeguide_test

import (
	"context"
	"flag"
	"log"
	"os"
	"testing"

	routeguide "github.com/qclaogui/gaip/genproto/routeguide/apiv1"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestMain(m *testing.M) {
	flag.Parse()

	conn, err := grpc.Dial("localhost:9095", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = conn.Close() }()

	opt := option.WithGRPCConn(conn)
	ctx := context.Background()

	routeguideGRPC, err = routeguide.NewClient(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = routeguideGRPC.Close() }()

	os.Exit(m.Run())
}
