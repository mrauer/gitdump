package cmd

import (
	"github.com/mrauer/gitdump/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ownersCmd)
	ownersCmd.AddCommand(ownersListCmd)
	ownersCmd.AddCommand(ownersGetCmd)
	ownersCmd.AddCommand(ownersDumpCmd)
}

var ownersCmd = &cobra.Command{
	Use:   "owners",
	Short: "Owners scope",
	Long: `
Private repositories commands:
  gitdump owners ls
  gitdump owners get [owner] [repo]
  gitdump owners dump [owner]`,
}

var ownersListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List private repositories.",
	Long:  `List the private repositories of the auth token owner`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.ListPrivateRepositories()
	},
}

var ownersGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Download a private repository",
	Long:  `Download a private repository from the auth token owner`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.GetPrivateRepository(args)
	},
}

var ownersDumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump all repositories",
	Long:  `Download all private repositories from the auth token owner`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.DumpPrivateRepositories(args)
	},
}
