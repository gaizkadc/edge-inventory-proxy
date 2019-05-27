/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package server

import (
	"fmt"
	"github.com/nalej/derrors"
	"github.com/nalej/edge-inventory-proxy/internal/app/eicproxy/config"
	"github.com/nalej/edge-inventory-proxy/internal/app/eicproxy/server/ecinventory"
	"github.com/nalej/edge-inventory-proxy/internal/app/eicproxy/server/ecproxy"
	"github.com/nalej/grpc-edge-inventory-proxy-go"
	"github.com/nalej/nalej-bus/pkg/queue/inventory/events"
	"github.com/nalej/nalej-bus/pkg/bus/pulsar-comcast"
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

type BusClients struct {
	inventoryEventsProducer * events.InventoryEventsProducer
}

func (s*Service) GetBusClients() (*BusClients, derrors.Error) {
	queueClient := pulsar_comcast.NewClient(s.Configuration.QueueAddress)
	invEventProducer , err := events.NewInventoryEventsProducer(queueClient, "eicproxy-invevents")
	if err != nil {
		return nil, err
	}
	return &BusClients{
		invEventProducer,
	}, nil
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

	busClients, bErr := s.GetBusClients()
	if err != nil{
		log.Fatal().Str("err", bErr.DebugReport()).Msg("Cannot create bus clients")
	}

	// Create handlers
	invManager := ecinventory.NewManager(s.Configuration, busClients.inventoryEventsProducer)
	invHandler := ecinventory.NewHandler(invManager)

	proxyManager := ecproxy.NewManager(s.Configuration)
	proxyHandler := ecproxy.NewHandler(proxyManager)

	// gRPC Server
	grpcServer := grpc.NewServer()

	grpc_edge_inventory_proxy_go.RegisterEdgeInventoryProxyServer(grpcServer, invHandler)
	grpc_edge_inventory_proxy_go.RegisterEdgeControllerProxyServer(grpcServer, proxyHandler)

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