/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package entities

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-inventory-go"
	"github.com/nalej/grpc-inventory-manager-go"
)

func ValidEdgeControllerId(id * grpc_inventory_go.EdgeControllerId) derrors.Error{
	if id.OrganizationId == ""{
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if id.EdgeControllerId == ""{
		return derrors.NewInvalidArgumentError("edge_controller_id cannot be empty")
	}
	return nil
}

func ValidEICStartInfo(info * grpc_inventory_manager_go.EICStartInfo) derrors.Error{
	if info.OrganizationId == ""{
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if info.EdgeControllerId == ""{
		return derrors.NewInvalidArgumentError("edge_controller_id cannot be empty")
	}
	if info.Ip == ""{
		return derrors.NewInvalidArgumentError("ip cannot be empty")
	}
	return nil
}

func ValidAgentJoinRequest (request *grpc_inventory_manager_go.AgentJoinRequest) derrors.Error {
	if request.OrganizationId == ""{
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if request.EdgeControllerId == ""{
		return derrors.NewInvalidArgumentError("edge_controller_id cannot be empty")
	}
	if request.AgentId == "" {
		return derrors.NewInvalidArgumentError("agent_id cannot be empty")
	}
	return nil
}

func ValidAgentsAlive (request *grpc_inventory_manager_go.AgentsAlive) derrors.Error {
	if request.OrganizationId == ""{
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if request.EdgeControllerId == ""{
		return derrors.NewInvalidArgumentError("edge_controller_id cannot be empty")
	}
	if request.Agents == nil || len(request.Agents) <= 0 {
		return derrors.NewInvalidArgumentError("agents cannot be empty")
	}
	return nil
}


func ValidAgentOpRequest(request *grpc_inventory_manager_go.AgentOpRequest) derrors.Error {
	if request.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if request.EdgeControllerId == "" {
		return derrors.NewInvalidArgumentError("edge_controller_id cannot be empty")
	}
	if request.AssetId == "" {
		return derrors.NewInvalidArgumentError("asset_id cannot be empty")
	}
	if request.OperationId == "" {
		return derrors.NewInvalidArgumentError("operation_id cannot be empty")
	}
	if request.Plugin == "" {
		return derrors.NewInvalidArgumentError("plugin cannot be empty")
	}
	return nil
}

func ValidAgentOpResponse (request *grpc_inventory_manager_go.AgentOpResponse) derrors.Error {
	if request.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if request.EdgeControllerId == "" {
		return derrors.NewInvalidArgumentError("edge_controller_id cannot be empty")
	}
	if request.AssetId == "" {
		return derrors.NewInvalidArgumentError("asset_id cannot be empty")
	}
	if request.OperationId == "" {
		return derrors.NewInvalidArgumentError("operation_id cannot be empty")
	}
	return nil
}

func ValidAssetSelector(selector *grpc_inventory_manager_go.AssetSelector) derrors.Error {
	// For the proxy to properly connect we need an Edge Controller ID.
	// Other validation is done by the Edge Controller.
	if selector == nil {
		return derrors.NewInvalidArgumentError("empty asset selector")
	}
	if selector.GetOrganizationId() == "" {
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if selector.GetEdgeControllerId() == "" {
		return derrors.NewInvalidArgumentError("edge_controller_id cannot be empty")
	}
	return nil
}

func ValidQueryMetricsRequest(request *grpc_inventory_manager_go.QueryMetricsRequest) derrors.Error {
	// We check the asset selector so we know we have an edge controller ID.
	// The rest is verified by the Edge Controller so we don't have to
	// adapt the proxy if functionality changes.
	derr := ValidAssetSelector(request.GetAssets())
	if derr != nil {
		return derr
	}

	return nil
}
