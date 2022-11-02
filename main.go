package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
)

func main() {
	fileName := "input.json"
	inputB, err := readFile(fileName)
	if err != nil {
		fmt.Printf("error happended %v", err)
	}
	inputs := unmarshallBytes(inputB)

	sort.Slice(inputs, func(i, j int) bool {
		return inputs[i].Content < inputs[j].Content
	})

	filterByContent(inputs)

}

func filterByContent(inputs []Input) []Input {
	var response []Input
	var higherInput Input
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		if i < len(inputs)-1 {
			if input.Content != inputs[i+1].Content {
				response = append(response, input)
			} else {
				if input.Answers[0].Rating > higherInput.Answers[0].Rating {
				}
				higherInput = input
			}
		} else {
			response = append(response, input)
		}
	}
	return response
}

func unmarshallBytes(inputB []byte) []Input {
	var input []Input
	err := json.Unmarshal(inputB, &input)
	if err != nil {
		fmt.Printf("error happended unmarshalling Json%v", err)
	}
	return input
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error opening file %v", err))
	}
	byteArr := make([]byte, 1024)
	_, err = file.Read(byteArr)

	return byteArr, nil
}
