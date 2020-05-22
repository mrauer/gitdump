package cmd

import (
	"fmt"
	"github.com/mrauer/gitdump/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(orgsCmd)
	orgsCmd.AddCommand(orgsGetCmd)
}

var orgsCmd = &cobra.Command{
	Use:   "orgs",
	Short: "Orgs model.",
	Long:  `TBD`,
}

var orgsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "TBD",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			lib.GetPrivateRepository()
		} else {
			fmt.Println("\nUsage: gitdump orgs get <USERNAME> <REPOSITORY>\n")
		}
	},
}
