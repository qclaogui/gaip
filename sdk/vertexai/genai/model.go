// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	pb "github.com/qclaogui/gaip/genproto/aiplatform/apiv1beta1/aiplatformpb"
	"google.golang.org/api/iterator"
)

// GenerativeModel is a model that can generate text.
// Create one with [Client.GenerativeModel], then configure
// it by setting the exported fields.
//
// The model holds all the config for a GenerateContentRequest, so the GenerateContent method
// can use a vararg for the content.
type GenerativeModel struct {
	c        *Client
	name     string
	fullName string

	GenerationConfig
	SafetySettings    []*SafetySetting
	Tools             []*Tool
	ToolConfig        *ToolConfig // configuration for tools
	SystemInstruction *Content
	// The name of the CachedContent to use.
	// Must have already been created with [Client.CreateCachedContent].
	CachedContentName string
}

// GenerativeModel creates a new instance of the named model.
// name is a string model name like "gemini-1.0-pro" or "models/gemini-1.0-pro"
// for Google-published models.
// See https://cloud.google.com/vertex-ai/generative-ai/docs/learn/model-versioning
// for details on model naming and versioning, and
// https://cloud.google.com/vertex-ai/generative-ai/docs/model-garden/explore-models
// for providing model garden names. The SDK isn't familiar with custom model
// garden models, and will pass your model name to the backend API server.
func (c *Client) GenerativeModel(name string) *GenerativeModel {
	return &GenerativeModel{
		c:        c,
		name:     name,
		fullName: inferFullModelName(c.projectID, c.location, name),
	}
}

// inferFullModelName infers the full model name (with all the required prefixes)
func inferFullModelName(project, location, name string) string {
	pubName := name
	if !strings.Contains(name, "/") {
		pubName = "publishers/google/models/" + name
	} else if strings.HasPrefix(name, "models/") {
		pubName = "publishers/google/" + name
	}
	if !strings.HasPrefix(pubName, "publishers/") {
		return pubName
	}
	return fmt.Sprintf("projects/%s/locations/%s/%s", project, location, pubName)
}

// Name returns the name of the model.
func (m *GenerativeModel) Name() string {
	return m.name
}

// GenerateContent produces a single request and response.
func (m *GenerativeModel) GenerateContent(ctx context.Context, parts ...Part) (*GenerateContentResponse, error) {
	req, err := m.newGenerateContentRequest(NewUserContent(parts...))
	if err != nil {
		return nil, err
	}
	resp, err := m.c.pc.GenerateContent(ctx, req)
	if err != nil {
		return nil, err
	}
	return protoToResponse(resp)
}

// GenerateContentStream returns an iterator that enumerates responses.
func (m *GenerativeModel) GenerateContentStream(ctx context.Context, parts ...Part) *GenerateContentResponseIterator {
	iter := &GenerateContentResponseIterator{}
	req, err := m.newGenerateContentRequest(NewUserContent(parts...))
	if err != nil {
		iter.err = err
	} else {
		iter.sc, iter.err = m.c.pc.StreamGenerateContent(ctx, req)
	}
	return iter
}

