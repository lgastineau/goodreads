package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:     "--search",
	Aliases: []string{"-s"},
	Short:   "Search the Goodreads' API and display the results on screen",
	Args:    cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("search command")
	},
}
