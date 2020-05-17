package cmd

import (
	"fmt"
	"github.com/mrauer/gitdump/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(usersCmd)
	usersCmd.AddCommand(usersGetCmd)
}

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Users model.",
	Long:  `Act on publicly available information about someone with a GitHub account.`,
}

var usersGetCmd = &cobra.Command{
	Use:   "get",
	Short: "TBD",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			lib.GetRepository(args)
		} else {
			fmt.Println("\nUsage: gitdump users get <USERNAME> <REPOSITORY>\n")
		}
	},
}
