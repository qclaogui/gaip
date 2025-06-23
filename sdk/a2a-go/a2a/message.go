package a2a

import (
	"context"
	"errors"
	"fmt"
	"io"

	pb "github.com/qclaogui/gaip/genproto/a2a/apiv1/a2apb"
	"google.golang.org/api/iterator"
)

type Messages struct {
	apiClient *apiClient
}

func (m Messages) SendMessage(ctx context.Context, msg *Message, config *SendMessageConfiguration) (*SendMessageResponse, error) {
	if config != nil {
		config.setDefaults()
	}

	req, err := m.newSendMessageRequest(msg, config)
	if err != nil {
		return nil, err
	}
	res, err := m.apiClient.a2aClient.SendMessage(ctx, req)
	if err != nil {
		return nil, err
	}

	return fromProto[SendMessageResponse](res)
}

func (m Messages) SendStreamingMessage(ctx context.Context, msg *Message, config *SendMessageConfiguration) *SendMessageResponseIterator {
	if config != nil {
		config.setDefaults()
	}

	iter := &SendMessageResponseIterator{}
	req, err := m.newSendMessageRequest(msg, config)
	if err != nil {
		iter.err = err
	} else {
		iter.sc, iter.err = m.apiClient.a2aClient.SendStreamingMessage(ctx, req)
	}
	return iter
}

func (m Messages) newSendMessageRequest(msg *Message, config *SendMessageConfiguration) (*pb.SendMessageRequest, error) {
	return pvCatchPanic(func() *pb.SendMessageRequest {
		req := &pb.SendMessageRequest{
			Request:       msg.toProto(),
			Configuration: config.toProto(),
		}
		debugPrint(req)
		return req
	})
}

// SendMessageResponseIterator is an iterator over StreamResponse.
type SendMessageResponseIterator struct {
	sc  pb.A2AService_SendStreamingMessageClient
	err error
}

func (iter *SendMessageResponseIterator) Next() (*StreamResponse, error) {
	if iter.err != nil {
		return nil, iter.err
	}

	res, err := iter.sc.Recv()
	iter.err = err
	if errors.Is(err, io.EOF) {
		return nil, iterator.Done
	}

	if err != nil {
		return nil, err
	}

	sp, err := fromProto[StreamResponse](res)
	if err != nil {
		iter.err = err
		return nil, err
	}
	return sp, nil
}

// NewUserMessageFromParts returns a *Message with a "user" role set and one or more
// parts.
func NewUserMessageFromParts(parts ...Part) *Message {
	return NewMessageFromParts(RoleUser, parts...)
}

// NewMessageFromParts returns a *Message with the specified role and one or more parts.
// This is a convenience function for creating messages with different roles.
func NewMessageFromParts(role Role, parts ...Part) *Message {
	message := &Message{Role: role, Content: make([]Part, 0, len(parts))}
	message.Content = append(message.Content, parts...)
	return message
}

// NewMessageFromText returns a *Message with a "user" role set and a single text part.
// This is a convenience function for creating user messages with text content.
func NewMessageFromText(role Role, text string) *Message {
	return NewMessageFromParts(role, Text(text))
}

// NewUserMessageFromText returns a *Message with a "user" role set and a single text part.
func NewUserMessageFromText(text string) *Message {
	return NewMessageFromParts(RoleUser, Text(text))
}

// Part is an interface that represents a part of a message.
// It can be a text, file, or data part.
// Each part can be converted to a protobuf representation.
type Part interface {
	toPart() *pb.Part
}

func partToProto(p Part) *pb.Part {
	if p == nil {
		return nil
	}
	return p.toPart()
}

func partFromProto(p *pb.Part) Part {
	switch d := p.Part.(type) {
	case *pb.Part_Text:
		return Text(d.Text)
	case *pb.Part_File:
		return (FilePart{}).fromProto(d.File)
	case *pb.Part_Data:
		return (DataPart{}).fromProto(d.Data)
	default:
		panic(fmt.Errorf("unknown Part.Part type %T", p.Part))
	}
}

// A Text is a piece of text, like a question or phrase.
type Text string

func (t Text) toPart() *pb.Part {
	return &pb.Part{
		Part: &pb.Part_Text{
			Text: string(t)},
	}
}

func (v DataPart) toPart() *pb.Part {
	return &pb.Part{
		Part: &pb.Part_Data{
			Data: v.toProto(),
		},
	}
}

