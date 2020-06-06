package cmd

import (
	"fmt"
	"github.com/mrauer/gitdump/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(orgsCmd)
	orgsCmd.AddCommand(orgsGetCmd)
	orgsCmd.AddCommand(orgsDumpCmd)
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
		lib.GetOrgsPrivateRepository(args)
	},
}

var orgsDumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump private repositories",
	Long:  `Dump all the private repositories of a given organization`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			lib.GetPrivateRepositories(args[0])
		} else {
			fmt.Println("\nUsage: gitdump orgs dump <ORGANIZATION>\n")
		}
	},
}
