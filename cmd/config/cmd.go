package configcommand

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"totalcmd/util"
)

var shouldSave bool

var RootCmd = &cobra.Command{
	Use:   "config",
	Short: "Thùy chỉnh, hiện thị các config",
	RunE: func(cmd *cobra.Command, args []string) error {
		if shouldSave {
			configOutPath := getDefaultConfigPath()
			if err := viper.WriteConfigAs(configOutPath); err != nil {
				return err
			}
			fmt.Printf("Configs was save to: %s\n", configOutPath)
			return nil
		}
		return printConfig()
	},
}

func init() {
	RootCmd.Flags().BoolVar(&shouldSave, "save", false, "save config to default path")
}

func getDefaultConfigPath() string{
	defaultPath := viper.ConfigFileUsed()
	if defaultPath == "" {
		homedirPath, err := homedir.Dir()
		if err != nil {
			log.Fatalln(err)
		}
		defaultPath = homedirPath + "/" + util.CfgFileName + ".yml"
	}
	return defaultPath
}

func printConfig() error {
	allSettings := viper.AllSettings()
	for key, value := range allSettings {
		fmt.Printf("%s: %v\n", key, value)
	}
	return nil
}
