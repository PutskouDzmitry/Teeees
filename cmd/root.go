package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "load-testing",
	Short: "Golang application for load testing PostgreSQL or Timescale databases",
	Long: `Golang application for load testing PostgreSQL or Timescale databases. It uses:
  - host: postgres/timescale;
  - port: 5432;
  - user: postgres;
  - password: root;
  - database: test;
  - ssl-mode: disable. 
You can change it using environment variables. (see Docker-compose.yaml)`,
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
