/*
 * Copyright 2019 Nalej
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
	if conf.NetworkManagerAddress == "" {
		return derrors.NewInvalidArgumentError("networkManagerAddress cannot be empty")
	}

	if conf.Username == "" || conf.Password == "" {
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
