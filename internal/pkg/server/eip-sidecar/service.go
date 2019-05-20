package eip_sidecar

import (
	"github.com/rs/zerolog/log")

type Service struct {
	Configuration Config
}
// NewService creates a new service.
func NewService(conf Config) *Service {
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
