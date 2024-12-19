package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func addDevice(dev string) error {
	err := queries.Devices.Add(dev)
	if err != nil {
		return err
	}
	return nil
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			println("Please add a device name.")
			return
		}
		if err := addDevice(args[0]); err != nil {
			println(err.Error())
		}
		fmt.Println("Device successfully added")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
