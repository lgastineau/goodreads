package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(hostCmd)
}

var hostCmd = &cobra.Command{
	Use:     "--host",
	Aliases: []string{"-h"},
	Short:   "the hostname or ip address where the server can be found. Default is 127.0.0.1",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("host command")
	},
}
