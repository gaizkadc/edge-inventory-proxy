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

func NewHandler(manager Manager) *Handler {
	return &Handler{
		manager,
	}
}

// ------------
// Agent
// ------------
func (h *Handler) AgentJoin(_ context.Context, request *grpc_inventory_manager_go.AgentJoinRequest) (*grpc_inventory_manager_go.AgentJoinResponse, error) {
	vErr := entities.ValidAgentJoinRequest(request)
	if vErr != nil {
		return nil, conversions.ToGRPCError(vErr)
	}

	log.Debug().Str("organization_id", request.OrganizationId).Str("edge_controller_id", request.EdgeControllerId).
		Str("agent_id", request.AgentId).Msg("Agent join")

	return h.manager.AgentJoin(request)
}

func (h *Handler)  LogAgentAlive(_ context.Context, request *grpc_inventory_manager_go.AgentsAlive) (*grpc_common_go.Success, error){
	vErr := entities.ValidAgentsAlive(request)
	if vErr != nil {
		return nil, conversions.ToGRPCError(vErr)
	}

	log.Debug().Str("organization_id", request.OrganizationId).Str("edge_controller_id", request.EdgeControllerId).
		Int("agents", len(request.Agents)).Msg("Agents Alive")

	err := h.manager.LogAgentAlive(request)
	if err != nil{
		return nil, conversions.ToGRPCError(err)
	}
	return &grpc_common_go.Success{}, nil
}

// CallbackAgentOperation is called by the EIC upon execution of the operation by the agent.
func (h *Handler) CallbackAgentOperation(_ context.Context, opResponse *grpc_inventory_manager_go.AgentOpResponse) (*grpc_common_go.Success, error) {
	vErr := entities.ValidAgentOpResponse(opResponse)
	if  vErr != nil {
		return nil, conversions.ToGRPCError(vErr)
	}

	log.Debug().Str("organization_id", opResponse.OrganizationId).Str("edge_controller_id", opResponse.EdgeControllerId).
		Str("asset_id", opResponse.AssetId).Str("operation_id", opResponse.OperationId).Msg("CallbackAgentOperation")


	err := h.manager.CallbackAgentOperation(opResponse)
	if err != nil {
		return nil, err
	}

	return &grpc_common_go.Success{}, nil

}

func (h *Handler) AgentUninstalled(_ context.Context, assetId *grpc_inventory_go.AssetUninstalledId) (*grpc_common_go.Success, error) {

	vErr := entities.ValidAssetUninstalledId(assetId)
	if vErr != nil {
		return nil, conversions.ToGRPCError(vErr)
	}
	err := h.manager.AgentUninstalled(assetId)
	if err != nil {
		return nil, err
	}

	return &grpc_common_go.Success{}, nil
}


// ----------------
// Edge Controller
// ----------------
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
	err := h.manager.EICAlive(id)
	if err != nil{
		return nil, conversions.ToGRPCError(err)
	}
	return &grpc_common_go.Success{}, nil
}


func (h *Handler) CallbackECOperation(ctx context.Context, response *grpc_inventory_manager_go.EdgeControllerOpResponse) (*grpc_common_go.Success, error) {
	vErr := entities.ValidEdgeControllerOpResponse(response)
	if vErr != nil{
		return nil, conversions.ToGRPCError(vErr)
	}
	err := h.manager.CallbackECOperation(response)
	if err != nil{
		return nil, conversions.ToGRPCError(err)
	}
	return &grpc_common_go.Success{}, nil
}