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
		Use:     "add <key>",
		Short:   "Create a new local password",
		Long:    "Creates a local password with a key and value. The first parameter needs to be key and the second parameter needs to be the password. No white spaces will be recognised, either for the password or key",
		Example: `vault add dev-password asdjshdsjasdjs-asdsdsad`,
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
