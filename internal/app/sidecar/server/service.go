package server

import (
	"github.com/nalej/edge-inventory-proxy/internal/app/sidecar/config"
	"github.com/rs/zerolog/log")

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
	if err != nil{
		log.Fatal().Err(err).Msg("cannot configure VPN")
	}

	err = helper.RegisterVPNAddress()
	if err != nil{
		log.Fatal().Err(err).Msg("cannot register Proxy DNS")
	}

	return nil
}
