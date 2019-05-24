/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package ecinventory

import (
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-inventory-go"
	"github.com/nalej/grpc-inventory-manager-go"
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

func (*Handler) EICStart(context.Context, *grpc_inventory_manager_go.EICStartInfo) (*grpc_common_go.Success, error) {
	panic("implement me")
}

func (*Handler) EICAlive(context.Context, *grpc_inventory_go.EdgeControllerId) (*grpc_common_go.Success, error) {
	panic("implement me")
}

