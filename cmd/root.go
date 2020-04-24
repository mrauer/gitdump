package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var str = `
Usage: gitdump COMMAND

A tool for downloading GitHub repositories

Commands:
  get      Download a single GitHub repository
  dump     Download an entire GitHub account
`

var rootCmd = &cobra.Command{
	Use:   "gitdump",
	Short: "TBD",
	Long:  `TBD`,
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
