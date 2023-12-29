// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package service

import (
	"net"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	lg "github.com/qclaogui/gaip/tools/log"
	"github.com/stretchr/testify/require"
)

func TestName(t *testing.T) {
	serverCfg := getServerConfig(t)
	srv, err := NewServer(serverCfg)
	require.NoError(t, err)

	go func() { _ = srv.Run() }()
	t.Cleanup(srv.Stop)

	router := mux.NewRouter()
	metrics := NewServerMetrics(serverCfg)
	_, _, _ = newEndpointREST(serverCfg, router, metrics, lg.Logger)
	_, _, _ = newEndpointGRPC(serverCfg, router, metrics, lg.Logger)

}

// Generates server config, with gRPC listening on random port.
func getServerConfig(t *testing.T) Config {
	grpcHost, grpcPortNum := getHostnameAndRandomPort(t)
	httpHost, httpPortNum := getHostnameAndRandomPort(t)

	cfg := Config{
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
