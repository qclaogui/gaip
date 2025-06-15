// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

import (
	"context"
	"errors"
	"fmt"
	"strings"

	pb "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta/generativelanguagepb"
)

// A Client is a Google generative AI client.
type Client struct {
	// Models provides access to the Models service.
	Models *Models
	// Live provides access to the Live service.
	Live *Live
	// Caches provides access to the Caches service.
	Caches *Caches
	// Chats provides util functions for creating a new chat session.
	Chats *Chats
	// Files provides access to the Files service.
	Files *Files
	// Operations provides access to long-running operations.
	Operations *Operations
}

// NewClient creates a new GenAI client.
func NewClient(ctx context.Context) (*Client, error) {
	ac, err := newAPIClient(ctx)
	if err != nil {
		return nil, err
	}

	c := &Client{
		Models:     &Models{apiClient: ac},
		Live:       &Live{apiClient: ac},
		Caches:     &Caches{apiClient: ac},
		Chats:      &Chats{apiClient: ac},
		Files:      &Files{apiClient: ac},
		Operations: &Operations{apiClient: ac},
	}

	return c, nil
}

// transformSlice applies f to each element of from and returns
// a new slice with the results.
func transformSlice[From, To any](from []From, f func(From) To) []To {
	if from == nil {
		return nil
	}
	to := make([]To, len(from))
	for i, e := range from {
		to[i] = f(e)
	}
	return to
}

func fromProto[V interface{ fromProto(P) *V }, P any](p P) (*V, error) {
	var v V
	return pvCatchPanic(func() *V { return v.fromProto(p) })
}

func protoToResponse(resp *pb.GenerateContentResponse) (*GenerateContentResponse, error) {
	gcp, err := fromProto[GenerateContentResponse](resp)
	if err != nil {
		return nil, err
	}
	if gcp == nil {
		return nil, errors.New("empty response from model")
	}
	// Assume a non-nil PromptFeedback is an error.
	if gcp.PromptFeedback != nil && gcp.PromptFeedback.BlockReason != BlockReasonUnspecified {
		return nil, &BlockedError{PromptFeedback: gcp.PromptFeedback}
	}

	// If any candidate is blocked, error.
	for _, c := range gcp.Candidates {
		if c.FinishReason == FinishReasonSafety || c.FinishReason == FinishReasonRecitation {
			return nil, &BlockedError{Candidate: c}
		}
	}
	return gcp, nil
}

// A BlockedError indicates that the model's response was blocked.
// There can be two underlying causes: the prompt or a candidate response.
type BlockedError struct {
	// If non-nil, the model's response was blocked.
	// Consult the FinishReason field for details.
	Candidate *Candidate

	// If non-nil, there was a problem with the prompt.
	PromptFeedback *PromptFeedback
}

func (e *BlockedError) Error() string {
	var b strings.Builder
	fmt.Fprintf(&b, "blocked: ")
	if e.Candidate != nil {
		fmt.Fprintf(&b, "candidate: %s", e.Candidate.FinishReason)
	}
	if e.PromptFeedback != nil {
		if e.Candidate != nil {
			fmt.Fprintf(&b, ", ")
		}
		fmt.Fprintf(&b, "prompt: %v", e.PromptFeedback.BlockReason)
	}
	return b.String()
}
