/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

// This is an example of an executable command.

package commands

import (
	"github.com/nalej/edge-inventory-proxy/internal/pkg/config"
	"github.com/nalej/edge-inventory-proxy/internal/pkg/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var cfg = config.Config{}

var helloCmd = &cobra.Command{
	Use:   "run",
	Short: "Launch EIP",
	Long:  `Launch EIP`,
	Run: func(cmd *cobra.Command, args []string) {
		SetupLogging()
		log.Info().Msg("Launching gRPC EIP!")
		cfg.Debug = debugLevel

		cfg.Print()
		err := cfg.Validate()
		if err != nil {
			log.Fatal().Err(err)
		}

		server := server.NewService(cfg)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}