package test

import (
	"github.com/HistoireDeBabar/tyne-quiz-api/data"
	"github.com/HistoireDeBabar/tyne-quiz-api/models"
	"github.com/HistoireDeBabar/tyne-quiz-api/test/fixtures"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"testing"
)

func TestQuizSaverReturnsErrorIfServiceIsNil(t *testing.T) {
	saver := &data.DynamoQuizSaver{}
	saver.Save(&models.AnsweredQuiz{})
}

func TestQuizSaverReturnsErrorIfAnsweredQuizHasNoAnswers(t *testing.T) {
	service := &fixtures.MockParamsReturnDataService{}
	quizSaver := data.DynamoQuizSaver{
		Service: service,
	}
	quizSaver.Save(&models.AnsweredQuiz{
		Answers: []*models.Answer{},
	})
}

func TestQuizSaverReturnsErrorIfAnsweredQuizHasNilAnswers(t *testing.T) {
	service := &fixtures.MockParamsReturnDataService{}
	quizSaver := data.DynamoQuizSaver{
		Service: service,
	}
	quizSaver.Save(&models.AnsweredQuiz{})
}

func TestQuizSaverReturnsErrorIfDynamoReturnsError(t *testing.T) {
	service := &fixtures.MockErrorDataService{}
	quizSaver := data.DynamoQuizSaver{
		Service: service,
	}
	answer := []*models.Answer{
		&models.Answer{
			Id:         "answer1",
			Answer:     "shearer",
			QuestionId: "question1",
		},
	}
	quizSaver.Save(&models.AnsweredQuiz{
		Id:      "test",
		Answers: answer,
	})
}

func TestQuizSaverUsesCorrectParams(t *testing.T) {
	service := &fixtures.MockParamsReturnDataService{}
	quizSaver := data.DynamoQuizSaver{
		Service: service,
	}
	answer := []*models.Answer{
		&models.Answer{
			Id:         "answer1",
			Answer:     "shearer",
			QuestionId: "question1",
		},
		&models.Answer{
			Id:         "answer2",
			Answer:     "South Shields",
			QuestionId: "question2",
		},
	}
	quizSaver.Save(&models.AnsweredQuiz{
		Id:      "test",
		Answers: answer,
	})
	for _, v := range service.SaveParams {
		put, ok := v.(*dynamodb.PutItemInput)

		if ok == false {
			t.Error("Expected Params to be type PutItemOutput")
		}

		if *put.TableName != "Answer" {
			t.Errorf("TableName expected: %v Got %v", *put.TableName)
		}
		quizId := *put.Item["quizId"]
		if quizId.S == nil || *quizId.S != "test" {
			t.Errorf("Expected quiz id to be test got %v", *quizId.S)
		}

		id := *put.Item["id"]
		if id.S == nil || (*id.S != "answer1" && *id.S != "answer2") {
			t.Errorf("Expected id to be answer1 or answer2 got %v", *id.S)
		}

		questionId := *put.Item["questionId"]
		if questionId.S == nil || (*questionId.S != "question1" && *questionId.S != "question2") {
			t.Errorf("Expected quiz id to question1 or question2 got %v", *questionId.S)
		}

		a := *put.Item["answer"]
		if a.S == nil || (*a.S != "shearer" && *a.S != "South Shields") {
			t.Errorf("Expected answer to be shearer or South Shields")
		}
	}
}

func BenchmarkDynamoSaveConnection(b *testing.B) {
	dataSaver := data.CreateDynamoDataSaver()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		answer := []*models.Answer{
			models.CreateAnswer("question2", "south shields"),
			models.CreateAnswer("question1", "shearer"),
		}
		dataSaver.Save(&models.AnsweredQuiz{
			Id:      "test",
			Answers: answer,
		})
	}
}
