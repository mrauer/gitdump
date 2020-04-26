package cmd

import (
	"fmt"
	"github.com/mrauer/gitdump/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(usersCmd)
}

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "TBD",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			lib.GetUserRepos(args[0])
		} else {
			fmt.Println("Usage: gitdump users USERNAME <ACTION>")
		}
	},
}
