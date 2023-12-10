// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package routeguide

import (
	"context"
	"testing"

	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
	lg "github.com/qclaogui/gaip/tools/log"
	"google.golang.org/protobuf/proto"
)

func Test_ServiceServer_GetFeature(t *testing.T) {
	ctx := context.Background()
	ssv, err := NewServiceServer(lg.Logger, WithMemoryRepository())
	if err != nil {
		t.Fatalf("an error '%s' was not expected when NewServiceServer", err)
	}

	type args struct {
		ctx context.Context
		req *routeguidepb.GetFeatureRequest
	}
	tests := []struct {
		name string
		ssv  routeguidepb.RouteGuideServiceServer
		args args
		want *routeguidepb.GetFeatureResponse
	}{
		{
			name: "Looking for a valid feature",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &routeguidepb.GetFeatureRequest{Point: &routeguidepb.Point{Latitude: 409146138, Longitude: -746188906}},
			},
			want: &routeguidepb.GetFeatureResponse{Feature: &routeguidepb.Feature{Name: "Berkshire Valley Management Area Trail, Jefferson, NJ, USA",
				Location: &routeguidepb.Point{Latitude: 409146138, Longitude: -746188906}}},
		},
		{
			name: "Feature missing",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &routeguidepb.GetFeatureRequest{Point: &routeguidepb.Point{}},
			},
			want: &routeguidepb.GetFeatureResponse{Feature: &routeguidepb.Feature{Location: &routeguidepb.Point{}}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.ssv.GetFeature(tc.args.ctx, tc.args.req)
			if err != nil {
				t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
					"✘got: %v\n\x1b[92m"+
					"want: %v\x1b[39m", got, nil)
			}

			if !proto.Equal(got, tc.want) {
				t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
					"✘got: %v\n\x1b[92m"+
					"want: %v\x1b[39m", got, tc.want)
			}
		})
	}

}
