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
		log.Println("QuizLoader::Load::No Data Service Instantiated")
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
		log.Println("QuizLoader::Load::Error from DataService", err)
		return quiz, err
	}
	return ql.createQuizFromDynamoResults(id, returnValue), nil
}

func (ql DynamoQuizLoader) createQuizFromDynamoResults(id string, result interface{}) models.Quiz {
	output, ok := result.(*dynamodb.QueryOutput)
	if ok == false {
		log.Println("QuizLoader::createQuizFromDynamoResults::Unexpected Type")
		return models.Quiz{}
	}
	if *output.Count == 0 {
		log.Println("QuizLoader::createQuizFromDynamoResults::Found No Results in Dynamo")
		return models.Quiz{}
	}
	questions := make([]models.Question, len(output.Items))
	for i, v := range output.Items {
		question := models.Question{
			Id: *v["id"].S,
		}

		valueQuestion, hasQuestion := v["question"]
		if hasQuestion == true {
			question.Question = *valueQuestion.S
		}

		answers, err := ql.AnswerLoader.LoadBatch(question.Id)
		if err != nil {
			log.Println("QuizLoader::createQuizFromDynamoResults::Error from DataService", err)
		}
		question.Answers = answers
		questions[i] = question
	}
	return models.Quiz{
		Id:        id,
		Questions: questions,
	}
}
