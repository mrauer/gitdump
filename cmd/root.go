package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var str = `
Usage: gitdump MODEL COMMAND [ARG1] [ARG2]...

A tool for downloading GitHub repositories

Models:
  users    Public repositories
  orgs     Private repositories (require authentication)

Commands:
  describe   Display information about the repository
  get        Download a copy of the repository

Args:
  slug          Unique account identifier
  repository    ID of the repository
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
