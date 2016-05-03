package data

import (
	"testing"

	"github.com/HistoireDeBabar/tyne-quiz-api/fixtures"
	"github.com/HistoireDeBabar/tyne-quiz-api/models"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestQuizSaverReturnsErrorIfServiceIsNil(t *testing.T) {
	saver := &DynamoQuizSaver{}
	saver.Save(&models.AnsweredQuiz{})
}

func TestQuizSaverReturnsErrorIfAnsweredQuizHasNoAnswers(t *testing.T) {
	service := &fixtures.MockParamsReturnDataService{}
	quizSaver := DynamoQuizSaver{
		Service: service,
	}
	quizSaver.Save(&models.AnsweredQuiz{
		Answers: []models.Answer{},
	})
}

func TestQuizSaverReturnsErrorIfAnsweredQuizHasNilAnswers(t *testing.T) {
	service := &fixtures.MockParamsReturnDataService{}
	quizSaver := DynamoQuizSaver{
		Service: service,
	}
	quizSaver.Save(&models.AnsweredQuiz{})
}

func TestQuizSaverReturnsErrorIfDynamoReturnsError(t *testing.T) {
	service := &fixtures.MockErrorDataService{}
	quizSaver := DynamoQuizSaver{
		Service: service,
	}
	answer := []models.Answer{
		models.Answer{
			Id:         "answer1",
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
	quizSaver := DynamoQuizSaver{
		Service: service,
	}
	answer := []models.Answer{
		models.Answer{
			Id:         "answer1",
			QuestionId: "question1",
		},
		models.Answer{
			Id:         "answer2",
			QuestionId: "question2",
		},
	}
	quizSaver.Save(&models.AnsweredQuiz{
		Id:      "test",
		Answers: answer,
	})
	for _, v := range service.SaveParams {
		put, ok := v.(*dynamodb.UpdateItemInput)

		if ok == false {
			t.Error("Expected Params to be type PutItemOutput")
		}

		if *put.TableName != "Answer" {
			t.Errorf("TableName expected: %v Got %v", *put.TableName)
		}

		id := *put.Key["id"]
		if id.S == nil || (*id.S != "answer1" && *id.S != "answer2") {
			t.Errorf("Expected id to be answer1 or answer2 got %v", *id.S)
		}

		questionId := *put.Key["questionId"]
		if questionId.S == nil || (*questionId.S != "question1" && *questionId.S != "question2") {
			t.Errorf("Expected quiz id to question1 or question2 got %v", *questionId.S)
		}

		updateExpression := *put.UpdateExpression
		if updateExpression != "SET answerCount = answerCount + :inc" {
			t.Errorf("Expected Update Expression to increment counter got %s", updateExpression)
		}
	}
}

func BenchmarkDynamoSaveConnection(b *testing.B) {
	dataSaver := CreateDynamoDataSaver()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		answer := []models.Answer{
			models.CreateAnswer("question2", "south shields"),
			models.CreateAnswer("question1", "shearer"),
		}
		dataSaver.Save(&models.AnsweredQuiz{
			Id:      "test",
			Answers: answer,
		})
	}
}
