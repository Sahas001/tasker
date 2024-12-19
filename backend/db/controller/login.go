package controller

import (
	"context"
)

const getUserByEmail = `SELECT * FROM users WHERE email=$1`

func (q *Queries) CheckUser(ctx context.Context, email string) (*User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email, &i.Password)
	return &i, err
}
