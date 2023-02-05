package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pageCmd)
}

var pageCmd = &cobra.Command{
	Use:   "-p",
	Short: "display the NUMBER page of results",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("pagination command")
	},
}
