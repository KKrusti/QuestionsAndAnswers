package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_readFromFile(t *testing.T) {
	fileName := "resources/input.json"
	got, err := readFromFile(fileName)
	require.NoError(t, err)

	assert.Equal(t, 9, len(got))
	assert.Equal(t, 98673, got[0].Id)
	assert.Equal(t, 18345, got[1].Id)
}
