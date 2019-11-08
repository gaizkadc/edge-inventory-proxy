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
