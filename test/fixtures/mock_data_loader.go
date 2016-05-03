package fixtures

import (
	"errors"

	"github.com/HistoireDeBabar/tyne-quiz-api/models"
)

type MockQuizLoaderReturnsBasicQuiz struct{}

func (m MockQuizLoaderReturnsBasicQuiz) Load(id string) (quiz models.Quiz, err error) {
	return models.Quiz{
		Id: "1",
		Questions: []models.Question{
			{
				Id:       "a",
				Question: "whats your name",
				Answers: []*models.Answer{
					{
						Id:         "b",
						QuestionId: "a",
					},
				},
			},
		},
	}, nil
}

type MockQuizLoaderAccessParams struct{}

func (m MockQuizLoaderAccessParams) Load(id string) (quiz models.Quiz, err error) {
	if id == "test" {
		return models.Quiz{
			Id: "1",
			Questions: []models.Question{
				{
					Id:       "a",
					Question: "whats your name",
					Answers: []*models.Answer{
						{
							Id:         "b",
							QuestionId: "a",
						},
					},
				},
			},
		}, nil
		return models.Quiz{
			Id: "1",
		}, nil
	} else {
		return models.Quiz{}, errors.New("ID NOT PRESENT")
	}
}

type MockQuizLoaderAccessParamsEmpty struct{}

func (m MockQuizLoaderAccessParamsEmpty) Load(id string) (quiz models.Quiz, err error) {
	return
}

type MockError struct{}

func (m MockError) Load(id string) (quiz models.Quiz, err error) {
	return models.Quiz{}, errors.New("ID NOT PRESENT")
}
