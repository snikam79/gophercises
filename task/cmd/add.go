package cmd

import (
	"fmt"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

// go file dedicated for "add command"

var addCommand = &cobra.Command{
	// actual command name
	Use: "add",
	// short description
	Short: "Adds a task to the task list",
	//anonymous function that will invoked while executing add command
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong :", err.Error())
			return
		}
		fmt.Printf("Added \"%s \" task to task list. \n", task)
	},
}

// bootstrapping add command
func init() {
	RootCmd.AddCommand(addCommand)
}
