package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckUser(t *testing.T) {
	t.Parallel()
	// Add test cases
	_, err := testQueries.CheckUser(context.Background(), "z0fkbG@SKyi.com")
	require.NoError(t, err)

}
