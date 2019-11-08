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
	"github.com/nalej/edge-inventory-proxy/internal/app/eicproxy/config"
	"github.com/nalej/edge-inventory-proxy/internal/app/eicproxy/server"
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
		cfg.Debug = debugLevel
		server := server.NewService(cfg)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVar(&cfg.EipPort, "port", 5544, "Port to receive management communications")
	runCmd.Flags().StringVar(&cfg.QueueAddress, "queueAddress", "localhost:6650", "Queue address (host:port)")
	runCmd.Flags().StringVar(&cfg.DNSServer, "dnsServer", "", "DNS server to resolve EC hostnames")
	runCmd.Flags().StringVar(&cfg.InventoryManagerAddress, "inventoryManagerAddress", "localhost:5510", "Inventory Manager address (host:port)")
}