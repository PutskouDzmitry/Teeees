package cmd

import (
	"github.com/spf13/cobra"
)

var postgresCmd *cobra.Command

func init() {
	postgresCmd = &cobra.Command{
		Use:   "postgres",
		Short: "run PostgreSQL application version",
		Long:  "run PostgreSQL application version which expects database on host=postgres and port=5432",
		Run: func(cmd *cobra.Command, args []string) {
			// runs server with database located on host=postgre
			start("postgres")
		},
	}
	rootCmd.AddCommand(postgresCmd)
}
