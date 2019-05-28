/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package commands

import (
	"github.com/nalej/edge-inventory-proxy/internal/app/sidecar/config"
	"github.com/nalej/edge-inventory-proxy/internal/app/sidecar/server"
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
		log.Info().Msg("Launching sidecar!")

		server := server.NewService(cfg)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVar(&cfg.VPNAddress, "vpnAddress", "localhost:5555", " VPN Server internal address (host:port)")
	runCmd.Flags().StringVar(&cfg.NetworkManagerAddress, "networkManagerAddress", "localhost:8000",
		"Network Manager address (host:port)")
	runCmd.Flags().StringVar(&cfg.ProxyName, "proxyName", "proxy0", "Name of the proxy for the DNS registration without .vpn.service.nalej")
	runCmd.Flags().StringVar(&cfg.Username, "username", "admin", "Username for the VPN")
	runCmd.Flags().StringVar(&cfg.Password, "password", "ecb75e08-27ad-412f-acd0-436f6f7b7c29", "Password for the VPN user")
}