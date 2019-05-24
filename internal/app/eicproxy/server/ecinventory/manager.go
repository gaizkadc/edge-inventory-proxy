/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package ecinventory

import (
	"context"
	"github.com/nalej/derrors"
	"github.com/nalej/edge-inventory-proxy/internal/app/eicproxy/config"
	"github.com/nalej/grpc-inventory-go"
	"github.com/nalej/grpc-inventory-manager-go"
	"github.com/nalej/nalej-bus/pkg/queue/inventory/events"
	"time"
)

const defaultTimeout = time.Second * 10

// Manager structure with the entities involved in the management of VPN users
type Manager struct {
	config config.Config
	inventoryProducer *events.InventoryEventsProducer
}

func NewManager(config config.Config, producer *events.InventoryEventsProducer) Manager{
	return Manager{
		config: config,
		inventoryProducer: producer,
	}
}

func (m * Manager) EICStart(request *grpc_inventory_manager_go.EICStartInfo) derrors.Error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	return m.inventoryProducer.Send(ctx, request)
}

func (m *Manager) EICAlive(id *grpc_inventory_go.EdgeControllerId) derrors.Error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	return m.inventoryProducer.Send(ctx, id)
}