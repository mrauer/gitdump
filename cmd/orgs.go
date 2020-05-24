package cmd

import (
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
		lib.GetPrivateRepository(args)
	},
}
