package types

import (
	"database/sql"
	"fmt"
	"time"
)

const tableName = "tasks"

type TaskDB struct {
	db      *sql.DB
	dataDir string
}

func New(db *sql.DB, dir string) TaskDB {
	return TaskDB{
		db:      db,
		dataDir: dir,
	}
}

func (t *TaskDB) TasksTableExists() bool {
	qStr := fmt.Sprintf("SELECT * FROM %v", tableName)
	if _, err := t.db.Query(qStr); err != nil {
		return false
	}
	return true
}

func (t *TaskDB) CreateTasksTable() error {
	qStr := fmt.Sprintf(`CREATE TABLE "%v" ( "id" INTEGER, "name" TEXT NOT NULL, "project" TEXT, "status" TEXT, "created" DATETIME, "updated" DATETIME, PRIMARY KEY("id" AUTOINCREMENT))`, tableName)

	_, err := t.db.Exec(qStr)

	return err
}

func (t *TaskDB) Insert(name, project string) error {

	qry := fmt.Sprintf(
		"INSERT INTO %v(name, project, status, created, updated) VALUES(?, ?, ?, ?, ?)",
		tableName,
	)

	_, err := t.db.Exec(
		qry,
		name,
		project,
		todo.String(),
		time.Now(),
		time.Now(),
	)

	return err
}

func (t *TaskDB) Delete(id uint) error {
	qry := fmt.Sprintf(
		"DELETE FROM %v WHERE id = ?",
		tableName,
	)
	_, err := t.db.Exec(
		qry,
		id,
	)

	return err
}

func (t *TaskDB) Update(task Task) error {

	qry := fmt.Sprintf(
		"UPDATE %v SET name = ?, project = ?, status = ?, updated = ? WHERE id = ?",
		tableName,
	)

	orig, err := t.Get(task.ID)
	if err != nil {
		return err
	}

	orig.Merge(task)
	_, err = t.db.Exec(
		qry,
		orig.Name,
		orig.Project,
		orig.Status,
		time.Now(),
	)

	return err
}

func (t *TaskDB) Get(id uint) (Task, error) {
	var task Task
	qry := fmt.Sprintf("SELECT * FROM %v WHERE id = ?", tableName)

	err := t.db.
		QueryRow(qry, id).
		Scan(
			&task.ID,
			&task.Name,
			&task.Project,
			&task.Status,
			&task.Created,
			&task.Updated,
		)

	return task, err
}

func (t *TaskDB) GetAll() ([]Task, error) {
	var tasks []Task
	qry := fmt.Sprintf("SELECT * FROM %v", tableName)

	rows, err := t.db.Query(qry)
	if err != nil {
		return tasks, fmt.Errorf("Unable to get values: %w", err)
	}

	for rows.Next() {
		var task Task
		if err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Project,
			&task.Status,
			&task.Created,
			&task.Updated,
		); err != nil {
			return tasks, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *TaskDB) GetAllByStatus(status string) ([]Task, error) {
	var tasks []Task
	qry := fmt.Sprintf("SELECT * FROM %v WHERE status = ?", tableName)

	rows, err := t.db.Query(qry, status)
	if err != nil {
		return tasks, fmt.Errorf("Unable to get values: %w", err)
	}

	for rows.Next() {
		var task Task
		if err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Project,
			&task.Status,
			&task.Created,
			&task.Updated,
		); err != nil {
			return tasks, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
