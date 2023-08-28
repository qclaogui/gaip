package routeguide

import (
	"context"
	"testing"

	"google.golang.org/protobuf/proto"

	pb "github.com/qclaogui/golang-api-server/pkg/api/routeguidepb"
)

func Test_ServiceServer_GetFeature(t *testing.T) {
	ctx := context.Background()
	ssv, err := NewServiceServer(WithMemoryRepository())
	if err != nil {
		t.Fatalf("an error '%s' was not expected when NewServiceServer", err)
	}

	type args struct {
		ctx context.Context
		req *pb.Point
	}
	tests := []struct {
		name string
		ssv  pb.RouteGuideServer
		args args
		want *pb.Feature
	}{
		{
			name: "Looking for a valid feature",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.Point{Latitude: 409146138, Longitude: -746188906},
			},
			want: &pb.Feature{Name: "Berkshire Valley Management Area Trail, Jefferson, NJ, USA",
				Location: &pb.Point{Latitude: 409146138, Longitude: -746188906}},
		},
		{
			name: "Feature missing",
			ssv:  ssv,
			args: args{
				ctx: ctx,
				req: &pb.Point{},
			},
			want: &pb.Feature{Location: &pb.Point{}},
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
