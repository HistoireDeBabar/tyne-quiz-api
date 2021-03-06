package data

import (
	"testing"

	"github.com/HistoireDeBabar/tyne-quiz-api/fixtures"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestLoadQuizWithNoDataService(t *testing.T) {
	quizLoader := DynamoQuizLoader{}
	_, err := quizLoader.Load("test")
	if err == nil {
		t.Error("Expected an Error")
	}
}

func TestLoadWhenDataServiceReturnsAnError(t *testing.T) {
	service := &fixtures.MockErrorDataService{}
	quizLoader := DynamoQuizLoader{
		Service: service,
	}
	_, err := quizLoader.Load("test")
	if err == nil {
		t.Error("Expected an Error")
	}
}

func TestLoadQuizCallsDynamoWithCorrectParams(t *testing.T) {
	service := &fixtures.MockParamsReturnDataService{}
	quizLoader := DynamoQuizLoader{
		Service: service,
	}
	_, _ = quizLoader.Load("test")
	result, ok := service.Params.(*dynamodb.QueryInput)
	if ok == false {
		t.Error("Expected Params to be type QueryInput")
	}
	if *result.TableName != "Question" {
		t.Errorf("Expected TableName: Question, Got %v", result.TableName)
	}
	if *result.KeyConditionExpression != "quizId = :quizId" {
		t.Errorf("Expected KeyConditionExpression to equal quizId =: quizId, got %v", *result.KeyConditionExpression)
	}
	quizIdValue := *result.ExpressionAttributeValues[":quizId"]
	if quizIdValue.S == nil {
		t.Error("Expected :quizId to have value got nil")
	}
	if *quizIdValue.S != "test" {
		t.Errorf("Expected quiz id to eql test got %v", *quizIdValue.S)
	}
}

func TestDynamoCreateQuizFunctionHandlesNonDynamoTypes(t *testing.T) {
	service := &fixtures.MockReturnsWrongTypedResultsDataService{}
	quizLoader := DynamoQuizLoader{
		Service: service,
	}
	result, _ := quizLoader.Load("test")
	if result.IsEmpty() == false {
		t.Error("Expected an empty result set")
	}
}

func TestDynamoCreateQuizFunctionHandlesPoorlyStructuredTypes(t *testing.T) {
	service := &fixtures.MockReturnsResultsBadStructureDataService{}
	quizLoader := DynamoQuizLoader{
		Service:      service,
		AnswerLoader: &fixtures.MockReturnsEmptyAnswers{},
	}
	result, _ := quizLoader.Load("test")
	if result.IsEmpty() == false {
		t.Error("Expected Empty Result")
	}
}

func TestDynamoReturnsNoResults(t *testing.T) {
	service := &fixtures.MockReturnsEmptyResultsDataService{}
	quizLoader := DynamoQuizLoader{
		Service: service,
	}
	result, _ := quizLoader.Load("test")
	if result.IsEmpty() == false {
		t.Error("Expected an empty result set")
	}
}

func TestDynamoReturnsResults(t *testing.T) {
	service := &fixtures.MockReturnsResultsDataService{}
	quizLoader := DynamoQuizLoader{
		Service:      service,
		AnswerLoader: &fixtures.MockReturnsEmptyAnswers{},
	}
	result, _ := quizLoader.Load("test")
	if result.IsEmpty() == true {
		t.Error("Expected an empty result set")
	}
	if result.Id != "test" {
		t.Errorf("Expected Id to be test.  Got %v", result.Id)
	}

	if len(result.Questions) != 2 {
		t.Errorf("Expected Questions length to be 2.  Got %v", len(result.Questions))
	}
	question1 := result.Questions[0]

	if question1.Question != "number 9 for newcastle" {
		t.Errorf("Expected question1 to be number 9 for newcastle.  Got %v", question1.Question)
	}

	if question1.Id != "question-1" {
		t.Errorf("Expected question1 to be question1.  Got %v", question1.Id)
	}

}

func BenchmarkDynamoConnect(b *testing.B) {
	ddl := CreateDynamoDataLoader()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ddl.Load("test")
	}
}
