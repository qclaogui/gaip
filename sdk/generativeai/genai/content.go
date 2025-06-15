// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

import pb "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta/generativelanguagepb"

type Role string

const (
	RoleUser  = "user"
	RoleModel = "model"
)

func roleString(role Role) string {
	if role == "" {
		return "user"
	}
	return string(role)
}

// NewContentFromParts builds a Content from a list of parts and a [Role].
// If role is the empty string, it defaults to [RoleUser].
func NewContentFromParts(parts []*Part, role Role) *Content {
	return &Content{
		Parts: parts,
		Role:  roleString(role),
	}
}

// NewContentFromText builds a Content from a text string.
// If role is the empty string, it defaults to [RoleUser].
func NewContentFromText(text string, role Role) *Content {
	return &Content{
		Parts: []*Part{
			NewPartFromText(text),
		},
		Role: roleString(role),
	}
}

// NewContentFromBytes builds a Content from a byte slice and mime type.
// If role is the empty string, it defaults to [RoleUser].
func NewContentFromBytes(data []byte, mimeType string, role Role) *Content {
	return &Content{
		Parts: []*Part{
			NewPartFromBytes(data, mimeType),
		},
		Role: roleString(role),
	}
}

// NewContentFromURI builds a Content from a file URI and mime type.
// If role is the empty string, it defaults to [RoleUser].
func NewContentFromURI(fileURI, mimeType string, role Role) *Content {
	return &Content{
		Parts: []*Part{
			NewPartFromURI(fileURI, mimeType),
		},
		Role: roleString(role),
	}
}

// populatePartTo populates some fields of p from v.
func populatePartTo(p *pb.Part, v *Part) {
	p.Data = nil
	if v.Text != "" {
		p.Data = &pb.Part_Text{
			Text: v.Text,
		}
	}
	if v.InlineData != nil {
		p.Data = &pb.Part_InlineData{
			InlineData: v.InlineData.toProto(),
		}
	}

	if v.FileData != nil {
		p.Data = &pb.Part_FileData{
			FileData: v.FileData.toProto(),
		}
	}

	if v.FunctionCall != nil {
		p.Data = &pb.Part_FunctionCall{
			FunctionCall: v.FunctionCall.toProto(),
		}
	}

	if v.FunctionResponse != nil {
		p.Data = &pb.Part_FunctionResponse{
			FunctionResponse: v.FunctionResponse.toProto(),
		}
	}

	if v.ExecutableCode != nil {
		p.Data = &pb.Part_ExecutableCode{
			ExecutableCode: v.ExecutableCode.toProto(),
		}
	}
	if v.CodeExecutionResult != nil {
		p.Data = &pb.Part_CodeExecutionResult{
			CodeExecutionResult: v.CodeExecutionResult.toProto(),
		}
	}
}

// populatePartFrom populates some fields of v from p.
func populatePartFrom(v *Part, p *pb.Part) {
	if p.Data == nil {
		return
	}
	switch d := p.Data.(type) {
	case *pb.Part_Text:
		v.Text = d.Text
	case *pb.Part_InlineData:
		v.InlineData = (Blob{}).fromProto(d.InlineData)
	case *pb.Part_FileData:
		v.FileData = (FileData{}).fromProto(d.FileData)
	case *pb.Part_FunctionCall:
		v.FunctionCall = (FunctionCall{}).fromProto(d.FunctionCall)
	case *pb.Part_FunctionResponse:
		v.FunctionResponse = (FunctionResponse{}).fromProto(d.FunctionResponse)
	case *pb.Part_ExecutableCode:
		v.ExecutableCode = (ExecutableCode{}).fromProto(d.ExecutableCode)
	case *pb.Part_CodeExecutionResult:
		v.CodeExecutionResult = (CodeExecutionResult{}).fromProto(d.CodeExecutionResult)
	}
}

// NewPartFromURI builds a Part from a given file URI and mime type.
func NewPartFromURI(fileURI, mimeType string) *Part {
	return &Part{
		FileData: &FileData{
			FileURI:  fileURI,
			MIMEType: mimeType,
		},
	}
}

// NewPartFromFile builds a Part from a given [File].
func NewPartFromFile(file File) *Part {
	return &Part{
		FileData: &FileData{
			FileURI:  file.URI,
			MIMEType: file.MIMEType,
		},
	}
}

// NewPartFromText builds a Part from a given text.
func NewPartFromText(text string) *Part {
	return &Part{
		Text: text,
	}
}

// NewPartFromBytes builds a Part from a given byte array and mime type.
func NewPartFromBytes(data []byte, mimeType string) *Part {
	return &Part{
		InlineData: &Blob{
			Data:     data,
			MIMEType: mimeType,
		},
	}
}