func (f FilePart) toPart() *pb.Part {
	return &pb.Part{
		Part: &pb.Part_File{
			File: f.toProto(),
		},
	}
}

type FilePartData struct {
	FileURI   string
	FileBytes []byte
}

func populateFilePartTo(p *pb.FilePart, v *FilePart) {
	f := v.File

	if f.FileURI != "" {
		p.File = &pb.FilePart_FileWithUri{
			FileWithUri: f.FileURI,
		}
	} else if f.FileBytes != nil {
		p.File = &pb.FilePart_FileWithBytes{
			FileWithBytes: f.FileBytes,
		}
	}

	// If both fields of v.File are zero, leave p.File unset.
}

func populateFilePartFrom(v *FilePart, p *pb.FilePart) {
	if p.File == nil {
		return
	}
	switch f := p.File.(type) {
	case *pb.FilePart_FileWithUri:
		v.File.FileURI = f.FileWithUri
	case *pb.FilePart_FileWithBytes:
		v.File.FileBytes = f.FileWithBytes
	default:
		panic(fmt.Sprintf("unknown type of FilePart.File: %T", p.File))
	}
}

type MessageResponsePayload struct {
	Msg  *Message
	Task *Task
}

func populateMessageResponseTo(p *pb.SendMessageResponse, v *SendMessageResponse) {
	pl := v.Payload
	if pl.Msg != nil {
		p.Payload = &pb.SendMessageResponse_Msg{
			Msg: pl.Msg.toProto(),
		}
	} else if pl.Task != nil {
		p.Payload = &pb.SendMessageResponse_Task{
			Task: pl.Task.toProto(),
		}
	}
	// If both fields of v.Payload are nil, leave p.Payload unset.
}
func populateMessageResponseFrom(v *SendMessageResponse, p *pb.SendMessageResponse) {
	if p.Payload == nil {
		return
	}
	switch pl := p.Payload.(type) {
	case *pb.SendMessageResponse_Msg:
		v.Payload.Msg = (Message{}).fromProto(pl.Msg)
	case *pb.SendMessageResponse_Task:
		v.Payload.Task = (Task{}).fromProto(pl.Task)
	default:
		panic(fmt.Sprintf("unknown type of SendMessageResponse.Payload: %T", p.Payload))
	}
}

type StreamResponsePayload struct {
	Msg            *Message
	Task           *Task
	StatusUpdate   *TaskStatusUpdateEvent
	ArtifactUpdate *TaskArtifactUpdateEvent
}

func populateStreamResponseTo(p *pb.StreamResponse, v *StreamResponse) {
	pl := v.Payload
	if pl.Msg != nil {
		p.Payload = &pb.StreamResponse_Msg{
			Msg: pl.Msg.toProto(),
		}
	} else if pl.Task != nil {
		p.Payload = &pb.StreamResponse_Task{
			Task: pl.Task.toProto(),
		}
	} else if pl.StatusUpdate != nil {
		p.Payload = &pb.StreamResponse_StatusUpdate{
			StatusUpdate: pl.StatusUpdate.toProto(),
		}
	} else if pl.ArtifactUpdate != nil {
		p.Payload = &pb.StreamResponse_ArtifactUpdate{
			ArtifactUpdate: pl.ArtifactUpdate.toProto(),
		}
	}
	// If both fields of v.Payload are nil, leave p.Payload unset.
}
func populateStreamResponseFrom(v *StreamResponse, p *pb.StreamResponse) {
	if p.Payload == nil {
		return
	}
	switch pl := p.Payload.(type) {
	case *pb.StreamResponse_Msg:
		v.Payload.Msg = (Message{}).fromProto(pl.Msg)
	case *pb.StreamResponse_Task:
		v.Payload.Task = (Task{}).fromProto(pl.Task)
	case *pb.StreamResponse_StatusUpdate:
		v.Payload.StatusUpdate = (TaskStatusUpdateEvent{}).fromProto(pl.StatusUpdate)
	case *pb.StreamResponse_ArtifactUpdate:
		v.Payload.ArtifactUpdate = (TaskArtifactUpdateEvent{}).fromProto(pl.ArtifactUpdate)
	default:
		panic(fmt.Sprintf("unknown type of StreamResponse.Payload: %T", p.Payload))
	}
}

func (c *SendMessageConfiguration) setDefaults() {
	if c == nil {
		return
	}

	if c.PushNotification != nil {
		c.PushNotification.setDefaults()
	}
}
