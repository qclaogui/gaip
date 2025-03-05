// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.
//go:build requires_docker

package showcase_test

import (
	"context"
	"flag"
	"log"
	"os"
	"testing"

	showcase "github.com/qclaogui/gaip/genproto/showcase/apiv1beta1"
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

	echoGRPC, err = showcase.NewEchoClient(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = echoGRPC.Close() }()

	echoREST, err = showcase.NewEchoRESTClient(ctx, restClientOpts...)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = echoREST.Close() }()

	identityGRPC, err = showcase.NewIdentityClient(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = identityGRPC.Close() }()

	identityREST, err = showcase.NewIdentityRESTClient(ctx, restClientOpts...)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = identityREST.Close() }()

	messagingGRPC, err = showcase.NewMessagingClient(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = messagingGRPC.Close() }()

	messagingREST, err = showcase.NewMessagingRESTClient(ctx, restClientOpts...)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = messagingREST.Close() }()

	os.Exit(m.Run())
}
