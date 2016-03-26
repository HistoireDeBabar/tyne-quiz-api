package models

type AnsweredQuiz struct {
	Id      string    `json:"id"`
	Answers []*Answer `json:"answers"`
}
