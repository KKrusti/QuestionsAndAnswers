package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
)

func main() {
	fileName := "resources/input.json"
	questions, err := readFromFile(fileName)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	sort.Slice(questions, func(i, j int) bool {
		return questions[i].Content < questions[j].Content
	})

	mapOfContents := convertToMap(questions)

	println(mapOfContents)
}

func convertToMap(questions []Question) map[string][]Question {
	questionMap := make(map[string][]Question)
	for _, question := range questions {
		questionMap[question.Content] = append(questionMap[question.Content], question)
	}
	return questionMap
}

func readFromFile(fileName string) ([]Question, error) {
	jsonInput, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("opening config file %v", err.Error()))

	}

	jsonParser := json.NewDecoder(jsonInput)
	questions := make([]Question, 0)
	if err = jsonParser.Decode(&questions); err != nil {
		return nil, errors.New(fmt.Sprintf("parsing config file %v", err.Error()))
	}
	return questions, nil
}
