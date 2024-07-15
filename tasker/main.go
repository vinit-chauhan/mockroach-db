package main

import (
	"log"
	"os"

	"github.com/vinit-chauhan/tasker/cmd/tasker"
	"github.com/vinit-chauhan/tasker/internal/db/types"
)

func init() {
	tasker.AddCmd.Flags().StringP(
		"project",
		"p",
		"",
		"specify a project for your task",
	)

	tasker.RootCmd.AddCommand(tasker.AddCmd)
	tasker.RootCmd.AddCommand(tasker.WhereCmd)
	tasker.RootCmd.AddCommand(tasker.ListCmd)
	tasker.RootCmd.AddCommand(tasker.RemoveDBCmd)

	tasker.UpdateCmd.Flags().StringP(
		"name",
		"n",
		"",
		"specify a name for your task",
	)
	tasker.UpdateCmd.Flags().StringP(
		"project",
		"p",
		"",
		"specify a project for your task",
	)

	tasker.UpdateCmd.Flags().IntP(
		"status",
		"s",
		int(types.Todo),
		"specify a status for your task (0: Backlog, 1: Waiting, 2: Todo, 3: InDesign, 4: InProgress, 5: Done)",
	)
	tasker.RootCmd.AddCommand(tasker.UpdateCmd)

}

func main() {
	if err := tasker.RootCmd.Execute(); err != nil {
		log.Fatalf("error executing root command: %v", err.Error())
		os.Exit(1)
	}
}
