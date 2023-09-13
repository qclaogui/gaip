// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package v1

import (
	"context"
	"testing"

	pb "github.com/qclaogui/golang-api-server/api/routeguide/v1/routeguidepb"
	util_log "github.com/qclaogui/golang-api-server/tools/log"
	"google.golang.org/protobuf/proto"
)

func Test_ServiceServer_GetFeature(t *testing.T) {
	ctx := context.Background()
	ssv, err := NewServiceServer(util_log.Logger, WithMemoryRepository())
	if err != nil {
		t.Fatalf("an error '%s' was not expected when NewServiceServer", err)
	}

	type args struct {
		ctx context.Context
		req *pb.GetFeatureRequest
	}
	tests := []struct {
		name string
		ssv  pb.RouteGuideServiceServer
		args args
		want *pb.GetFeatureResponse
	}{
		{
			name: "Looking for a valid feature",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.GetFeatureRequest{Point: &pb.Point{Latitude: 409146138, Longitude: -746188906}},
			},
			want: &pb.GetFeatureResponse{Feature: &pb.Feature{Name: "Berkshire Valley Management Area Trail, Jefferson, NJ, USA",
				Location: &pb.Point{Latitude: 409146138, Longitude: -746188906}}},
		},
		{
			name: "Feature missing",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.GetFeatureRequest{Point: &pb.Point{}},
			},
			want: &pb.GetFeatureResponse{Feature: &pb.Feature{Location: &pb.Point{}}},
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
