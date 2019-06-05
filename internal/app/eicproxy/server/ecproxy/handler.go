/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package ecproxy

import (
	"context"
	"github.com/nalej/edge-inventory-proxy/internal/pkg/entities"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-inventory-go"
	"github.com/nalej/grpc-inventory-manager-go"
	"github.com/nalej/grpc-utils/pkg/conversions"
)

type Handler struct {
	manager Manager
}

func NewHandler(manager Manager) *Handler {
	return &Handler{
		manager,
	}
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
	panic("implement me")
}

func (h*Handler) ListMetrics(_ context.Context, selector *grpc_inventory_manager_go.AssetSelector) (*grpc_inventory_manager_go.MetricsList, error) {
	panic("implement me")
}

func (h*Handler) QueryMetrics(_ context.Context, request *grpc_inventory_manager_go.QueryMetricsRequest) (*grpc_inventory_manager_go.QueryMetricsResult, error) {
	panic("implement me")
}

