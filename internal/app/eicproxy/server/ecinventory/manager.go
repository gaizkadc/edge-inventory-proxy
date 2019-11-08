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

package ecinventory

import (
	"context"
	"github.com/nalej/derrors"
	"github.com/nalej/edge-inventory-proxy/internal/app/eicproxy/config"
	"github.com/nalej/grpc-inventory-go"
	"github.com/nalej/grpc-inventory-manager-go"
	"github.com/nalej/nalej-bus/pkg/queue/inventory/events"
	"github.com/nalej/nalej-bus/pkg/queue/inventory/ops"
	"time"
)

const defaultTimeout = time.Second * 10

// Manager structure with the entities involved in the management of VPN users
type Manager struct {
	config               config.Config
	inventoryProducer    *events.InventoryEventsProducer
	inventoryOpsProducer *ops.InventoryOpsProducer
	agentClient          grpc_inventory_manager_go.AgentClient
}

func NewManager(config config.Config, producer *events.InventoryEventsProducer, opProducer *ops.InventoryOpsProducer, client grpc_inventory_manager_go.AgentClient) Manager {
	return Manager{
		config:               config,
		inventoryProducer:    producer,
		inventoryOpsProducer: opProducer,
		agentClient:          client,
	}
}

// ----------------
// Edge Controller
// ----------------
func (m *Manager) EICStart(request *grpc_inventory_manager_go.EICStartInfo) derrors.Error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	return m.inventoryProducer.Send(ctx, request)
}

func (m *Manager) EICAlive(id *grpc_inventory_go.EdgeControllerId) derrors.Error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	return m.inventoryProducer.Send(ctx, id)
}

func (m *Manager) CallbackECOperation(response *grpc_inventory_manager_go.EdgeControllerOpResponse) derrors.Error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	return m.inventoryOpsProducer.Send(ctx, response)
}

// ------------
// Agent
// ------------
func (m *Manager) AgentJoin(request *grpc_inventory_manager_go.AgentJoinRequest) (*grpc_inventory_manager_go.AgentJoinResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return m.agentClient.AgentJoin(ctx, request)
}

func (m *Manager) LogAgentAlive(request *grpc_inventory_manager_go.AgentsAlive) derrors.Error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	return m.inventoryProducer.Send(ctx, request)

}

func (m *Manager) CallbackAgentOperation(opResponse *grpc_inventory_manager_go.AgentOpResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	return m.inventoryOpsProducer.Send(ctx, opResponse)
}

func (m *Manager) AgentUninstalled(assetId *grpc_inventory_go.AssetUninstalledId) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	return m.inventoryProducer.Send(ctx, assetId)
}
