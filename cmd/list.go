package cmd

import (
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ll"},
	Short:   "List all the memos",
	Long:    "List all the memos",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var memos = fetchMemos()
		render(memos)

	},
}

func init() {
	rootCmd.AddCommand(list)
}
