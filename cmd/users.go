package cmd

import (
	"github.com/mrauer/gitdump/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(usersCmd)
	usersCmd.AddCommand(usersListCmd)
	usersCmd.AddCommand(usersGetCmd)
	usersCmd.AddCommand(usersDumpCmd)
}

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Users scope",
	Long: `
Public repositories commands:
  gitdump users ls
  gitdump users get [user] [repo]
  gitdump users dump [user]`,
}

var usersListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List public repositories.",
	Long:  `List the public repositories from a public user account`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.ListPublicRepositories(args)
	},
}

var usersGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Download a public repository",
	Long:  `Download a public repository from a public user account`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.GetPublicRepository(args)
	},
}

var usersDumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump all public repositories",
	Long:  `Download all public repositories from a public user account`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.DumpPublicRepositories(args)
	},
}
