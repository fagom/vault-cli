package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeletePasswordCmd *cobra.Command

func DeletePassword() *cobra.Command {
	if DeletePasswordCmd != nil {
		return DeletePasswordCmd
	}

	DeletePasswordCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete password",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Password deleted successfully")
		},
	}

	return DeletePasswordCmd
}
