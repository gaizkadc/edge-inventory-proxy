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
	// QueueAddress with the bus address
	QueueAddress string
}

func (conf *Config) Validate() derrors.Error {

	if conf.EipPort <= 0 {
		return derrors.NewInvalidArgumentError("port must be valid")
	}

	if conf.QueueAddress == ""{
		return derrors.NewInvalidArgumentError("queueAddress must not be empty")
	}

	return nil
}

func (conf *Config) Print() {
	log.Info().Str("app", version.AppVersion).Str("commit", version.Commit).Msg("Version")
	log.Info().Int("port", conf.EipPort).Msg("EIP port")
	log.Info().Str("URL", conf.QueueAddress).Msg("Queue")
}