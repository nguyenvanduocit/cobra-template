package cmd

import (
	"cobra-template/cmd/config"
	"cobra-template/util"
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use: "cobra-template",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is ./%s.yaml or $HOME/%s.yaml)", util.CfgFileName, util.CfgFileName))
	RootCmd.PersistentFlags().StringP("token", "t", "", "Moneylover JWT, you can get it from web.moneylover.me.")
	viper.BindPFlag("token", RootCmd.PersistentFlags().Lookup("token"))

	RootCmd.AddCommand(config.ConfigCmd)
}

func initConfig() {
	viper.SetEnvPrefix("ml")
	viper.AutomaticEnv()
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(util.CfgFileName)
	}

	if err := viper.ReadInConfig(); err != nil {
		// log.Println(err)
	}
}