func (m *GenerativeModel) generateContent(ctx context.Context, req *pb.GenerateContentRequest) (*GenerateContentResponse, error) {
	iter := &GenerateContentResponseIterator{}
	iter.sc, iter.err = m.c.pc.StreamGenerateContent(ctx, req)

	for {
		_, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			return iter.MergedResponse(), nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func (m *GenerativeModel) newGenerateContentRequest(contents ...*Content) (*pb.GenerateContentRequest, error) {
	return pvCatchPanic(func() *pb.GenerateContentRequest {
		var cc *string
		if m.CachedContentName != "" {
			cc = &m.CachedContentName
		}

		req := &pb.GenerateContentRequest{
			Model:             m.fullName,
			Contents:          pvTransformSlice(contents, (*Content).toProto),
			SafetySettings:    pvTransformSlice(m.SafetySettings, (*SafetySetting).toProto),
			Tools:             pvTransformSlice(m.Tools, (*Tool).toProto),
			ToolConfig:        m.ToolConfig.toProto(),
			GenerationConfig:  m.GenerationConfig.toProto(),
			SystemInstruction: m.SystemInstruction.toProto(),
			CachedContent:     *cc,
		}
		debugPrint(req)
		return req
	})
}

// GenerateContentResponseIterator is an iterator over GnerateContentResponse.
type GenerateContentResponseIterator struct {
	sc     pb.PredictionService_StreamGenerateContentClient
	err    error
	merged *GenerateContentResponse
	cs     *ChatSession
}

// Next returns the next response.
func (iter *GenerateContentResponseIterator) Next() (*GenerateContentResponse, error) {
	if iter.err != nil {
		return nil, iter.err
	}

	resp, err := iter.sc.Recv()
	iter.err = err
	if errors.Is(err, io.EOF) {
		if iter.cs != nil && iter.merged != nil {
			iter.cs.addToHistory(iter.merged.Candidates)
		}
		return nil, iterator.Done
	}
	if err != nil {
		return nil, err
	}

	gcp, err := protoToResponse(resp)
	if err != nil {
		iter.err = err
		return nil, err
	}

	// Merge this response in with the ones we've already seen.
	iter.merged = joinResponses(iter.merged, gcp)
	// If this is part of a ChatSession, remember the response for the history.
	return gcp, nil
}

// joinResponses  merges the two responses, which should be the result of a streaming call.
// The first argument is modified.
func joinResponses(dest, src *GenerateContentResponse) *GenerateContentResponse {
	if dest == nil {
		return src
	}
	dest.Candidates = joinCandidateLists(dest.Candidates, src.Candidates)
	// Keep dest.PromptFeedback.
	// TODO: Take the last UsageMetadata.
	return dest
}

func joinCandidateLists(dest, src []*Candidate) []*Candidate {
	indexToSrcCandidate := map[int32]*Candidate{}
	for _, s := range src {
		indexToSrcCandidate[s.Index] = s
	}
	for _, d := range dest {
		s := indexToSrcCandidate[d.Index]
		if s != nil {
			d.Content = joinContent(d.Content, s.Content)
			// Take the last of these.
			d.FinishReason = s.FinishReason
			d.FinishMessage = s.FinishMessage
			d.SafetyRatings = s.SafetyRatings
			d.CitationMetadata = joinCitationMetadata(d.CitationMetadata, s.CitationMetadata)
		}
	}
	return dest
}

func joinCitationMetadata(dest, src *CitationMetadata) *CitationMetadata {
	if dest == nil {
		return src
	}
	if src == nil {
		return dest
	}
	dest.Citations = append(dest.Citations, src.Citations...)
	return dest
}

func joinContent(dest, src *Content) *Content {
	if dest == nil {
		return src
	}
	if src == nil {
		return dest
	}
	// Assume roles are the same.
	dest.Parts = joinParts(dest.Parts, src.Parts)
	return dest
}

func joinParts(dest, src []Part) []Part {
	return mergeTexts(append(dest, src...))
}

func mergeTexts(in []Part) []Part {
	var out []Part
	i := 0
	for i < len(in) {
		if t, ok := in[i].(Text); ok {
			texts := []string{string(t)}
			var j int
			for j = i + 1; j < len(in); j++ {
				if t, ok := in[j].(Text); ok {
					texts = append(texts, string(t))
				} else {
					break
				}
			}
			// j is just after the last Text.
			out = append(out, Text(strings.Join(texts, "")))
			i = j
		} else {
			out = append(out, in[i])
			i++
		}
	}
	return out
}

// MergedResponse returns the result of combining all the streamed responses seen so far.
// After iteration completes, the merged response should match the response obtained without streaming
// (that is, if [GenerativeModel.GenerateContent] were called).
func (iter *GenerateContentResponseIterator) MergedResponse() *GenerateContentResponse {
	return iter.merged
}

func protoToResponse(resp *pb.GenerateContentResponse) (*GenerateContentResponse, error) {
	gcp, err := fromProto[GenerateContentResponse](resp)
	if err != nil {
		return nil, err
	}
	if gcp == nil {
		return nil, errors.New("empty response from model")
	}

	if gcp.PromptFeedback != nil {
		return nil, &BlockedError{PromptFeedback: gcp.PromptFeedback}
	}
	// Assume a non-nil PromptFeedback is an error.
	// TODO: confirm.
	// if gcp.PromptFeedback != nil && gcp.PromptFeedback.BlockReason != BlockReasonUnspecified {
	// 	return nil, &BlockedError{PromptFeedback: gcp.PromptFeedback}
	// }

	// If any candidate is blocked, error.
	// TODO: is this too harsh?
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
	// Consult the Candidate and SafetyRatings fields for details.
	Candidate *Candidate

	// If non-nil, there was a problem with the prompt.
	PromptFeedback *PromptFeedback
}

func (e *BlockedError) Error() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "blocked: ")
	if e.Candidate != nil {
		fmt.Fprintf(&sb, "candidate: %s", e.Candidate.FinishReason)
	}
	if e.PromptFeedback != nil {
		if e.Candidate != nil {
			fmt.Fprintf(&sb, ", ")
		}
		fmt.Fprintf(&sb, "prompt: %v (%s)", e.PromptFeedback.BlockReason, e.PromptFeedback.BlockReasonMessage)
	}
	return sb.String()
}
