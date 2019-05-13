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
	// Organization ID
	OrganizationId string
	// Edge Controller ID
	EdgeControllerId string
	// Edge Controller IP
	EdgeControllerIp string
}

func (conf *Config) Validate() derrors.Error {

	if conf.EipPort <= 0 {
		return derrors.NewInvalidArgumentError("port must be valid")
	}

	if conf.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("organization id must be set")
	}

	if conf.EdgeControllerId == "" {
		return derrors.NewInvalidArgumentError("edge controller id must be set")
	}

	if conf.EdgeControllerIp == "" {
		return derrors.NewInvalidArgumentError("edge controller ip must be set")
	}

	return nil
}

func (conf *Config) Print() {
	log.Info().Str("app", version.AppVersion).Str("commit", version.Commit).Msg("Version")
	log.Info().Int("port", conf.EipPort).Msg("EIP port")
	log.Info().Str("orgid", conf.EdgeControllerId).Msg("Organization ID")
	log.Info().Str("controller_id", conf.EdgeControllerId).Msg("Controller ID")
	log.Info().Str("controller_ip", conf.EdgeControllerIp).Msg("Controller IP")
}