package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var str = `
Usage: gitdump SCOPE COMMAND [ARG1] [ARG2]...

A tool for downloading GitHub repositories

Scopes:
  users    Public repositories
  owners   Private repositories (require authentication)
  orgs     Organizations repositories (require authentication)

Commands:
  ls       List the repositories
  get      Download a single repository
  dump     Download all repositories

Args:
  user     Public account
  owner    Account associated with the the token
  org      Organization name
  repo     Repository name
`

var rootCmd = &cobra.Command{
	Use:   "gitdump",
	Short: "A tool for downloading GitHub repositories",
	Long:  `A tool for downloading GitHub repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(str)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
