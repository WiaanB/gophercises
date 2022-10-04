package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "a cli based task manager",
}
