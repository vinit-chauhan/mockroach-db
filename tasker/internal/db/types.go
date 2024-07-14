package db

import (
	"database/sql"
	"fmt"
)

const tableName = "tasks"

type TaskDB struct {
	db      *sql.DB
	dataDir string
}

func (t *TaskDB) tasksTableExists() bool {
	qStr := fmt.Sprintf("SELECT * FROM %v", tableName)
	if _, err := t.db.Query(qStr); err != nil {
		return false
	}
	return true
}

func (t *TaskDB) createTasksTable() error {
	qStr := fmt.Sprintf(`CREATE TABLE "%v" ( "id" INTEGER, "name" TEXT NOT NULL, "project" TEXT, "status" TEXT, "created" DATETIME, PRIMARY KEY("id" AUTOINCREMENT))`, tableName)

	_, err := t.db.Exec(qStr)

	return err
}
