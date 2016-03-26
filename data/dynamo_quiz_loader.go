package data

import (
	"errors"
	"github.com/HistoireDeBabar/tyne-quiz-api/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const QuestionTableName = "Question"
const QuizQuestionQuery = "quizId = :quizId"

type DynamoQuizLoader struct {
	Service DataService
}

func CreateDynamoDataLoader() QuizLoader {
	return DynamoQuizLoader{
		Service: CreateDynamoDataService(),
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
	return CreateQuizFromDynamoResults(id, returnValue), nil
}

func CreateQuizFromDynamoResults(id string, result interface{}) models.Quiz {
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
			Id: *v["id"].S,
		}

		valueAnswer, hasAnswer := v["answer"]
		if hasAnswer == true {
			question.Answer = *valueAnswer.S
		}

		valueQuestion, hasQuestion := v["question"]
		if hasQuestion == true {
			question.Question = *valueQuestion.S
		}

		valueOptions, hasOptions := v["options"]
		if hasOptions == true {
			options := make([]string, len(valueOptions.L))
			for j, ov := range valueOptions.L {
				options[j] = *ov.S
			}
			question.Options = options
		}
		if question.IsValid() == true {
			questions[i] = question
		}
	}
	return models.Quiz{
		Id:        id,
		Questions: questions,
	}
}
