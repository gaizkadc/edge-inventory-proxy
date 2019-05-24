/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package entities

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-inventory-manager-go"
	"github.com/nalej/grpc-inventory-go"
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
