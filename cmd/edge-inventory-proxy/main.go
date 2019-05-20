/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package main

import (
	"github.com/nalej/edge-inventory-proxy/cmd/edge-inventory-proxy/commands"
	"github.com/nalej/edge-inventory-proxy/version"
)

var MainVersion string

var MainCommit string

func main() {
	version.AppVersion = MainVersion
	version.Commit = MainCommit
	commands.Execute()
}
