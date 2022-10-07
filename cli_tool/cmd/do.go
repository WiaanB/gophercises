package cmd

import (
	"fmt"
	"os"
	"strconv"
	"workspace/cli_tool/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "crosses off a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			if err != nil {
				fmt.Println("Something went wrong: ", err.Error())
				os.Exit(1)
			}
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invaled task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as complete. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as complete\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
