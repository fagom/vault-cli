package main

import (
	"cvault/cmd"
	"cvault/internal"
	"fmt"
	"os"
)

func main() {
	db := internal.InitDb()
	defer db.Close()
	var rootCmd = cmd.GetRootCmd()
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	var AddCmd = cmd.AddLocalPassword(db)
	var ListCmd = cmd.ListPasswords(db)
	var DeleteCmd = cmd.DeletePassword(db)
	var GetCmd = cmd.GetPasswordByKey(db)

	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(DeleteCmd)
	rootCmd.AddCommand(GetCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
