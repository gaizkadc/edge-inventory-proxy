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
	"github.com/nalej/grpc-monitoring-go"
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
			return d.DialContext(ctx, "udp", net.JoinHostPort(m.config.DNSServer, "53"))
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

func (m*Manager) getEICClient(edgeControllerID string) (grpc_edge_controller_go.EICClient, *grpc.ClientConn, derrors.Error){
	dnsName := fmt.Sprintf("%s-vpn.service.nalej", edgeControllerID)
	ecIP, err := m.getIP(dnsName)
	if err != nil{
		return nil, nil, derrors.AsError(err, "cannot resolve IP address for EC")
	}
	ipAddr := fmt.Sprintf("%s:5577", *ecIP)
	log.Debug().Str("DNS", dnsName).Str("IP", ipAddr).Msg("Creating agent client")
	conn, err := grpc.Dial(ipAddr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, derrors.AsError(err, "cannot create connection with edge controller")
	}
	client := grpc_edge_controller_go.NewEICClient(conn)
	return client, conn, nil
}

func (m*Manager) InstallAgent(request *grpc_inventory_manager_go.InstallAgentRequest) (*grpc_inventory_manager_go.EdgeControllerOpResponse, error) {
	edgeClient, conn, aErr := m.getEICClient(request.EdgeControllerId)
	if aErr != nil{
		return nil, conversions.ToGRPCError(aErr)
	}
	if conn != nil {
		defer conn.Close()
	}
	ctx, cancel := context.WithTimeout(context.Background(), ControllerTimeout)
	defer cancel()
	response, err := edgeClient.InstallAgent(ctx, request)
	if err != nil{
		return nil, m.ConvertError(err, "install agent")
	}
	log.Debug().Interface("response", response).Msg("install agent request sent")
	return response, nil
}

func (m*Manager) CreateAgentJoinToken(edgeControllerID *grpc_inventory_go.EdgeControllerId) (*grpc_inventory_manager_go.AgentJoinToken, error) {
	edgeClient, conn, aErr := m.getEICClient(edgeControllerID.EdgeControllerId)
	if aErr != nil{
		return nil, conversions.ToGRPCError(aErr)
	}
	if conn != nil {
		defer conn.Close()
	}
	ctx, cancel := context.WithTimeout(context.Background(), ControllerTimeout)
	defer cancel()
	token, err := edgeClient.CreateAgentJoinToken(ctx, edgeControllerID)
	if err != nil{
		return nil, m.ConvertError(err, "create agent-join-token")
	}
	log.Debug().Str("token", token.Token).Msg("agent join token has been created")
	return token, nil
}

func (m*Manager) TriggerAgentOperation(request *grpc_inventory_manager_go.AgentOpRequest) (*grpc_inventory_manager_go.AgentOpResponse, error) {
	edgeClient, conn, aErr := m.getEICClient(request.EdgeControllerId)
	if aErr != nil{
		return nil, conversions.ToGRPCError(aErr)
	}
	if conn != nil {
		defer conn.Close()
	}
	ctx, cancel := context.WithTimeout(context.Background(), ControllerTimeout)
	defer cancel()
	res, err :=  edgeClient.TriggerAgentOperation(ctx, request)
	if err != nil {
		return nil, m.ConvertError(err, "trigger agent operation")
	}
	return res, nil
}

func (m*Manager) Configure(request *grpc_inventory_manager_go.ConfigureEICRequest) (*grpc_common_go.Success, error) {
	return nil, conversions.ToGRPCError(derrors.NewUnimplementedError("configure call not implemented"))
}

func (m*Manager) ListMetrics(selector *grpc_inventory_go.AssetSelector) (*grpc_monitoring_go.MetricsList, error) {
	edgeClient, conn, aErr := m.getEICClient(selector.GetEdgeControllerId())
	if aErr != nil{
		return nil, conversions.ToGRPCError(aErr)
	}
	if conn != nil {
		defer conn.Close()
	}
	ctx, cancel := context.WithTimeout(context.Background(), ControllerTimeout)
	defer cancel()
	metrics, err :=  edgeClient.ListMetrics(ctx, selector)
	if err != nil {
		return nil, m.ConvertError(err, "list metrics")
	}
	return metrics, nil
}

func (m*Manager) QueryMetrics(request *grpc_monitoring_go.QueryMetricsRequest) (*grpc_monitoring_go.QueryMetricsResult, error) {
	edgeClient, conn, aErr := m.getEICClient(request.GetAssets().GetEdgeControllerId())
	if aErr != nil{
		return nil, conversions.ToGRPCError(aErr)
	}
	if conn != nil {
		defer conn.Close()
	}
	ctx, cancel := context.WithTimeout(context.Background(), ControllerTimeout)
	defer cancel()
	metrics, err := edgeClient.QueryMetrics(ctx, request)
	if err != nil {
		return nil, m.ConvertError(err, "query metrics")
	}
	return metrics, nil
}

func (m *Manager) UnlinkEC(edge *grpc_inventory_go.EdgeControllerId) (*grpc_common_go.Success, error){
	log.Debug().Msg("UnlinkEIC received")

	edgeClient, conn, aErr := m.getEICClient(edge.EdgeControllerId)
	if aErr != nil{
		return nil, conversions.ToGRPCError(aErr)
	}
	if conn != nil {
		defer conn.Close()
	}
	ctx, cancel := context.WithTimeout(context.Background(), ControllerTimeout)
	defer cancel()
	res, err :=  edgeClient.Unlink(ctx, &grpc_common_go.Empty{})
	if err != nil {
		return nil, m.ConvertError(err, "unlink")
	}else {
		return res, err
	}
}

func (m *Manager) UninstallAgent(assetID *grpc_inventory_manager_go.FullUninstallAgentRequest) (*grpc_inventory_manager_go.EdgeControllerOpResponse, error) {
	edgeClient, conn, aErr := m.getEICClient(assetID.EdgeControllerId)
	if aErr != nil{
		return nil, conversions.ToGRPCError(aErr)
	}
	if conn != nil {
		defer conn.Close()
	}
	ctx, cancel := context.WithTimeout(context.Background(), ControllerTimeout)
	defer cancel()

	res, err := edgeClient.UninstallAgent(ctx, assetID)

	if err != nil {
		return nil, m.ConvertError(err, "uninstall agent")
	}else {
		return res, err
	}

}

func (m *Manager) ConvertError(err error, msg string)  error {

	if conversions.ToDerror(err).Type() == derrors.Unavailable ||
		conversions.ToDerror(err).Type() == derrors.DeadlineExceeded{
		return conversions.ToGRPCError(derrors.NewUnavailableError(
			fmt.Sprintf("unable to %s, EC not connected", msg), err))
	}else{
		return err
	}

}
