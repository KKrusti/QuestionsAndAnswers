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

	response := make([]Question, 0)
	for _, content := range mapOfContents {
		response = append(response, getHighestRate(content))
	}

	sort.Slice(response, func(i, j int) bool {
		return response[i].Id < response[j].Id
	})

	out, err := json.Marshal(response)

	fmt.Println(string(out))
}

func getHighestRate(questions []Question) Question {
	responseQuestion := questions[0]
	responseQuestionRate := 0
	for _, question := range questions {
		rate := checkHighestRangeAnswer(question.Answers)
		if responseQuestionRate < rate[0].Rating {
			responseQuestionRate = rate[0].Rating
			responseQuestion = question
		} else if question.CreateTimestamp < responseQuestion.CreateTimestamp {
			responseQuestionRate = rate[0].Rating
			responseQuestion = question
		}
	}
	return responseQuestion
}

func checkHighestRangeAnswer(answers []Answers) []Answers {
	sort.Slice(answers, func(i, j int) bool {
		return answers[i].Rating > answers[j].Rating
	})
	return answers
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
