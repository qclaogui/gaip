// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

import (
	"context"

	pb "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta/generativelanguagepb"
)

// Chats provides util functions for creating a new chat session.
// You don't need to initiate this struct. Create a client instance via NewClient, and
// then access Chats through client.Models field.
type Chats struct {
	apiClient *apiClient
}

// Chat represents a single chat session (multi-turn conversation) with the model.
//
//		client, _ := genai.NewClient(ctx, &genai.ClientConfig{})
//		chat, _ := client.Chats.Create(ctx, "gemini-2.0-flash", nil, nil)
//	  result, err = chat.SendMessage(ctx, genai.Part{Text: "What is 1 + 2?"})
type Chat struct {
	Models
	apiClient *apiClient
	model     string
	config    *pb.GenerationConfig

	// History of the chat.
	comprehensiveHistory []*pb.Content
}

// Create initializes a new chat session.
func (c *Chats) Create(_ context.Context, model string, config *pb.GenerationConfig, history []*pb.Content) (*Chat, error) {
	chat := &Chat{
		apiClient:            c.apiClient,
		model:                model,
		config:               config,
		comprehensiveHistory: history,
	}
	chat.Models.apiClient = c.apiClient
	return chat, nil
}
