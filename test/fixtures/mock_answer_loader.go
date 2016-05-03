package fixtures

import "github.com/HistoireDeBabar/tyne-quiz-api/models"

type MockReturnsEmptyAnswers struct{}

func (m MockReturnsEmptyAnswers) LoadBatch(id string) (answers []*models.Answer, err error) {
	return
}
