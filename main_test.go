package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_readFile(t *testing.T) {
	path := "input.json"

	byteSlice, err := readFile(path)
	require.NoError(t, err)

	assert.NotEmpty(t, byteSlice)
}

func Test_unmarshallBytes(t *testing.T) {
	fileName := "input.json"
	inputB, err := readFile(fileName)
	require.NoError(t, err)

	inputSlice := unmarshallBytes(inputB)

	expectedId := 98673
	assert.Equal(t, expectedId, inputSlice[0].Id)
}

func Test_filterbyContent(t *testing.T) {
	inputs := []Input{
		{
			Id:              0,
			CreateTimestamp: 0,
			Content:         "No more!",
			Answers:         nil,
		},
		{
			Id:              0,
			CreateTimestamp: 0,
			Content:         "Baby",
			Answers: []Answers{
				{
					Id:      0,
					Rating:  20,
					Content: "",
				},
			},
		},
		{
			Id:              0,
			CreateTimestamp: 0,
			Content:         "Baby",
			Answers: []Answers{
				{
					Id:      0,
					Rating:  15,
					Content: "",
				},
			},
		},
	}

	response := filterByContent(inputs)

	expected := []Input{
		{
			Id:              0,
			CreateTimestamp: 0,
			Content:         "No more!",
			Answers:         nil,
		},
		{
			Id:              0,
			CreateTimestamp: 0,
			Content:         "Baby",
			Answers: []Answers{
				{
					Id:      0,
					Rating:  20,
					Content: "",
				},
			},
		},
	}

	assert.Equal(t, expected, response)
}
