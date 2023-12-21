// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package routeguide

import (
	"context"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
	"github.com/qclaogui/gaip/pkg/service/routeguide/repository"
	lg "github.com/qclaogui/gaip/tools/log"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func Test_ServiceServer_GetFeature(t *testing.T) {
	ctx := context.Background()
	var cfg = Config{}
	// set repository database driver
	cfg.RepoCfg.Driver = repository.DriverMemory

	repo, err := repository.NewRepository(cfg.RepoCfg)
	require.NoError(t, err)

	ssv := &RouteGuide{
		Cfg:        cfg,
		logger:     lg.Logger,
		Registerer: prometheus.DefaultRegisterer,
		repo:       repo,
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
