package models

import (
	"github.com/satori/go.uuid"
)

type Answer struct {
	Id         string `json:"id"`
	QuestionId string `json: "questionId"`
	Answer     string `json:"answer"`
}

func CreateAnswer(questionId string, answer string) *Answer {
	return &Answer{
		QuestionId: questionId,
		Answer:     answer,
		Id:         uuid.NewV4().String(),
	}
}
