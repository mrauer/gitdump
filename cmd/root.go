package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var str = `
Usage: gitdump MODEL [OWNER] COMMAND

A tool for downloading GitHub repositories

Models:
  users    Public repositories
  orgs     Private repositories

Commands:
  dump     Download the repositories

Owner: account name
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
