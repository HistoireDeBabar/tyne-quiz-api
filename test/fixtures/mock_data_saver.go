package fixtures

import (
	"github.com/HistoireDeBabar/tyne-quiz-api/models"
)

type MockSaver struct {
	Params *models.AnsweredQuiz
}

func (ms *MockSaver) Save(answered *models.AnsweredQuiz) {
	ms.Params = answered
	return
}
