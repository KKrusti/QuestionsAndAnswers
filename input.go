package main

type Input struct {
	Id              int       `json:"id"`
	CreateTimestamp int       `json:"createTimestamp"`
	Content         string    `json:"content"`
	Answers         []Answers `json:"answers"`
}

type Answers struct {
	Id      int    `json:"id"`
	Rating  int    `json:"rating"`
	Content string `json:"content"`
}
