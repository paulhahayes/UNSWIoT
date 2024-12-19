/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func deleteDevice(dev string) error {
	res, err := queries.Devices.Delete(dev)
	if err != nil {
		return err
	}
	if res {
		fmt.Printf("Successfully deleted device %s\n", dev)
	} else {
		fmt.Printf("Failed to delete %s \n", dev)
	}
	return nil
}

func deleteAllDevices() error {
	return queries.Devices.DeleteAll()
}

var (
	deleteAll bool
)
var delCmd = &cobra.Command{
	Use:   "del [deviceName]",
	Short: "Delete a device or all devices",
	Long:  `Delete a specific device by name, or use -A flag to delete all devices`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		switch {
		case deleteAll:
			err = deleteAllDevices()
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Failed to delete all devices: %v\n", err)
				return
			}
			fmt.Println("Successfully deleted all devices")

		case len(args) == 0:
			fmt.Fprintln(cmd.ErrOrStderr(), "Please include a device name or use -A to delete all devices")
			cmd.Usage()
			return

		default:
			err = deleteDevice(args[0])
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Failed to delete device %s: %v\n", args[0], err)
				return
			}
			fmt.Printf("Successfully deleted device %s\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
	delCmd.Flags().BoolVarP(&deleteAll, "all", "A", false, "Delete all devices")
}
