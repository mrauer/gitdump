package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	CONFIG_FILE_NAME = ".gitdump"
	CONFIG_FILE_TYPE = "yaml"
	CONFIG_FILE_PERSM = 0700
)

var rootCmd = &cobra.Command{
	Use:   "gitdump",
	Short: "A tool for downloading GitHub repositories",
	Long:  `A tool for downloading GitHub repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Welcome to GitDump!")
		log.Println("For usage information, run 'gitdump --help'")
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
		log.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	} else {
		configFilePath := fmt.Sprintf("%s/%s.%s", home, CONFIG_FILE_NAME, CONFIG_FILE_TYPE)

		newFile, err := os.Create(configFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer newFile.Close()

		err = os.Chmod(configFilePath, CONFIG_FILE_PERSM)
		if err != nil {
			log.Fatal(err)
		}

		viper.Set("created_at", time.Now().Format("2006-01-02 15:04:05"))
		viper.WriteConfig()

		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Error reading config file: %s\n", err)
		}
		log.Printf("Config file created at: %s\n", configFilePath)
	}
}
