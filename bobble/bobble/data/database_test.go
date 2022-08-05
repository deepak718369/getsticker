package data

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSticker(t *testing.T) {

	GetConnection()
	err, _ := NewBobbleDatabaseLayer().GetSticker(context.TODO())
	require.NoError(t, err)

}
