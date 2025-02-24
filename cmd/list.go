package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"go.etcd.io/bbolt"
)

var ListPasswordsCmd *cobra.Command

func ListPasswords(db *bbolt.DB) *cobra.Command {
	if ListPasswordsCmd != nil {
		return ListPasswordsCmd
	}

	ListPasswordsCmd = &cobra.Command{
		Use:     "list",
		Short:   "Fetch all passwords",
		Long:    "Retrieves all local passwords",
		Example: "vault list",
		Run: func(cmd *cobra.Command, args []string) {
			var keys []string
			db.View(func(tx *bbolt.Tx) error {
				b := tx.Bucket([]byte("passwords"))
				if b == nil {
					return fmt.Errorf("No passwords stored")
				}

				b.ForEach(func(k, v []byte) error {
					keys = append(keys, string(k))
					return nil
				})
				return nil
			})

			if len(keys) == 0 {
				fmt.Println("No passwords found")
				return
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"#", "Password Key"})

			for i, key := range keys {
				table.Append([]string{fmt.Sprintf("%d", i+1), key})
			}

			table.Render()
		},
	}

	return ListPasswordsCmd
}
