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
	message := a2a.NewUserMessageFromText("Good morning! How are you?")

	// Call the SendMessage method with a text-only message
	resp, err := client.Messages.SendMessage(ctx, message, nil)
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

// printResponse Helping for printing the response.
func printResponse(resp *a2a.SendMessageResponse) {
	msg := resp.Payload.Msg
	if msg != nil && msg.Content != nil {
		for _, part := range msg.Content {
			fmt.Println(part)
		}
	}
	fmt.Println("---")
}
