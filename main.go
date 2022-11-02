package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func main() {
	fileName := "input.json"
	questions := readFromFile(fileName)

	sort.Slice(questions, func(i, j int) bool {
		return questions[i].Content < questions[j].Content
	})

	mapOfContents := convertToMap(questions)
}

func convertToMap(questions []Question) map[string][]Question {
	questionMap := make(map[string][]Question)
	for _, question := range questions {
		questionMap[question.Content] = append(questionMap[question.Content], question)
	}
	return questionMap
}

func readFromFile(fileName string) []Question {
	jsonInput, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("opening config file %v", err.Error())
	}

	jsonParser := json.NewDecoder(jsonInput)
	questions := make([]Question, 0)
	if err = jsonParser.Decode(&questions); err != nil {
		fmt.Printf("parsing config file %v", err.Error())
	}
	return questions
}
