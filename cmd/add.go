package cmd

import (
	"cvault/internal"
	"fmt"

	"github.com/spf13/cobra"
	"go.etcd.io/bbolt"
)

var AddLocalPasswordCmd *cobra.Command

func AddLocalPassword(db *bbolt.DB) *cobra.Command {
	if AddLocalPasswordCmd != nil {
		return AddLocalPasswordCmd
	}

	AddLocalPasswordCmd = &cobra.Command{
		Use:   "add",
		Short: "Create a new local password",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				fmt.Println("Usage: vault add <key> <value>")
				return
			}

			internal.CreatePassword(db, args[0], args[1])

			fmt.Println("Password added successfully")
		},
	}

	return AddLocalPasswordCmd
}
