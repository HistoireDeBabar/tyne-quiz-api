package data

import (
	"fmt"
	"github.com/HistoireDeBabar/tyne-quiz-api/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/satori/go.uuid"
	"sync"
)

const AnswerTableName = "Answer"

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
		go dqs.upload(v, quiz.Id)
	}
	dqs.wg.Wait()
}

func (dqs *DynamoQuizSaver) upload(answer *models.Answer, quizId string) {
	answerId := answer.Id
	if answerId == "" {
		answerId = uuid.NewV4().String()
	}
	params := &dynamodb.PutItemInput{
		TableName: aws.String(AnswerTableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id": &dynamodb.AttributeValue{
				S: aws.String(answerId),
			},
			"quizId": &dynamodb.AttributeValue{
				S: aws.String(quizId),
			},
			"questionId": &dynamodb.AttributeValue{
				S: aws.String(answer.QuestionId),
			},
			"answer": &dynamodb.AttributeValue{
				S: aws.String(answer.Answer),
			},
		},
	}
	_, e := dqs.Service.PutItem(params)
	if e != nil {
		fmt.Println(e)
	}
	dqs.wg.Done()
}
