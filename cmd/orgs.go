package cmd

import (
	"github.com/mrauer/gitdump/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(orgsCmd)
	orgsCmd.AddCommand(orgsListCmd)
	orgsCmd.AddCommand(orgsGetCmd)
	orgsCmd.AddCommand(orgsDumpCmd)
}

var orgsCmd = &cobra.Command{
	Use:   "orgs",
	Short: "Orgs scope",
	Long: `
Orgs repositories commands:
  gitdump orgs ls
  gitdump orgs ls [org]
  gitdump orgs get [org] [repo]
  gitdump orgs dump [org]`,
}

var orgsListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List organizations/repositories.",
	Long:  `List the organizations granted to the auth token owner`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			lib.ListOrganizations(args)
		}
		if len(args) == 1 {
			lib.ListOrganizationRepositories(args)
		}
	},
}

var orgsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Download an organization repository",
	Long:  `Download an organization repository granted to the auth token owner`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.GetOrganizationRepository(args)
	},
}

var orgsDumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump organization repositories",
	Long:  `Download all organization repositories granted to the auth token owner`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.DumpOrganizationRepositories(args)
	},
}
