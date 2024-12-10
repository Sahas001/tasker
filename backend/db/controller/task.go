package controller

import "context"

type AddTaskParams struct {
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

const createTask = `INSERT INTO tasks (user_id, title, description, category) VALUES ($1, $2, $3, $4) RETURNING id, user_id, title, description,category, status, created_at, updated_at`

func (q *Queries) CreateTask(ctx context.Context, arg AddTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask, arg.UserID, arg.Title, arg.Description, arg.Category)
	var i Task
	err := row.Scan(&i.ID, &i.UserID, &i.Title, &i.Descr, &i.Category, &i.Status, &i.CreatedAt, &i.UpdatedAt)
	return i, err
}

const getTasksByUserID = `SELECT * FROM tasks WHERE user_id=$1`

func (q *Queries) GetTasksByUserID(ctx context.Context, userID int) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getTasksByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(&i.ID, &i.UserID, &i.Title, &i.Descr, &i.Category, &i.Status, &i.CreatedAt, &i.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, i)
	}
	return tasks, nil
}

const deleteTask = `DELETE FROM tasks WHERE id=$1`

func (q *Queries) DeleteTask(ctx context.Context, id int) error {
	_, err := q.db.ExecContext(ctx, deleteTask, id)
	return err
}

type UpdateTaskParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

const updateTask = `UPDATE tasks SET title=$1, description=$2, status=$3 WHERE id=$4 RETURNING id, user_id, title, description,category, status, created_at, updated_at`

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams, id int) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTask, arg.Title, arg.Description, arg.Status, id)
	var i Task
	err := row.Scan(&i.ID, &i.UserID, &i.Title, &i.Descr, &i.Category, &i.Status, &i.CreatedAt, &i.UpdatedAt)
	return i, err
}
