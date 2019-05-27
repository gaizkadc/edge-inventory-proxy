/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package ecproxy

import (
	"context"
	"errors"
	"fmt"
	"github.com/nalej/derrors"
	"github.com/nalej/edge-inventory-proxy/internal/app/eicproxy/config"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-edge-controller-go"
	"github.com/nalej/grpc-inventory-go"
	"github.com/nalej/grpc-inventory-manager-go"
	"github.com/nalej/grpc-utils/pkg/conversions"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"time"
)

const ControllerTimeout = time.Minute

// Manager structure with the entities involved in the management of VPN users
type Manager struct {
	config config.Config
}

func NewManager(config config.Config) Manager{
	return Manager{
		config: config,
	}
}

func (m*Manager) getIP(dnsName string) (*string, error){
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", net.JoinHostPort("51.144.230.81", "53"))
		},
	}

	ip, err := resolver.LookupIPAddr(context.Background(), dnsName)
	if err != nil{
		log.Warn().Err(err).Msg("cannot resolve IP")
		return nil, err
	}
	if len(ip) > 0{
		ip := ip[0].IP.String()
		return &ip, nil
	}
	return nil, errors.New("empty result")
}

func (m*Manager) getAgentClient(edgeControllerID string) (grpc_edge_controller_go.AgentClient, derrors.Error){
	dnsName := fmt.Sprintf("%s-vpn.service.nalej", edgeControllerID)
	ecIP, err := m.getIP(dnsName)
	if err != nil{
		return nil, derrors.AsError(err, "cannot resolve IP address for EC")
	}
	ipAddr := fmt.Sprintf("%s:5577", *ecIP)
	log.Debug().Str("DNS", dnsName).Str("IP", ipAddr).Msg("Creating agent client")
	conn, err := grpc.Dial(ipAddr, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with edge controller")
	}
	client := grpc_edge_controller_go.NewAgentClient(conn)
	return client, nil
}

func (m*Manager) CreateAgentJoinToken(edgeControllerID *grpc_inventory_go.EdgeControllerId) (*grpc_inventory_manager_go.AgentJoinToken, error) {
	edgeClient, aErr := m.getAgentClient(edgeControllerID.EdgeControllerId)
	if aErr != nil{
		return nil, conversions.ToGRPCError(aErr)
	}
	ctx, cancel := context.WithTimeout(context.Background(), ControllerTimeout)
	defer cancel()
	token, err := edgeClient.CreateAgentJoinToken(ctx, edgeControllerID)
	if err != nil{
		return nil, err
	}
	log.Debug().Str("token", token.Token).Msg("agent join token has been created")
	return token, nil
}

func (m*Manager) TriggerAgentOperation(request *grpc_inventory_manager_go.AgentOpRequest) (*grpc_inventory_manager_go.AgentOpResponse, error) {
	panic("implement me")
}

func (m*Manager) Configure(request *grpc_inventory_manager_go.ConfigureEICRequest) (*grpc_common_go.Success, error) {
	panic("implement me")
}

func (m*Manager) ListMetrics(selector *grpc_inventory_manager_go.AssetSelector) (*grpc_inventory_manager_go.MetricsList, error) {
	panic("implement me")
}

func (m*Manager) QueryMetrics(request *grpc_inventory_manager_go.QueryMetricsRequest) (*grpc_inventory_manager_go.QueryMetricsResult, error) {
	panic("implement me")
}
