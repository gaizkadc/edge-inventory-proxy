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
	"fmt"
	"github.com/nalej/derrors"
	"github.com/nalej/edge-inventory-proxy/internal/app/eicproxy/config"
	"github.com/nalej/edge-inventory-proxy/internal/app/eicproxy/server/ecinventory"
	"github.com/nalej/edge-inventory-proxy/internal/app/eicproxy/server/ecproxy"
	"github.com/nalej/grpc-edge-inventory-proxy-go"
	"github.com/nalej/grpc-inventory-manager-go"
	"github.com/nalej/nalej-bus/pkg/bus/pulsar-comcast"
	"github.com/nalej/nalej-bus/pkg/queue/inventory/events"
	"github.com/nalej/nalej-bus/pkg/queue/inventory/ops"
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
	inventoryOpsProducer * ops.InventoryOpsProducer
}

func (s*Service) GetBusClients() (*BusClients, derrors.Error) {
	queueClient := pulsar_comcast.NewClient(s.Configuration.QueueAddress)
	invEventProducer , err := events.NewInventoryEventsProducer(queueClient, "eicproxy-invevents")
	if err != nil {
		return nil, err
	}

	invOpProducer, err := ops.NewInventoryOpsProducer(queueClient, "eicproxy-invops")
	if err != nil {
		return nil, err
	}

	return &BusClients{
		inventoryEventsProducer:invEventProducer,
		inventoryOpsProducer:invOpProducer,
	}, nil
}

type Clients struct {
	agentClient grpc_inventory_manager_go.AgentClient
}
func (s *Service) GetClients() (*Clients, derrors.Error) {

	invManagerConn, err := grpc.Dial(s.Configuration.InventoryManagerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with inventory manager")
	}

	agentClient := grpc_inventory_manager_go.NewAgentClient(invManagerConn)

	return &Clients{agentClient:agentClient}, nil
}

// Run the service, launch the REST service handler.
func (s *Service) Run() error {
	vErr := s.Configuration.Validate()
	if vErr != nil {
		log.Fatal().Str("err", vErr.DebugReport()).Msg("invalid configuration")
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

	clients, cErr := s.GetClients()
	if cErr != nil {
		log.Fatal().Str("err", cErr.DebugReport()).Msg("cannot generate clients")
		return cErr
	}


	// Create handlers
	invManager := ecinventory.NewManager(s.Configuration, busClients.inventoryEventsProducer, busClients.inventoryOpsProducer, clients.agentClient)
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