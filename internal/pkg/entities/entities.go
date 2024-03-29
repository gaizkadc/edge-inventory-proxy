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

package entities

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-inventory-go"
	"github.com/nalej/grpc-inventory-manager-go"
	"github.com/nalej/grpc-monitoring-go"
)

func ValidEdgeControllerId(id *grpc_inventory_go.EdgeControllerId) derrors.Error {
	if id.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if id.EdgeControllerId == "" {
		return derrors.NewInvalidArgumentError("edge_controller_id cannot be empty")
	}
	return nil
}

func ValidEdgeControllerOpResponse(response *grpc_inventory_manager_go.EdgeControllerOpResponse) derrors.Error {
	if response.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if response.EdgeControllerId == "" {
		return derrors.NewInvalidArgumentError("edge_controller_id cannot be empty")
	}
	if response.OperationId == "" {
		return derrors.NewInvalidArgumentError("operation_id cannot be empty")
	}
	return nil
}

func ValidEICStartInfo(info *grpc_inventory_manager_go.EICStartInfo) derrors.Error {
	if info.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if info.EdgeControllerId == "" {
		return derrors.NewInvalidArgumentError("edge_controller_id cannot be empty")
	}
	if info.Ip == "" {
		return derrors.NewInvalidArgumentError("ip cannot be empty")
	}
	return nil
}

func ValidInstallAgentRequest(request *grpc_inventory_manager_go.InstallAgentRequest) derrors.Error {
	if request.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if request.EdgeControllerId == "" {
		return derrors.NewInvalidArgumentError("edge_controller_id cannot be empty")
	}
	if request.TargetHost == "" {
		return derrors.NewInvalidArgumentError("target_host cannot be empty")
	}
	return nil
}

func ValidAgentJoinRequest(request *grpc_inventory_manager_go.AgentJoinRequest) derrors.Error {
	if request.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if request.EdgeControllerId == "" {
		return derrors.NewInvalidArgumentError("edge_controller_id cannot be empty")
	}
	if request.AgentId == "" {
		return derrors.NewInvalidArgumentError("agent_id cannot be empty")
	}
	return nil
}

func ValidAgentsAlive(request *grpc_inventory_manager_go.AgentsAlive) derrors.Error {
	if request.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if request.EdgeControllerId == "" {
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

func ValidAgentOpResponse(request *grpc_inventory_manager_go.AgentOpResponse) derrors.Error {
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

func ValidAssetSelector(selector *grpc_inventory_go.AssetSelector) derrors.Error {
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

func ValidQueryMetricsRequest(request *grpc_monitoring_go.QueryMetricsRequest) derrors.Error {
	// We check the asset selector so we know we have an edge controller ID.
	// The rest is verified by the Edge Controller so we don't have to
	// adapt the proxy if functionality changes.
	derr := ValidAssetSelector(request.GetAssets())
	if derr != nil {
		return derr
	}

	return nil
}

func ValidFullUninstallAgentRequest(request *grpc_inventory_manager_go.FullUninstallAgentRequest) derrors.Error {
	if request.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if request.EdgeControllerId == "" {
		return derrors.NewInvalidArgumentError("edge_controller_id cannot be empty")
	}
	if request.AssetId == "" {
		return derrors.NewInvalidArgumentError("asset_id cannot be empty")
	}
	return nil
}

func ValidAssetUninstalledId(request *grpc_inventory_go.AssetUninstalledId) derrors.Error {
	if request.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("organization_id cannot be empty")
	}
	if request.AssetId == "" {
		return derrors.NewInvalidArgumentError("asset_id cannot be empty")
	}
	return nil
}
