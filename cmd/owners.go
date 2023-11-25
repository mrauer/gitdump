package cmd

import (
	"log"

	"github.com/mrauer/gitdump/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ownersCmd)
	ownersCmd.AddCommand(ownersListCmd)
	ownersCmd.AddCommand(ownersGetCmd)
	ownersCmd.AddCommand(ownersDumpCmd)
}

var ownersCmd = &cobra.Command{
	Use:   "owners",
	Short: "Owners scope",
	Long: `
Private repositories commands:
  gitdump owners ls
  gitdump owners get [owner] [repo]
  gitdump owners dump [owner]`,
}

var ownersListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List private repositories.",
	Long:  `List the private repositories of the authenticated token owner`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Listing private repositories...")
		lib.ListPrivateRepositories()
	},
}

var ownersGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Download a private repository",
	Long:  `Download a private repository from the authenticated token owner`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Downloading private repository...")
		lib.GetPrivateRepository(args)
		log.Println("Download completed.")
	},
}

var ownersDumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump all repositories",
	Long:  `Download all private repositories from the authenticated token owner`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Dumping all private repositories...")
		lib.DumpPrivateRepositories(args)
		log.Println("Dump completed.")
	},
}
