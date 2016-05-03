package data

import (
	"log"
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
		log.Println("QuizSaver::Save::No Data Service Instantiated")
		return
	}

	if quiz.Answers == nil || len(quiz.Answers) == 0 {
		log.Println("QuizSaver::Save::No Answers To Save")
		return
	}

	dqs.wg.Add(len(quiz.Answers))
	for _, v := range quiz.Answers {
		go dqs.upload(v)
	}
	dqs.wg.Wait()
}

func (dqs *DynamoQuizSaver) upload(answer *models.Answer) {
	defer dqs.wg.Done()
	if answer.IsValid() == false {
		log.Println("QuizSaver::upload::Answer value is Invalid")
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
		log.Println("QuizSaver::upload::Error From DataService::", e)
		return
	}
}
