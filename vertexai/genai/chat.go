// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

import "context"

// A ChatSession provides interactive chat.
type ChatSession struct {
	m       *GenerativeModel
	History []*Content
}

// StartChat starts a chat session.
func (m *GenerativeModel) StartChat() *ChatSession {
	return &ChatSession{m: m}
}

// SendMessage sends a request to the model as part of a chat session.
func (cs *ChatSession) SendMessage(ctx context.Context, parts ...Part) (*GenerateContentResponse, error) {
	// Call the underlying client with the entire history plus the argument Content.
	cs.History = append(cs.History, NewUserContent(parts...))

	req := cs.m.newGenerateContentRequest(cs.History...)
	var cc int32 = 1
	req.GenerationConfig.CandidateCount = &cc

	resp, err := cs.m.generateContent(ctx, req)
	if err != nil {
		return nil, err
	}

	cs.addToHistory(resp.Candidates)
	return resp, nil
}

// SendMessageStream is like SendMessage, but with a streaming request.
func (cs *ChatSession) SendMessageStream(ctx context.Context, parts ...Part) *GenerateContentResponseIterator {
	// Call the underlying client with the entire history plus the argument Content.
	cs.History = append(cs.History, NewUserContent(parts...))

	req := cs.m.newGenerateContentRequest(cs.History...)
	var cc int32 = 1
	req.GenerationConfig.CandidateCount = &cc

	streamClient, err := cs.m.c.pc.StreamGenerateContent(ctx, req)

	return &GenerateContentResponseIterator{
		sc:  streamClient,
		err: err,
		cs:  cs,
	}
}

// By default, use the first candidate for history. The user can modify that if they want.
func (cs *ChatSession) addToHistory(cands []*Candidate) {
	if len(cands) > 0 {
		c := cands[0].Content
		if c == nil {
			return
		}
		cs.History = append(cs.History, copySanitizedModelContent(c))
	}
}

// copySanitizedModelContent creates a (shallow) copy of c with role set to
// model and empty text parts removed.
func copySanitizedModelContent(c *Content) *Content {
	newc := &Content{Role: roleModel}
	for _, part := range c.Parts {
		if t, ok := part.(Text); !ok || len(string(t)) > 0 {
			newc.Parts = append(newc.Parts, part)
		}
	}
	return newc
}
