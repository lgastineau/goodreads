package commands

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sortCmd)
}

var sortCmd = &cobra.Command{
	Use:       "--sort",
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ValidArgs: []string{"author", "title"},
	Short:     "Sorts the results by the specified field, if no sort is specified, title is the default",
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "author":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "title":
			cmd.Root().GenBashCompletion(os.Stdout)
		}
	},
}
