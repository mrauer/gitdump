package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

const (
	CONFIG_FILE_NAME = ".gitdump"
	CONFIG_FILE_TYPE = "yaml"
)

var str = `
Usage: gitdump SCOPE COMMAND [ARG1] [ARG2]...

A tool for downloading GitHub repositories

Scopes:
  users    Public repositories
  owners   Private repositories (require authentication)
  orgs     Organizations repositories (require authentication)

Commands:
  ls       List the repositories
  get      Download a single repository
  dump     Download all repositories
  config   Configuration

Args:
  user     Public account
  owner    Account associated with the the token
  org      Organization name
  repo     Repository name
`

var rootCmd = &cobra.Command{
	Use:   "gitdump",
	Short: "A tool for downloading GitHub repositories",
	Long:  `A tool for downloading GitHub repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(str)
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigName(CONFIG_FILE_NAME)
	viper.SetConfigType(CONFIG_FILE_TYPE)

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		config_file_path := fmt.Sprintf("%s/%s.%s", home, CONFIG_FILE_NAME, CONFIG_FILE_TYPE)

		new, err := os.Create(config_file_path)
		if err != nil {
			fmt.Println(err)
		}
		defer new.Close()

		err = os.Chmod(config_file_path, 0700)
		if err != nil {
			fmt.Println(err)
		}

		viper.Set("created_at", fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05")))
		viper.WriteConfig()

		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file, %s", err)
		}
	}
}
