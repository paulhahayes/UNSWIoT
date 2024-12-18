/*
Copyright © 2024 NAME HERE <z5303576@ad.unsw.edu.au>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func listDevices() error {
	devices, err := queries.Devices.GetAll()
	if err != nil {
		return err
	}

	// Print devices
	for _, device := range devices {
		fmt.Printf("ID: %d, Name: %s\n", device.ID, device.Name)
	}
	return nil
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all devices",
	Long:  `List all devices in the IoT database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return listDevices()
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
