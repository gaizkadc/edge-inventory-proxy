/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package commands

import (
	"github.com/nalej/edge-inventory-proxy/internal/pkg/server/eip-sidecar"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var cfg = eip_sidecar.Config{}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Launch EIP",
	Long:  `Launch EIP`,
	Run: func(cmd *cobra.Command, args []string) {
		SetupLogging()
		log.Info().Msg("Launching gRPC EIP!")

		server := eip_sidecar.NewService(cfg)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVar(&cfg.VpnAddress, "vpnAddress", "vpn-server.nalej:5555", " VPN Server internal address with port")
}