package tasker

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"

	"github.com/vinit-chauhan/tasker/internal/db"
	"github.com/vinit-chauhan/tasker/internal/db/types"
)

var RootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A CLI task manager tool for your home project",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var AddCmd = &cobra.Command{
	Use:   "add TASK_NAME",
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

var WhereCmd = &cobra.Command{
	Use:   "where",
	Short: "Show where your tasks are stored",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := fmt.Println(db.SetupPath())
		return err
	},
}

var DeleteCmd = &cobra.Command{
	Use:   "delete TASK_ID",
	Short: "Delete a task by TASK_ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := db.OpenDB(db.SetupPath())
		if err != nil {
			return err
		}
		defer t.GetDB().Close()

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		return t.Delete(uint(id))
	},
}

var UpdateCmd = &cobra.Command{
	Use:   "update TASK_ID",
	Short: "Update a task by TASK_ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := db.OpenDB(db.SetupPath())
		if err != nil {
			return err
		}
		defer t.GetDB().Close()

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		project, err := cmd.Flags().GetString("project")
		if err != nil {
			return err
		}
		status, err := cmd.Flags().GetInt("status")
		if err != nil {
			return err
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		newTask := types.Task{
			ID:      uint(id),
			Name:    name,
			Project: project,
			Status:  types.Status(status),
			Created: time.Time{},
			Updated: time.Now(),
		}

		return t.Update(newTask)
	},
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your tasks",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := db.OpenDB(db.SetupPath())
		if err != nil {
			return err
		}
		defer t.GetDB().Close()

		tasks, err := t.GetAll()
		if err != nil {
			return err
		}

		fmt.Print(setupTable(tasks))
		return nil
	},
}

var RemoveDBCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes the database",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		path := db.SetupPath()
		var buf string

		fmt.Printf("Do you really want to delete? (Y/n) ")
		fmt.Scanf("%s", &buf)
		if strings.ToLower(buf) == "y" {
			if err := os.RemoveAll(path); err != nil {
				return err
			}

			println("Cleanup Successful!!!")
			return nil
		}

		println("Removal discarded.")
		return nil
	},
}

func setupTable(tasks []types.Task) *table.Table {
	columns := []string{
		"ID",
		"Name",
		"Project",
		"Status",
		"Created At",
		"Updated At",
	}

	var rows [][]string

	for _, task := range tasks {
		rows = append(rows, []string{
			fmt.Sprintf("%d", task.ID),
			task.Name,
			task.Project,
			task.Status.String(),
			task.Created.Format("2006-01-02 16:04:05"),
			task.Updated.Format("2006-01-02 16:04:05"),
		})
	}

	t := table.New().Border(lipgloss.HiddenBorder()).
		Headers(columns...).
		Rows(rows...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return lipgloss.NewStyle().
					Foreground(lipgloss.Color("212")).
					Border(lipgloss.NormalBorder()).
					BorderTop(false).
					BorderLeft(false).
					BorderRight(false).
					BorderBottom(true).
					Bold(true)
			}
			if row%2 == 0 {
				return lipgloss.NewStyle().Foreground(lipgloss.Color("246"))
			}
			return lipgloss.NewStyle()
		})

	return t
}
