package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "retorrent",
	Short: "a simple torrent client",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
		██████  ███████ ████████  ██████  ██████  ██████  ███████ ███    ██ ████████ 
		██   ██ ██         ██    ██    ██ ██   ██ ██   ██ ██      ████   ██    ██    
		██████  █████      ██    ██    ██ ██████  ██████  █████   ██ ██  ██    ██    
		██   ██ ██         ██    ██    ██ ██   ██ ██   ██ ██      ██  ██ ██    ██    
		██   ██ ███████    ██     ██████  ██   ██ ██   ██ ███████ ██   ████    ██    
		`)

		cmd.Help()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Can't read config:", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
