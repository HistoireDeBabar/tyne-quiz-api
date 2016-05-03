package data

import (
	"fmt"
	"sync"

	"github.com/HistoireDeBabar/tyne-quiz-api/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoQuizSaver struct {
	Service DataService
	wg      sync.WaitGroup
}

func CreateDynamoDataSaver() QuizSaver {
	return DynamoQuizSaver{
		Service: CreateDynamoDataService(),
	}
}

func (dqs DynamoQuizSaver) Save(quiz *models.AnsweredQuiz) {
	if dqs.Service == nil {
		//log
		return
	}

	if quiz.Answers == nil || len(quiz.Answers) == 0 {
		//log
		return
	}

	dqs.wg.Add(len(quiz.Answers))
	for _, v := range quiz.Answers {
		go dqs.upload(v)
	}
	dqs.wg.Wait()
}

func (dqs *DynamoQuizSaver) upload(answer *models.Answer) {
	answerId := answer.Id
	if answerId == "" {
		fmt.Println("Answer has no Id.  Can not upload.")
		return
	}
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String(AnswerTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(answer.Id),
			},
			"questionId": {
				S: aws.String(answer.QuestionId),
			},
		},
		UpdateExpression: aws.String(AnswerUpdateExpression),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":inc": {
				N: aws.String("1"),
			},
		},
	}
	_, e := dqs.Service.UpdateItem(params)
	if e != nil {
		fmt.Println(e)
	}
	dqs.wg.Done()
}
