package settings

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "ngrok_restater",
		Short: "A helper tool for restart ngork for free",
		Long:  "A helper tool for restart ngork for free",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version: 0.1")
		},
	}
)

// Execute executes the root command.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return nil
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config name", "config file (default is $APP_HOME/config.yaml)")
	rootCmd.PersistentFlags().StringP("mode", "m", "client", "execute mode client / server")
	rootCmd.PersistentFlags().StringP("port", "p", "8001", "server port. defult port: 8001")
	viper.BindPFlag("mode", rootCmd.PersistentFlags().Lookup("mode"))
	viper.BindPFlag("server.port", rootCmd.PersistentFlags().Lookup("port"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
