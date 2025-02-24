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

	var AddCmd = cmd.AddLocalPassword(db)
	var ListCmd = cmd.ListPasswords()
	var DeleteCmd = cmd.DeletePassword()
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
