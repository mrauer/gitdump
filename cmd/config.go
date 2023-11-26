package cmd

import (
	"log"

	"github.com/mrauer/gitdump/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configTokenCmd)
	configCmd.AddCommand(configPathCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "GitDump configuration",
	Long: `
Configuration commands:
  gitdump config token [github_token]
  gitdump config path [download_path]`,
}

var configTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Set the GitHub Token",
	Long:  `Add your GitHub Token as an argument`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Setting GitHub Token...")
		lib.ConfigToken(args)
		log.Println("GitHub Token set successfully.")
	},
}

var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Set the path to download files",
	Long:  `This will be where your files will be downloaded`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Setting download path...")
		lib.ConfigPath(args)
		log.Println("Download path set successfully.")
	},
}
