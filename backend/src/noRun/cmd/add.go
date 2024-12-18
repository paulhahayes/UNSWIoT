package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func addDevice() {
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
