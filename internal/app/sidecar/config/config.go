package config

import (
	"github.com/nalej/derrors"
	"github.com/nalej/edge-inventory-proxy/version"
	"github.com/rs/zerolog/log"
	"strings"
)

type Config struct {
	// Debug level is active.
	Debug bool

	//VPNAddress with the address of the VPN server accepting connections
	VPNAddress string

	// ProxyName with the name of the current proxy
	ProxyName string

	// NetworkManagerAddress with the address of the network manager
	NetworkManagerAddress string

	// Username to connect to the VPN
	Username string
	// Password to connect to the VPN
	Password string
}

func (conf *Config) Validate() derrors.Error {

	if conf.VPNAddress == "" {
		return derrors.NewInvalidArgumentError("VpnAddress must be defined")
	}
	if conf.ProxyName == "" {
		return derrors.NewInvalidArgumentError("proxyName cannot be empty")
	}
	if conf.NetworkManagerAddress == ""{
		return derrors.NewInvalidArgumentError("networkManagerAddress cannot be empty")
	}

	if conf.Username == "" || conf.Password == ""{
		return derrors.NewInvalidArgumentError("username and password must be specified")
	}

	return nil
}

func (conf *Config) Print() {
	log.Info().Str("app", version.AppVersion).Str("commit", version.Commit).Msg("Version")
	log.Info().Str("URL", conf.VPNAddress).Msg("VPN Server")
	log.Info().Str("URL", conf.NetworkManagerAddress).Msg("Network Manager")
	log.Info().Str("proxyName", conf.ProxyName).Msg("Proxy identity")
	log.Info().Str("username", conf.Username).Str("password", strings.Repeat("*", len(conf.Password))).Msg("VPN credentials")
}
