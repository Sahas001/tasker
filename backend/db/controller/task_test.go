package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateTask(t *testing.T) (Task, int) {
	user := CreateTestUser(t)
	arg := AddTaskParams{
		UserID:      user.ID,
		Title:       RandomString(6),
		Description: RandomString(10),
		Category:    RandomString(6),
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

	return task, user.ID

}

func TestCreateTask(t *testing.T) {
	task, _ := CreateTask(t)
	testQueries.DeleteTask(context.Background(), task.ID)

}

func TestGetTasksByUserID(t *testing.T) {
	userID := 1
	taskID := 1
	name := "Lvfbhr"
	description := "xpr5BdGVnO"
	tasks, err := testQueries.GetTasksByUserID(context.Background(), userID)
	require.NoError(t, err)
	require.NotEmpty(t, tasks)
	require.Equal(t, taskID, tasks[0].ID)
	require.Equal(t, name, tasks[0].Title)
	require.Equal(t, description, tasks[0].Descr)

}

func TestDeleteTask(t *testing.T) {
	task, userID := CreateTask(t)
	err := testQueries.DeleteTask(context.Background(), task.ID)
	require.NoError(t, err)

	tasks, err := testQueries.GetTasksByUserID(context.Background(), userID)
	require.NoError(t, err)
	require.Empty(t, tasks)
	testQueries.DeleteUser(context.Background(), userID)
}

func TestUpdateTask(t *testing.T) {
	task, _ := CreateTask(t)
	taskID := task.ID
	arg := UpdateTaskParams{
		Title:       RandomString(6),
		Description: RandomString(10),
		Status:      true,
	}
	updatedTask, err := testQueries.UpdateTask(context.Background(), arg, taskID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedTask)
	require.Equal(t, arg.Title, updatedTask.Title)
	require.Equal(t, arg.Description, updatedTask.Descr)
	require.Equal(t, arg.Status, updatedTask.Status)
	require.NotZero(t, updatedTask.UpdatedAt)
	testQueries.DeleteTask(context.Background(), taskID)
}

func TestDeleteTasksByUserID(t *testing.T) {
	_, id := CreateTask(t)
	err := testQueries.DeleteTasksByUserID(context.Background(), id)
	require.NoError(t, err)

	tasks, err := testQueries.GetTasksByUserID(context.Background(), id)
	require.NoError(t, err)
	require.Empty(t, tasks)
	// testQueries.DeleteTask(context.Background(), task.ID)
}
