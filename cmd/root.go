package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var rootCmd *cobra.Command

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "aggregate_errors",
		Short: "Aggregate Errors and Review them",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: 初期化処理
			return nil
		},
	}
}

func Execute() {
	rootCmd = newRootCmd()
	saveCmd := newSaveCmd()
	log.Print(saveCmd.Use)
	rootCmd.AddCommand(saveCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// TODO: フラグ関連の処理を追加する
func init() {
	// cobra.OnInitialize(initConfig)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aggregate_errors.yaml)")
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// }

	// func initConfig() {
	// 	if cfgFile != "" {
	// 		viper.SetConfigFile(cfgFile)
	// 	} else {
	// 		home, err := homedir.Dir()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			os.Exit(1)
	// 		}
	// 		viper.AddConfigPath(home)
	// 		viper.SetConfigName(".aggregate_errors")
}

// 	viper.AutomaticEnv()
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
