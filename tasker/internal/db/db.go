package db

import (
	"log"
	"os"
	"runtime"

	"database/sql"
	"path/filepath"

	gap "github.com/muesli/go-app-paths"
	"github.com/vinit-chauhan/tasker/internal/db/types"

	_ "github.com/mattn/go-sqlite3"
)

// opens a SQLite db and stores the db file in location specified by `path“
func OpenDB(path string) (*types.TaskDB, error) {
	db, err := sql.Open("sqlite3", filepath.Join(path, "tasks.db"))
	if err != nil {
		return nil, err
	}

	t := types.New(db, path)

	if !t.TasksTableExists() {
		if err := t.CreateTasksTable(); err != nil {
			return nil, err
		}

	}

	return &t, nil
}

func SetupPath() string {
	scope := gap.NewScope(gap.User, "tasks")

	dirs, err := scope.DataDirs()
	if err != nil {
		log.Fatal(err)
	}

	var taskDir string
	if len(dirs) > 0 {
		taskDir = dirs[0]
	} else {
		if taskDir, err = os.UserHomeDir(); err != nil {
			log.Fatal(err)
		}
	}

	if err := initTaskDir(taskDir); err != nil {
		log.Fatal(err)
	}

	// the / at the end is not added in Linux so adding it here.
	if runtime.GOOS == "linux" {
		taskDir = taskDir + "/"
	}

	return taskDir
}

func initTaskDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(path, 0o770)
		}
		return err
	}
	return nil
}
