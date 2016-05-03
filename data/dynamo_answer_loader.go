package data

import (
	"errors"
	"strconv"

	"github.com/HistoireDeBabar/tyne-quiz-api/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoAnswerLoader struct {
	Service DataService
}

func CreateDynamoAnswerLoader() AnswerLoader {
	return &DynamoAnswerLoader{
		Service: CreateDynamoDataService(),
	}
}

func (dal *DynamoAnswerLoader) LoadBatch(id string) (answers []*models.Answer, err error) {
	if dal.Service == nil {
		return answers, errors.New("DataService not instantiated")
	}

	params := &dynamodb.QueryInput{
		TableName:              aws.String(AnswerTableName),
		KeyConditionExpression: aws.String(AnswerQuestionIdQuery),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":questionId": {
				S: aws.String(id),
			},
		},
	}
	result, err := dal.Service.GetItem(params)
	if err != nil {
		return answers, err
	}
	return populateAnswers(result)
}

func populateAnswers(result interface{}) (answers []*models.Answer, err error) {
	output, ok := result.(*dynamodb.QueryOutput)
	if ok == false {
		return []*models.Answer{}, nil
	}
	if *output.Count == 0 {
		return []*models.Answer{}, nil
	}
	answers = make([]*models.Answer, len(output.Items))
	for i, v := range output.Items {
		answer := &models.Answer{
			Id:         *v["id"].S,
			QuestionId: *v["questionId"].S,
			Answer:     *v["answer"].S,
			Correct:    *v["correct"].BOOL,
		}
		count, _ := strconv.Atoi(*v["answerCount"].N)
		answer.AnswerCount = count
		answers[i] = answer
	}
	return
}
