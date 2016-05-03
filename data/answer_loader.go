package data

import "github.com/HistoireDeBabar/tyne-quiz-api/models"

type AnswerLoader interface {
	// Given an Id of a question, load in multiple answers.
	LoadBatch(id string) (answers []models.Answer, err error)
}
