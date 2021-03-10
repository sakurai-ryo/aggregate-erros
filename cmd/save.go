package cmd

import (
	"aggregate_errors/save"
	"fmt"

	"github.com/spf13/cobra"
)

func newSaveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "save",
		Short: "save error",
		Long:  `save error with context infomation to Aggregate_Errors Dir`,
		RunE:  runE,
	}
}

func runE(cmd *cobra.Command, args []string) error {
	fmt.Println("save called")
	save.Save()
	return nil
}
