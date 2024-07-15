package main

import (
	"log"
	"os"

	cmd "github.com/vinit-chauhan/tasker/cmd/tasker"
	"github.com/vinit-chauhan/tasker/internal/db/types"
)

func init() {
	// Adding flags to `AddCmd`
	cmd.AddCmd.Flags().StringP(
		"project",
		"p",
		"",
		"specify a project for your task",
	)

	// Adding flags to `UpdateCmd`
	cmd.UpdateCmd.Flags().StringP(
		"name",
		"n",
		"",
		"specify a name for your task",
	)
	cmd.UpdateCmd.Flags().StringP(
		"project",
		"p",
		"",
		"specify a project for your task",
	)
	cmd.UpdateCmd.Flags().IntP(
		"status",
		"s",
		int(types.Todo),
		"specify a status for your task (0: Backlog, 1: Waiting, 2: Todo, 3: InDesign, 4: InProgress, 5: Done)",
	)

	// Declaring sub-commands to RootCmd
	cmd.RootCmd.AddCommand(cmd.AddCmd)
	cmd.RootCmd.AddCommand(cmd.DeleteCmd)
	cmd.RootCmd.AddCommand(cmd.ListCmd)
	cmd.RootCmd.AddCommand(cmd.RemoveDBCmd)
	cmd.RootCmd.AddCommand(cmd.UpdateCmd)
	cmd.RootCmd.AddCommand(cmd.WhereCmd)
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalf("error executing root command: %v", err.Error())
		os.Exit(1)
	}
}
