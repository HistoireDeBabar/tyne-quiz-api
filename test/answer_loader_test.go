package test

import (
	"testing"

	"github.com/HistoireDeBabar/tyne-quiz-api/data"
	"github.com/HistoireDeBabar/tyne-quiz-api/test/fixtures"
)

func TestGivenAListOfAnswersCallsDynamoSeveralTime(t *testing.T) {
	answerLoader := data.DynamoAnswerLoader{
		Service: &fixtures.MockReturnsResultsAnswerDataService{},
	}
	result, err := answerLoader.LoadBatch("id")
	if err != nil {
		t.Errorf("Epected Error to be nil got %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected Length of results to equal 2 for %d", len(result))
	}
	if result[0].Id != "1" {
		t.Errorf("Expected id to eql 1 got, %s", result[0].Id)
	}
	if result[0].QuestionId != "a" {
		t.Errorf("Expected question id to eql a got, %s", result[0].QuestionId)
	}
	if result[0].Answer != "alan shearer" {
		t.Errorf("Expected answer to eql alan shearer got, %s", result[0].Answer)
	}
	if result[0].AnswerCount != 1 {
		t.Errorf("Expected answer count to eql 1 got, %d", result[0].AnswerCount)
	}
	if result[0].Correct != true {
		t.Errorf("Expected correct to eql true got, %d", result[0].Correct)
	}
	if result[1].Id != "2" {
		t.Errorf("Expected id to eql 2 got, %s", result[1].Id)
	}
	if result[1].QuestionId != "a" {
		t.Errorf("Expected question id to eql a got, %s", result[1].QuestionId)
	}
	if result[1].Answer != "wayne rooney" {
		t.Errorf("Expected answer to eql wayne rooney got, %s", result[1].Answer)
	}
	if result[1].AnswerCount != 2 {
		t.Errorf("Expected answer count to eql 2 got, %d", result[1].AnswerCount)
	}
	if result[1].Correct != false {
		t.Errorf("Expected correct to eql false got, %v", result[1].Correct)
	}
}
