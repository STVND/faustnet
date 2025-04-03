/*
Copyright © 2024 Steven Davis
*/
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"main.go/cmd"
)

func main() {

	faustPath, _ := os.UserHomeDir()
	faustPath = fmt.Sprintf("%s/faust", faustPath)
	faustPath = strings.ReplaceAll(faustPath, "\\", "/")

	if _, err := os.Stat(faustPath + "/config"); os.IsNotExist(err) {
		fmt.Println("No config folder found - creating directory")
		err := os.MkdirAll(faustPath+"/config", 0755)
		if err != nil {
			fmt.Println("Unable to create config folder")
		}

	}

	if _, err := os.Stat(faustPath + "/config/config.yaml"); os.IsNotExist(err) {
		fmt.Println("No config file found - creating file")
		confPath := fmt.Sprintf("%s/config/config.yaml", faustPath)
		_, err := os.Create(confPath)
		if err != nil {
			fmt.Println("Unable to create config file")
		}

	}

	initConfig(faustPath)

	cmd.Execute()
}

func makePath(fPath string, sdir string) {
	subPath := fPath + "/" + sdir

	if _, err := os.Stat(subPath); os.IsNotExist(err) {
		fmt.Println("Creating subdirectory: ", subPath)
		err := os.Mkdir(subPath, 0755)

		if err != nil {
			fmt.Println("Error creating directory: ", err)
			return
		}

	}
}

func initConfig(dir string) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir + "/config")

	viper.SetDefault("FaustDir", dir)
	viper.SetDefault("ServerSubDir", "server")
	viper.SetDefault("ConfigSubDir", "config")
	viper.SetDefault("LogSubDir", "log")
	viper.SetDefault("HostedSubDir", "hosting")

	viper.WriteConfig()

	subDirs := [4]string{viper.GetString("ExportSubDir"),
		viper.GetString("ServerSubDir"),
		viper.GetString("LogSubDir"),
		viper.GetString("HostedSubDir")}

	for _, sdir := range subDirs {
		go makePath(dir, sdir)
	}

}
