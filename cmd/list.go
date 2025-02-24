package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListPasswordsCmd *cobra.Command

func ListPasswords() *cobra.Command {
	if ListPasswordsCmd != nil {
		return ListPasswordsCmd
	}

	ListPasswordsCmd = &cobra.Command{
		Use:   "list",
		Short: "Fetch all passwords",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Passwords fetched successfully")
		},
	}

	return ListPasswordsCmd
}
