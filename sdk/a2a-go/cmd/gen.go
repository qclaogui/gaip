// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"google.golang.org/api/option"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/qclaogui/gaip/sdk/a2a-go/a2a"
)

func main() {
	ctx := context.Background()

	client, err := a2a.NewClient(ctx, nil,
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Construct the user message contents.
	message := a2a.NewUserMessageFromText("Good morning! How are you?")

	// Call the SendMessage method with a text-only message
	resp, err := client.Messages.SendMessage(ctx, message, nil)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	fmt.Println(resp)
}
