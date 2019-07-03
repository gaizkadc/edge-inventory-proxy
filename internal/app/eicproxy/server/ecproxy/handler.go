/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package ecproxy

import (
	"context"
	"github.com/nalej/derrors"
	"github.com/nalej/edge-inventory-proxy/internal/pkg/entities"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-inventory-go"
	"github.com/nalej/grpc-inventory-manager-go"
	"github.com/nalej/grpc-utils/pkg/conversions"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	manager Manager
}

func NewHandler(manager Manager) *Handler {
	return &Handler{
		manager,
	}
}


func (h *Handler) InstallAgent(_ context.Context, request *grpc_inventory_manager_go.InstallAgentRequest) (*grpc_inventory_manager_go.EdgeControllerOpResponse, error) {
	verr := entities.ValidInstallAgentRequest(request)
	if verr != nil {
		return nil, conversions.ToGRPCError(verr)
	}
	return h.manager.InstallAgent(request)
}

func (h *Handler) CreateAgentJoinToken(_ context.Context, edgeControllerID *grpc_inventory_go.EdgeControllerId) (*grpc_inventory_manager_go.AgentJoinToken, error) {
	verr := entities.ValidEdgeControllerId(edgeControllerID)
	if verr != nil {
		return nil, conversions.ToGRPCError(verr)
	}
	return h.manager.CreateAgentJoinToken(edgeControllerID)
}

func (h*Handler) TriggerAgentOperation(_ context.Context, request *grpc_inventory_manager_go.AgentOpRequest) (*grpc_inventory_manager_go.AgentOpResponse, error) {
	vErr := entities.ValidAgentOpRequest(request)
	if vErr != nil {
		return nil, conversions.ToGRPCError(vErr)
	}
	return h.manager.TriggerAgentOperation(request)
}

func (h*Handler) Configure(_ context.Context, request *grpc_inventory_manager_go.ConfigureEICRequest) (*grpc_common_go.Success, error) {
	return nil, conversions.ToGRPCError(derrors.NewUnimplementedError("Configure call not implemented"))
}

func (h*Handler) ListMetrics(_ context.Context, selector *grpc_inventory_manager_go.AssetSelector) (*grpc_inventory_manager_go.MetricsList, error) {
	derr := entities.ValidAssetSelector(selector)
	if derr != nil {
		return nil, conversions.ToGRPCError(derr)
	}
	return h.manager.ListMetrics(selector)
}

func (h*Handler) QueryMetrics(_ context.Context, request *grpc_inventory_manager_go.QueryMetricsRequest) (*grpc_inventory_manager_go.QueryMetricsResult, error) {
	derr := entities.ValidQueryMetricsRequest(request)
	if derr != nil {
		return nil, conversions.ToGRPCError(derr)
	}
	return h.manager.QueryMetrics(request)
}

// UnlinkEC operation to remove/uninstall an EIC.
func (h *Handler) UnlinkEC(_ context.Context, edgeControllerID *grpc_inventory_go.EdgeControllerId) (*grpc_common_go.Success, error){
	log.Debug().Msg("UnlinkEIC received")
	vErr := entities.ValidEdgeControllerId(edgeControllerID)
	if vErr != nil {
		return nil, conversions.ToGRPCError(vErr)
	}

	return h.manager.UnlinkEC(edgeControllerID)
}

func (h *Handler) UninstallAgent(_ context.Context, assetID *grpc_inventory_manager_go.FullUninstallAgentRequest) (*grpc_inventory_manager_go.EdgeControllerOpResponse, error) {
	log.Debug().Str("edge_controller_id", assetID.EdgeControllerId).Str("asset_id", assetID.AssetId).Msg("uninstall msg received")

	vErr := entities.ValidFullUninstallAgentRequest(assetID)
	if vErr != nil {
		return nil, conversions.ToGRPCError(vErr)
	}
	return h.manager.UninstallAgent(assetID)

}