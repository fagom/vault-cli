package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "dev"

func GetCliVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version of the CLI",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("vault CLI version:", Version)
		},
	}
}
