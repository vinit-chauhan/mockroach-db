package db

import (
	"database/sql"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// opens a SQLite db and stores the db file in location specified by `path“
func openDB(path string) (*TaskDB, error) {
	db, err := sql.Open("sqlite3", filepath.Join(path, "tasks.db"))
	if err != nil {
		return nil, err
	}

	t := TaskDB{db, path}

	if !t.tasksTableExists() {
		if err := t.createTasksTable(); err != nil {
			return nil, err
		}

	}

	return &t, nil
}
