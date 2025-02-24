package cmd

import (
	"cvault/internal"
	"fmt"

	"github.com/spf13/cobra"
	"go.etcd.io/bbolt"
)

var getPasswordCmd *cobra.Command

func GetPasswordByKey(db *bbolt.DB) *cobra.Command {
	if getPasswordCmd != nil {
		return getPasswordCmd
	}

	getPasswordCmd = &cobra.Command{
		Use:     "get <key>",
		Short:   "Get local password by key",
		Example: "vault get dev-password",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Usage: vault get <key>")
				return
			}

			v, err := internal.GetPassword(db, args[0])
			if err != nil {
				fmt.Println("Error")
				return
			}

			fmt.Println(v)
			return

		},
	}

	return getPasswordCmd
}
