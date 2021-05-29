package lib

import (
	"fmt"

	"github.com/spf13/viper"
)

func ConfigToken(args []string) {
	fmt.Println("GitHub Token has been set in config")
	viper.Set("github_token", fmt.Sprintf("%s", args[0]))
	viper.WriteConfig()
}

func ConfigPath(args []string) {
	fmt.Println("Download path has been set in config")
	viper.Set("download_path", fmt.Sprintf("%s", args[0]))
	viper.WriteConfig()
}
