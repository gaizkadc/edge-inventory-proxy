package eip_sidecar

import (
	"github.com/nalej/derrors"
	"github.com/nalej/edge-inventory-proxy/version"
	"github.com/rs/zerolog/log"
)

type Config struct {
	// Debug level is active.
	Debug bool

	VpnAddress string
}

func (conf *Config) Validate() derrors.Error {

	if conf.VpnAddress == "" {
		return derrors.NewInvalidArgumentError("VpnAddress must be defined")
	}

	return nil
}

func (conf *Config) Print() {
	log.Info().Str("app", version.AppVersion).Str("commit", version.Commit).Msg("Version")
	log.Info().Str("vpnAddress", conf.VpnAddress).Msg("EIP port")
}
