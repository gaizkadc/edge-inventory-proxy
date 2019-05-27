/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package ecinventory

import (
	"github.com/nalej/edge-inventory-proxy/internal/pkg/entities"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-inventory-go"
	"github.com/nalej/grpc-inventory-manager-go"
	"github.com/nalej/grpc-utils/pkg/conversions"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

type Handler struct {
	manager Manager
}

func (h *Handler) AgentJoin(context.Context, *grpc_inventory_manager_go.AgentJoinRequest) (*grpc_inventory_manager_go.AgentJoinResponse, error) {
	panic("implement me")
}

func NewHandler(manager Manager) *Handler {
	return &Handler{
		manager,
	}
}

func (h *Handler) EICStart(ctx context.Context, request *grpc_inventory_manager_go.EICStartInfo) (*grpc_common_go.Success, error) {
	vErr := entities.ValidEICStartInfo(request)
	if vErr != nil{
		return nil, conversions.ToGRPCError(vErr)
	}
	log.Debug().Str("organization_id", request.OrganizationId).Str("edge_controller_id", request.EdgeControllerId).Str("ip", request.Ip).Msg("EC starts")
	err := h.manager.EICStart(request)
	if err != nil{
		return nil, conversions.ToGRPCError(err)
	}
	return &grpc_common_go.Success{}, nil
}

func (h *Handler) EICAlive(ctx context.Context, id *grpc_inventory_go.EdgeControllerId) (*grpc_common_go.Success, error) {
	vErr := entities.ValidEdgeControllerId(id)
	if vErr != nil{
		return nil, conversions.ToGRPCError(vErr)
	}
	log.Debug().Str("organization_id", id.OrganizationId).Str("edge_controller_id", id.EdgeControllerId).Msg("EC is alive")
	err := h.manager.EICAlive(id)
	if err != nil{
		return nil, conversions.ToGRPCError(err)
	}
	return &grpc_common_go.Success{}, nil
}

