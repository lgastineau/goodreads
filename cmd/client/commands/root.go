package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goodreads",
	Short: "Goodreads is ____",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
