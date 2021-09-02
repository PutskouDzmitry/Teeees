package cmd

import (
	"github.com/spf13/cobra"
)

var timescaleCmd *cobra.Command

func init() {
	timescaleCmd = &cobra.Command{
		Use:   "timescale",
		Short: "run Timescale application version",
		Long:  "run Timescale application version which expects database on host=timescale and port=5432",
		Run: func(cmd *cobra.Command, args []string) {
			// runs server with database located on host=timescale
			start("timescale")
		},
	}
	rootCmd.AddCommand(timescaleCmd)
}
