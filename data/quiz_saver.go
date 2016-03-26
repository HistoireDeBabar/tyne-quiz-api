package data

import (
	"github.com/HistoireDeBabar/tyne-quiz-api/models"
)

type QuizSaver interface {
	Save(quiz *models.AnsweredQuiz)
}
