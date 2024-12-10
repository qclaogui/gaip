// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package rest

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

type wireBuffer struct {
	bytes.Buffer
	chunks []string
}

func (wb *wireBuffer) Flush() {
	wb.chunks = append(wb.chunks, wb.String())
	wb.Reset()
}

func TestServerStreamer(t *testing.T) {
	for _, tc := range []struct {
		name                     string
		messages                 []string
		expectedPerMessageChunks []string
		chunkSize                int
		expectedFixedSizeChunks  []string
	}{
		{
			name:     "Single chunk",
			messages: []string{"greetings"},
			expectedPerMessageChunks: []string{
				"[greetings",
				"]",
			},
			chunkSize: 4,
			expectedFixedSizeChunks: []string{
				"[gre",
				"etin",
				"gs]",
			},
		},
		{
			name:     "Two chunks",
			messages: []string{"greetings", "  earthling"},
			expectedPerMessageChunks: []string{
				"[greetings",
				",  earthling",
				"]",
			},
			chunkSize: 5,
			expectedFixedSizeChunks: []string{
				"[gree",
				"tings",
				",  ea",
				"rthli",
				"ng]",
			},
		},
		{
			name:     "Many chunks",
			messages: []string{"greetings", "  people", "of ", "Earth"},
			expectedPerMessageChunks: []string{
				"[greetings",
				",  people",
				",of ",
				",Earth",
				"]",
			},
			chunkSize: 4,
			expectedFixedSizeChunks: []string{
				"[gre",
				"etin",
				"gs, ",
				" peo",
				"ple,",
				"of ,",
				"Eart",
				"h]",
			},
		},
		{
			name:                     "No chunks",
			messages:                 []string{},
			expectedPerMessageChunks: nil,
			chunkSize:                3,
			expectedFixedSizeChunks:  nil,
		},
		{
			name:                     "Single empty chunk",
			messages:                 []string{""},
			expectedPerMessageChunks: nil,
			chunkSize:                3,
			expectedFixedSizeChunks:  nil,
		},
		{
			name:     "Intermediate empty chunk",
			messages: []string{"greetings", "", "earthlings"},
			expectedPerMessageChunks: []string{
				"[greetings",
				",earthlings",
				"]",
			},
			chunkSize: 4,
			expectedFixedSizeChunks: []string{
				"[gre",
				"etin",
				"gs,e",
				"arth",
				"ling",
				"s]",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			label := fmt.Sprintf("[%s:per-message chunks]", tc.name)
			streamWithChunkSize(t, label, 0, tc.messages, tc.expectedPerMessageChunks)
			label = fmt.Sprintf("[%s: chunk size %d]", tc.name, tc.chunkSize)
			streamWithChunkSize(t, label, tc.chunkSize, tc.messages, tc.expectedFixedSizeChunks)
		})
	}
}

func streamWithChunkSize(t *testing.T, label string, chunkSize int, messages, expectedChunks []string) {
	wire := &wireBuffer{}
	streamer, err := NewServerStreamer(wire, chunkSize)
	if err != nil {
		t.Fatalf("%s: could not construct ServerStreamer: %s", label, err)
	}

	for msgIdx, msg := range messages {
		if err = streamer.sendJSONArrayMessage([]byte(msg)); err != nil {
			t.Errorf("%s: error sending message #%d (%q): %s", label, msgIdx, msg, err)
			break
		}
	}
	if err = streamer.End(); err != nil {
		t.Errorf("%s: error ending stream: %s", label, err)
	}

	if got, want := wire.chunks, expectedChunks; !reflect.DeepEqual(got, want) {
		t.Errorf("%s: did not received expected chunks\n== got ===\n%#v\n== want ==\n%#v\n",
			label, got, want)
	}
}
