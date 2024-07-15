package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/vinit-chauhan/tasker/internal/db"
)

var RootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A CLI task manager tool for your home project",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var addCmd = &cobra.Command{
	Use:   "add NAME",
	Short: "Add a new task with optional project name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := db.OpenDB(db.SetupPath())
		if err != nil {
			return err
		}

		defer t.GetDB().Close()

		project, err := cmd.Flags().GetString("project")
		if err != nil {
			return err
		}
		if err := t.Insert(args[0], project); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	addCmd.Flags().StringP(
		"project",
		"p",
		"",
		"specify a project for your task",
	)

	RootCmd.AddCommand(addCmd)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalf("error executing root command: %v", err.Error())
		os.Exit(1)
	}
}
