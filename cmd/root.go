package cmd

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// SparkClient is Cisco Spark Client
var SparkClient *ciscospark.Client

// Max results to return
var Max int

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-spark",
	Short: "Displays the help",
	Long:  `Displays the go-spark help`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		SparkClient = ciscospark.NewClient(client)
		var token string
		viper.BindEnv("CISCO_SPARK_TOKEN")
		if viper.IsSet("CISCO_SPARK_TOKEN") {
			token = viper.GetString("CISCO_SPARK_TOKEN")
		} else if viper.IsSet("token") {
			token = viper.GetString("token")
		} else {
			fmt.Println(cmd.Help())
			os.Exit(-1)
		}
		SparkClient.Authorization = "Bearer " + token
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-spark.yaml)")
	RootCmd.PersistentFlags().IntVarP(&Max, "max", "m", 10, "Max results to return")

	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".go-spark") // name of config file (without extension)
	viper.AddConfigPath("$HOME")     // adding home directory as first search path
	viper.AutomaticEnv()             // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
