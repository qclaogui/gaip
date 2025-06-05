package genai

import (
	"fmt"

	pb "github.com/qclaogui/gaip/genproto/aiplatform/apiv1beta1/aiplatformpb"
)

const (
	roleUser  = "user"
	roleModel = "model"
)

// A Part is either a Text, a Blob, or a FileData.
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
	switch d := p.Data.(type) {
	case *pb.Part_Text:
		return Text(d.Text)
	case *pb.Part_InlineData:
		return Blob{
			MIMEType: d.InlineData.MimeType,
			Data:     d.InlineData.Data,
		}
	case *pb.Part_FileData:
		return FileData{
			MIMEType: d.FileData.MimeType,
			FileURI:  d.FileData.FileUri,
		}
	case *pb.Part_FunctionCall:
		return *(FunctionCall{}).fromProto(d.FunctionCall)
	case *pb.Part_FunctionResponse:
		panic("FunctionResponse unimplemented")
		// return *(FunctionResponse{}).fromProto(d.FunctionResponse)
	default:
		panic(fmt.Errorf("unknown Part.Data type %T", p.Data))
	}
}

// A Text is a piece of text, like a question or phrase.
type Text string

func (t Text) toPart() *pb.Part {
	return &pb.Part{
		Data: &pb.Part_Text{Text: string(t)},
	}
}

func (v Blob) toPart() *pb.Part {
	return &pb.Part{
		Data: &pb.Part_InlineData{
			InlineData: v.toProto(),
		},
	}
}

func (f FileData) toPart() *pb.Part {
	return &pb.Part{
		Data: &pb.Part_FileData{
			FileData: f.toProto(),
		},
	}
}

func (f FunctionCall) toPart() *pb.Part {
	return &pb.Part{
		Data: &pb.Part_FunctionCall{
			FunctionCall: f.toProto(),
		},
	}
}

func (f FunctionResponse) toPart() *pb.Part {
	return &pb.Part{
		Data: &pb.Part_FunctionResponse{
			FunctionResponse: f.toProto(),
		},
	}
}

// func (e ExecutableCode) toPart() *pb.Part {
// 	return &pb.Part{
// 		Data: &pb.Part_ExecutableCode{
// 			ExecutableCode: e.toProto(),
// 		},
// 	}
// }

// ImageData is a convenience function for creating an image
// Blob for input to a model.
// The format should be the second part of the MIME type, after "image/".
// For example, for a PNG image, pass "png".
func ImageData(format string, data []byte) Blob {
	return Blob{
		MIMEType: "image/" + format,
		Data:     data,
	}
}

// Ptr returns a pointer to its argument.
// It can be used to initialize pointer fields:
//
//	model.Temperature = genai.Ptr[float32](0.1)
func Ptr[T any](t T) *T { return &t }

// SetCandidateCount sets the CandidateCount field.
func (c *GenerationConfig) SetCandidateCount(x int32) { c.CandidateCount = &x }

// SetMaxOutputTokens sets the MaxOutputTokens field.
func (c *GenerationConfig) SetMaxOutputTokens(x int32) { c.MaxOutputTokens = &x }

// SetTemperature sets the Temperature field.
func (c *GenerationConfig) SetTemperature(x float32) { c.Temperature = &x }

// SetTopP sets the TopP field.
func (c *GenerationConfig) SetTopP(x float32) { c.TopP = &x }

// SetTopK sets the TopK field.
func (c *GenerationConfig) SetTopK(x int32) { c.TopK = &x }

// NewUserContent returns a [Content] with a "user" role set and one or more
// parts.
func NewUserContent(parts ...Part) *Content {
	content := &Content{Role: roleUser, Parts: []Part{}}
	for _, part := range parts {
		if part == nil {
			continue
		}
		content.Parts = append(content.Parts, part)
	}

	return content
}
