package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	// cfgFile     string
	// userLicense string

	rootCmd = &cobra.Command{
		Use:   "cbcli",
		Short: "cbcli - A cobra/viper cli",
		Long:  `aigen  - A cobra/viper cli`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

// func init() {
// 	cobra.OnInitialize(initConfig)

// 	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
// 	// rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
// 	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
// 	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
// 	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
// 	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
// 	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
// 	// viper.SetDefault("license", "apache")

// 	//rootCmd.AddCommand(ver)
// 	//rootCmd.AddCommand(initCmd)
// }

// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := os.UserHomeDir()
// 		cobra.CheckErr(err)

// 		// Search config in home directory with name ".cobra" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigType("yaml")
// 		viper.SetConfigName(".cobra")
// 	}

// 	viper.AutomaticEnv()

// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
