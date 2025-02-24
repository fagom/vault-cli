package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.etcd.io/bbolt"
)

var DeletePasswordCmd *cobra.Command

func DeletePassword(db *bbolt.DB) *cobra.Command {
	if DeletePasswordCmd != nil {
		return DeletePasswordCmd
	}

	DeletePasswordCmd = &cobra.Command{
		Use:     "delete <key>",
		Short:   "Delete local password",
		Long:    "Deletes the local stored password based on the key",
		Example: "vault delete dev-password",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Usage: vault delete <key>")
				return
			}

			key := args[0]

			err := db.Update(func(tx *bbolt.Tx) error {
				bucket := tx.Bucket([]byte("passwords"))
				if bucket == nil {
					return fmt.Errorf("Password DB not found")
				}

				if bucket.Get([]byte(key)) == nil {
					return fmt.Errorf("Key %s not found", key)
				}

				return bucket.Delete([]byte(key))
			})

			if err != nil {
				fmt.Errorf("Unable to delete key", err)
				return
			}

			fmt.Println("Password deleted successfully")
		},
	}

	return DeletePasswordCmd
}
