package db

import (
	"log"
	"os"

	gap "github.com/muesli/go-app-paths"
)

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
