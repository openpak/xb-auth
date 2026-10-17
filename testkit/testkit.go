// SPDX-License-Identifier: AGPL-3.0-only

// Package testkit exposes the xb-auth KDC to test/e2e, which sits outside this
// service's internal/ tree. Test support only; no command imports it.
package testkit

import (
	"github.com/openpak/xb-auth/internal/devicekeys"
	"github.com/openpak/xb-auth/internal/kdc"
	"github.com/openpak/xb-common/xbaccountapi"
)

type (
	Config = kdc.Config
	KDC    = kdc.KDC
)

func DefaultConfig() Config { return kdc.DefaultConfig() }

// Account is what xb-auth consumes from xb-account (the *xbaccountapi.Client).
type Account interface {
	xbaccountapi.Resolver
	xbaccountapi.DeviceKeys
}

// New wires the KDC exactly as cmd/xb-auth does: device keys come from xb-account.
func New(cfg Config, a Account) (*KDC, error) {
	return kdc.New(cfg, a, devicekeys.Remote{Keys: a})
}
