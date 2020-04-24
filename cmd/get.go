package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(tbdCmd)
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "TBD",
	Long:  `TBD`,
}

var tbdCmd = &cobra.Command{
	Use:   "tbd",
	Short: "TBD",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("my command here")
	},
}
