package data

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDataService struct {
	dynamoTable *dynamodb.DynamoDB
}

func CreateDynamoDataService() DataService {
	return DynamoDataService{
		dynamoTable: dynamodb.New(session.New(&aws.Config{Region: aws.String("eu-west-1")})),
	}
}

func (dds DynamoDataService) PutItem(params interface{}) (result interface{}, err error) {
	input, ok := params.(*dynamodb.PutItemInput)
	if ok == false {
		return nil, errors.New("Invalid Param Type. Expected QueryInput")
	}
	result, err = dds.dynamoTable.PutItem(input)
	return
}

func (dds DynamoDataService) GetItem(params interface{}) (result interface{}, err error) {
	input, ok := params.(*dynamodb.QueryInput)
	if ok == false {
		return nil, errors.New("Invalid Param Type. Expected QueryInput")
	}
	return dds.dynamoTable.Query(input)
}
