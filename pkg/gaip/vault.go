// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"github.com/go-kit/log/level"
	"github.com/qclaogui/gaip/pkg/vault"
)

// initVault init Vault
func (g *Gaip) initVault() error {
	if !g.Cfg.VaultCfg.Enabled {
		_ = level.Warn(g.Server.Log).Log("msg", "vault.enabled=false")
		return nil
	}

	v, err := vault.New(g.Cfg.VaultCfg)
	if err != nil {
		return err
	}
	g.Vault = v

	// Update Configs - KVStore
	//g.Cfg.MemberlistKV.TCPTransport.TLS.Reader = g.VaultCfg

	// Update Configs - GRPCServer Clients
	//g.Cfg.Worker.GRPCClientConfig.TLS.Reader = g.VaultCfg

	return nil
}
