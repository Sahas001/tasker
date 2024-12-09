package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateTask(t *testing.T) Task {
	user := CreateTestUser(t)
	arg := AddTaskParams{
		UserID:      user.ID,
		Title:       RandomString(6),
		Description: RandomString(10),
	}

	task, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, arg.UserID, task.UserID)
	require.Equal(t, arg.Title, task.Title)
	require.Equal(t, arg.Description, task.Descr)

	require.NotZero(t, task.ID)
	require.NotZero(t, task.CreatedAt)
	require.NotZero(t, task.UpdatedAt)

	return task
}

func TestCreateTask(t *testing.T) {
	CreateTask(t)
}
