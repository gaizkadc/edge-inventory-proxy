/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package commands

import (
	"github.com/nalej/edge-inventory-proxy/internal/pkg/config"
	"github.com/nalej/edge-inventory-proxy/internal/pkg/server/edge-inventory-proxy"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var cfg = config.Config{}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Launch EIP",
	Long:  `Launch EIP`,
	Run: func(cmd *cobra.Command, args []string) {
		SetupLogging()
		log.Info().Msg("Launching gRPC EIP!")

		server := edge_inventory_proxy.NewService(cfg)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVar(&cfg.EipPort, "port", 5544, "Port to receive management communications")
}