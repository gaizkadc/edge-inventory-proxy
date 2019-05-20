/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package edge_inventory_proxy

import (
	"fmt"
	"github.com/nalej/edge-inventory-proxy/internal/pkg/config"
	"github.com/nalej/grpc-edge-inventory-proxy-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
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

// Run the service, launch the REST service handler.
func (s *Service) Run() error {
	cErr := s.Configuration.Validate()
	if cErr != nil {
		log.Fatal().Str("err", cErr.DebugReport()).Msg("invalid configuration")
	}
	s.Configuration.Print()


	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Configuration.EipPort))
	if err != nil {
		log.Fatal().Errs("failed to listen: %v", []error{err})
	}

	// Create handlers
	manager := NewManager(s.Configuration)
	handler := NewHandler(manager)

	// gRPC Server
	grpcServer := grpc.NewServer()

	grpc_edge_inventory_proxy_go.RegisterEdgeInventoryProxyServer(grpcServer, handler)

	if s.Configuration.Debug{
		log.Info().Msg("Enabling gRPC server reflection")
		// Register reflection service on gRPC server.
		reflection.Register(grpcServer)
	}
	log.Info().Int("port", s.Configuration.EipPort).Msg("Launching gRPC server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Errs("failed to serve: %v", []error{err})
	}
	return nil
}