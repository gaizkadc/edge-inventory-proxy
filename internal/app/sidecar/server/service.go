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
	// connect to VPN
	vpnErr := NewVpnHelper(s.Configuration.VpnAddress).ConfigureLocalVPN()
	if vpnErr != nil {
		log.Fatal().Errs("failed to connect VPN: %v", []error{vpnErr})
	}
	return nil
}
