package controller

import (
	"context"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func RandomUser(t *testing.T) AddUserParams {
	// create a random user with random data that is not 'test' each time this function is called

	name := RandomString(6)
	email := RandomEmail()
	password := RandomString(10)
	return AddUserParams{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func RandomString(n int) string {
	// create a random string of length n
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func RandomEmail() string {
	// create a random email
	return RandomString(6) + "@" + RandomString(4) + ".com"
}

func CreateTestUser(t *testing.T) User {
	arg := RandomUser(t)

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)

	require.NotZero(t, user.ID)

	return user
}

func TestCreateUser(t *testing.T) {
	user := CreateTestUser(t)
	testQueries.DeleteUser(context.Background(), user.ID)
}

func TestDeleteUser(t *testing.T) {
	user := CreateTestUser(t)
	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	user, err := testQueries.GetUser(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, user)
}

func TestUpdateUser(t *testing.T) {
	user := CreateTestUser(t)
	arg := RandomUser(t)

	user1, err := testQueries.UpdateUser(context.Background(), arg, user.ID)
	require.NoError(t, err)
	require.NotEqual(t, user.Name, user1.Name)
	require.NotEqual(t, user.Email, user1.Email)
	require.NotEqual(t, user.Password, user1.Password)
	testQueries.DeleteUser(context.Background(), user.ID)

}
