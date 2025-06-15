// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

// printRequests controls whether request protobufs are written to stderr.
var printRequests = false

func debugPrint(m proto.Message) {
	if !printRequests {
		return
	}
	fmt.Fprintln(os.Stderr, "--------")
	fmt.Fprintf(os.Stderr, "%T\n", m)
	fmt.Fprint(os.Stderr, prototext.Format(m))
	fmt.Fprintln(os.Stderr, "^^^^^^^^")
}
