package data

import (
	"github.com/HistoireDeBabar/tyne-quiz-api/models"
)

type QuizLoader interface {
	Load(id string) (quiz models.Quiz, err error)
}
