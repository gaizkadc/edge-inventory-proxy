/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package ecinventory

import (
	"github.com/nalej/edge-inventory-proxy/internal/app/eicproxy/config"
)

// Manager structure with the entities involved in the management of VPN users
type Manager struct {
	config config.Config
}

func NewManager(config config.Config) Manager{
	return Manager{
		config: config,
	}
}