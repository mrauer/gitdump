package cmd

import (
	"fmt"
	"github.com/mrauer/gitdump/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(usersCmd)
	usersCmd.AddCommand(usersGetCmd)
	usersCmd.AddCommand(usersDumpCmd)
	usersCmd.AddCommand(usersPrivateCmd)
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
			lib.GetPublicRepository(args)
		} else {
			fmt.Println("\nUsage: gitdump users get <USERNAME> <REPOSITORY>\n")
		}
	},
}

var usersDumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump public repositories",
	Long:  `Dump all the public repositories of a given account`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			lib.GetPublicRepositories(args[0])
		} else {
			fmt.Println("\nUsage: gitdump users dump <USERNAME>\n")
		}
	},
}

var usersPrivateCmd = &cobra.Command{
	Use:   "private",
	Short: "List private repositories",
	Long:  `List private repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.GetUsersPrivateRepository(args)
	},
}
