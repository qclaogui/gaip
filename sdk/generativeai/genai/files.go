// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

import pb "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta/generativelanguagepb"

type Files struct {
	apiClient *apiClient
}

// FileMetadata holds metadata about a file.
type FileMetadata struct {
	// Set if the file contains video.
	Video *VideoMetadata
}

func populateFileTo(p *pb.File, f *File) {
	p.Metadata = nil
	if f.Metadata == nil {
		return
	}
	if f.Metadata.Video != nil {
		p.Metadata = &pb.File_VideoMetadata{
			VideoMetadata: f.Metadata.Video.toProto(),
		}
	}
}

func populateFileFrom(f *File, p *pb.File) {
	f.Metadata = nil
	if p.Metadata == nil {
		return
	}
	switch m := p.Metadata.(type) {
	case *pb.File_VideoMetadata:
		f.Metadata = &FileMetadata{
			Video: (VideoMetadata{}).fromProto(m.VideoMetadata),
		}
	default:
		// ignore other types
		// TODO: signal a problem
	}
}
