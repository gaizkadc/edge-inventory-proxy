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

package server

import (
	"github.com/nalej/edge-inventory-proxy/internal/app/sidecar/config"
	"github.com/rs/zerolog/log"
)

type Service struct {
	Configuration config.Config
}

// NewService creates a new service.
func NewService(conf config.Config) *Service {
	return &Service{
		conf,
	}
}

func (s *Service) Run() error {

	vErr := s.Configuration.Validate()
	if vErr != nil {
		log.Fatal().Str("err", vErr.DebugReport()).Msg("invalid configuration")
	}

	s.Configuration.Print()

	helper := NewVpnHelper(s.Configuration)

	err := helper.ConfigureLocalVPN()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot configure VPN")
	}

	err = helper.RegisterVPNAddress()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register Proxy DNS")
	}

	return nil
}
