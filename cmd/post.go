package cmd

import (
	"github.com/spf13/cobra"
)

var lis2 = &cobra.Command{
	Use:     "post",
	Aliases: []string{"p"},
	Short:   "Post new moemos",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		post(args[0])
	},
}

func init() {
	rootCmd.AddCommand(lis2)
}
