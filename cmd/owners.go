package cmd

import (
	"github.com/mrauer/gitdump/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ownersCmd)
	ownersCmd.AddCommand(ownersListCmd)
}

var ownersCmd = &cobra.Command{
	Use:   "owners",
	Short: "Owners scope",
	Long: `Private repositories commands:
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
