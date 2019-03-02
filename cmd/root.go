package cmd

import (
	"fmt"
	"os"

	"github.com/jwhitaker/gomyip/pkg/myip"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gomyip",
	Short: "Retrieve your IP address from an external service",
	Run: func(cmd *cobra.Command, args []string) {
		ipAddress, err := myip.GetIPAddress()

		if err != nil {
			fmt.Printf("Failed to retrieve IP Address: %s", err)
		}

		fmt.Printf("%s", ipAddress)
	},
}

// Execute run the commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
