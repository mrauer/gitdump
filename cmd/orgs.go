package cmd

import (
	"log"

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
			log.Println("Listing organizations...")
			lib.ListOrganizations(args)
		} else if len(args) == 1 {
			log.Printf("Listing repositories for organization '%s'...", args[0])
			lib.ListOrganizationRepositories(args)
		}
	},
}

var orgsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Download an organization repository",
	Long:  `Download an organization repository granted to the auth token owner`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Downloading organization repository...")
		lib.GetOrganizationRepository(args)
		log.Println("Download completed.")
	},
}

var orgsDumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump organization repositories",
	Long:  `Download all organization repositories granted to the auth token owner`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Dumping all organization repositories...")
		lib.DumpOrganizationRepositories(args)
		log.Println("Dump completed.")
	},
}
