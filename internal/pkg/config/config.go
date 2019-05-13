/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package config

import (
	"github.com/nalej/derrors"
	"github.com/nalej/edge-inventory-proxy/version"
	"github.com/rs/zerolog/log"
)

type Config struct {
	// Debug level is active.
	Debug bool
	// EIP port
	EipPort int
}

func (conf *Config) Validate() derrors.Error {

	if conf.EipPort <= 0 {
		return derrors.NewInvalidArgumentError("port must be valid")
	}

	return nil
}

func (conf *Config) Print() {
	log.Info().Str("app", version.AppVersion).Str("commit", version.Commit).Msg("Version")
	log.Info().Int("port", conf.EipPort).Msg("EIP port")
}