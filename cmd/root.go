package cmd

import "github.com/spf13/cobra"

var RootCmd *cobra.Command

func GetRootCmd() *cobra.Command {
	if RootCmd != nil {
		return RootCmd
	}

	RootCmd = &cobra.Command{
		Use:   "vault",
		Short: "Command Line Utility for managing passwords",
	}

	return RootCmd
}
