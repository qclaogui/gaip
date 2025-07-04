// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package a2a_test

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/qclaogui/gaip/sdk/a2a-go/a2a"
	"google.golang.org/api/iterator"
)

func ExampleAgentCards_GetAgentCard() {
	ctx := context.Background()
	client, err := a2a.NewClient(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.AgentCards.GetAgentCard(ctx)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Println(resp.Name)
}

func ExampleMessages_SendMessage_textOnly() {
	ctx := context.Background()
	client, err := a2a.NewClient(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Construct the user message contents.
	message := a2a.NewUserMessageFromText("Hello, agent!")

	config := &a2a.SendMessageConfiguration{
		Blocking:            true,
		AcceptedOutputModes: []string{"text/plain"},
	}

	// Call the SendMessage method with a text-only message
	resp, err := client.Messages.SendMessage(ctx, message, config)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	// Print the response
	printResponse(resp)
}

func ExampleMessages_SendStreamingMessage() {
	ctx := context.Background()
	client, err := a2a.NewClient(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Construct the user message contents with an image.
	message := a2a.NewUserMessageFromText("Check out this image!")
	// message.AddImageURL("https://example.com/image.jpg")

	// Call the SendMessage method with a message containing an image
	iter := client.Messages.SendStreamingMessage(ctx, message, nil)
	for {
		resp, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			log.Fatalf("Failed to get next response: %v", err)
		}
		// Print the response
		printStreamResponse(resp)
	}
}

// printStreamResponse Helping for printing the response.
func printStreamResponse(resp *a2a.StreamResponse) {
	msg := resp.Payload.Msg
	if msg != nil && msg.Content != nil {
		for _, part := range msg.Content {
			fmt.Println(part)
		}
	}
	fmt.Println("---")
}

// printResponse Helping for printing the response. the
// result can be a Task or a Message. Check which one it is.
func printResponse(resp *a2a.SendMessageResponse) {
	// The agent created a task.
	if task := resp.Payload.Task; task != nil {
		fmt.Printf("Send Message Result (Task): %v", task)
		// Save the task ID for the next call
	}

	// The agent responded with a direct message.
	if msg := resp.Payload.Msg; msg != nil && msg.Content != nil {
		for _, part := range msg.Content {
			fmt.Printf("Send Message Result (Direct Message): %v", part)
			// No task was created, so we can't get task status.
		}
	}

	fmt.Println("---")
}
