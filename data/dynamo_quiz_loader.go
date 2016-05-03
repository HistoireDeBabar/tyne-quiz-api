package data

import (
	"errors"
	"log"

	"github.com/HistoireDeBabar/tyne-quiz-api/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoQuizLoader struct {
	Service      DataService
	AnswerLoader AnswerLoader
}

func CreateDynamoDataLoader() QuizLoader {
	dataService := CreateDynamoDataService()
	return DynamoQuizLoader{
		Service: dataService,
		AnswerLoader: &DynamoAnswerLoader{
			Service: dataService,
		},
	}
}

func (ql DynamoQuizLoader) Load(id string) (quiz models.Quiz, err error) {

	if ql.Service == nil {
		return quiz, errors.New("DataService not instantiated")
	}

	params := &dynamodb.QueryInput{
		TableName:              aws.String(QuestionTableName),
		KeyConditionExpression: aws.String(QuizQuestionQuery),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":quizId": {
				S: aws.String(id),
			},
		},
	}
	returnValue, err := ql.Service.GetItem(params)
	if err != nil {
		return quiz, err
	}
	return ql.createQuizFromDynamoResults(id, returnValue), nil
}

func (ql DynamoQuizLoader) createQuizFromDynamoResults(id string, result interface{}) models.Quiz {
	output, ok := result.(*dynamodb.QueryOutput)
	if ok == false {
		return models.Quiz{}
	}
	if *output.Count == 0 {
		return models.Quiz{}
	}
	questions := make([]models.Question, len(output.Items))
	for i, v := range output.Items {
		question := models.Question{
			Id:       *v["id"].S,
			Question: *v["question"].S,
		}

		answers, err := ql.AnswerLoader.LoadBatch(question.Id)
		if err == nil {
			question.Answers = answers
		} else {
			log.Println("error processing answers from dynamo", err)
		}
		questions[i] = question
	}
	return models.Quiz{
		Id:        id,
		Questions: questions,
	}
}
