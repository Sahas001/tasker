package controller

import "context"

const createUser = `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email, password`

type AddUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg AddUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Email, arg.Password)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email, &i.Password)
	return i, err
}

const getUser = `SELECT * FROM users WHERE id=$1`

func (q *Queries) GetUser(ctx context.Context, id int) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email, &i.Password)
	return i, err
}

const updateUser = `UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4 RETURNING id, name, email, password`

func (q *Queries) UpdateUser(ctx context.Context, arg AddUserParams, id int) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.Name, arg.Email, arg.Password, id)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email, &i.Password)
	return i, err
}

const deleteUser = `DELETE FROM users WHERE id=$1`

func (q *Queries) DeleteUser(ctx context.Context, id int) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}
