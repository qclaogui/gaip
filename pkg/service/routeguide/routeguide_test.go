// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package routeguide

import (
	"context"
	"net"
	"strconv"
	"testing"

	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func Test_ServiceServer_GetFeature(t *testing.T) {
	ctx := context.Background()

	serverCfg := getServerConfig(t)
	srv, err := service.NewServer(serverCfg)
	require.NoError(t, err)

	// set repository database driver
	repoCfg := repository.Config{Driver: repository.DriverMemory}
	repo, err := repository.NewRouteGuide(repoCfg)
	require.NoError(t, err)

	cfg := Config{Repo: repo}
	ssv, err := New(cfg, srv)
	require.NoError(t, err)

	go func() { _ = srv.Run() }()
	t.Cleanup(srv.Stop)

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

// Generates server config, with gRPC listening on random port.
func getServerConfig(t *testing.T) service.Config {
	grpcHost, grpcPortNum := getHostnameAndRandomPort(t)
	httpHost, httpPortNum := getHostnameAndRandomPort(t)

	cfg := service.Config{
		HTTPListenAddress: httpHost,
		HTTPListenPort:    httpPortNum,

		GRPCListenAddress: grpcHost,
		GRPCListenPort:    grpcPortNum,

		GRPCServerMaxRecvMsgSize: 1024,
	}
	require.NoError(t, cfg.LogLevel.Set("info"))
	return cfg
}

func getHostnameAndRandomPort(t *testing.T) (string, int) {
	listen, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)

	host, port, err := net.SplitHostPort(listen.Addr().String())
	require.NoError(t, err)
	require.NoError(t, listen.Close())

	portNum, err := strconv.Atoi(port)
	require.NoError(t, err)
	return host, portNum
}
