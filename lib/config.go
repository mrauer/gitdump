package lib

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// ConfigToken sets the GitHub token in the configuration and logs the action.
func ConfigToken(args []string) {
	token := fmt.Sprintf("%s", args[0])
	viper.Set("github_token", token)
	viper.WriteConfig()
	log.Printf("GitHub Token has been set in config: %s", token)
}

// ConfigPath sets the download path in the configuration and logs the action.
func ConfigPath(args []string) {
	path := fmt.Sprintf("%s", args[0])
	viper.Set("download_path", path)
	viper.WriteConfig()
	log.Printf("Download path has been set in config: %s", path)
}
