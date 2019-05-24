/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package ecproxy

import (
	"context"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-inventory-manager-go"
)

type Handler struct {
	manager Manager
}

func NewHandler(manager Manager) *Handler {
	return &Handler{
		manager,
	}
}

func (h*Handler) TriggerAgentOperation(ctx context.Context, request *grpc_inventory_manager_go.AgentOpRequest) (*grpc_inventory_manager_go.AgentOpResponse, error) {
	panic("implement me")
}

func (h*Handler) Configure(ctx context.Context, request *grpc_inventory_manager_go.ConfigureEICRequest) (*grpc_common_go.Success, error) {
	panic("implement me")
}

func (h*Handler) ListMetrics(ctx context.Context, selector *grpc_inventory_manager_go.AssetSelector) (*grpc_inventory_manager_go.MetricsList, error) {
	panic("implement me")
}

func (h*Handler) QueryMetrics(ctx context.Context, request *grpc_inventory_manager_go.QueryMetricsRequest) (*grpc_inventory_manager_go.QueryMetricsResult, error) {
	panic("implement me")
}
