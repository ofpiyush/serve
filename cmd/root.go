package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	serve "github.com/ofpiyush/serve/cmd/serve"
	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "serve",
	Short: "A simple file server",
	Run:   serveCmd,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func serveCmd(cmd *cobra.Command, args []string) {
	var pathStr = "."

	if len(args) >= 1 {
		pathStr = args[0]
	}
	absPath, err := filepath.Abs(pathStr)
	if err != nil {
		log.Fatal(err)
	}
	serve.Start(absPath, 3000)
}

// func init() {
// 	cobra.OnInitialize(initConfig)

// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports Persistent Flags, which, if defined here,
// 	// will be global for your application.

// 	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.serve.yaml)")
// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

// // initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" { // enable ability to specify config file via flag
// 		viper.SetConfigFile(cfgFile)
// 	}

// 	viper.SetConfigName(".serve") // name of config file (without extension)
// 	viper.AddConfigPath("$HOME")  // adding home directory as first search path
// 	viper.AutomaticEnv()          // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
