package fixtures

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/**
  Allows Access to Params in Public Field.
*/
type MockParamsReturnDataService struct {
	Params     interface{}
	SaveParams [2]interface{}
	SaveLength int
}

func (mds *MockParamsReturnDataService) GetItem(params interface{}) (result interface{}, err error) {
	mds.Params = params
	return result, nil
}

func (mds *MockParamsReturnDataService) UpdateItem(params interface{}) (result interface{}, err error) {
	mds.SaveParams[mds.SaveLength] = params
	mds.SaveLength++
	return result, nil
}

/**
  Retuns an Error
*/
type MockErrorDataService struct{}

func (mds *MockErrorDataService) GetItem(params interface{}) (result interface{}, err error) {
	return result, errors.New("DataServiceError")
}

func (mds *MockErrorDataService) UpdateItem(params interface{}) (result interface{}, err error) {
	return result, errors.New("DataServiceError")
}

/**
  Retuns an Empty Response
*/
type MockReturnsEmptyResultsDataService struct{}

func (mds *MockReturnsEmptyResultsDataService) GetItem(params interface{}) (result interface{}, err error) {
	result = &dynamodb.QueryOutput{
		Count: new(int64),
	}
	return result, nil
}

func (mds *MockReturnsEmptyResultsDataService) UpdateItem(params interface{}) (result interface{}, err error) {
	return result, errors.New("DataServiceError")
}

/**
  Retuns a different type of type
*/
type MockReturnsWrongTypedResultsDataService struct{}

func (mds *MockReturnsWrongTypedResultsDataService) GetItem(params interface{}) (result interface{}, err error) {
	result = "I shouldn't be a string"
	return result, nil
}

func (mds *MockReturnsWrongTypedResultsDataService) UpdateItem(params interface{}) (result interface{}, err error) {
	return result, errors.New("DataServiceError")
}

/**
  Retuns correct data
*/
type MockReturnsResultsDataService struct{}

func (mds *MockReturnsResultsDataService) GetItem(params interface{}) (result interface{}, err error) {
	count := &[]int64{2}[0]
	// question 1
	vq := &dynamodb.AttributeValue{
		S: aws.String("number 9 for newcastle"),
	}
	vi := &dynamodb.AttributeValue{
		S: aws.String("question-1"),
	}
	vqi := &dynamodb.AttributeValue{
		S: aws.String("test"),
	}
	// question 2
	vq2 := &dynamodb.AttributeValue{
		S: aws.String("town in which the great north run ends"),
	}
	vi2 := &dynamodb.AttributeValue{
		S: aws.String("question-2"),
	}
	vqi2 := &dynamodb.AttributeValue{
		S: aws.String("test"),
	}
	values := []map[string]*dynamodb.AttributeValue{
		{
			"question": vq,
			"id":       vi,
			"quizId":   vqi,
		},
		{
			"question": vq2,
			"id":       vi2,
			"quizId":   vqi2,
		},
	}

	output := &dynamodb.QueryOutput{
		Count: count,
		Items: values,
	}
	return output, nil
}

func (mds *MockReturnsResultsDataService) UpdateItem(params interface{}) (result interface{}, err error) {
	return result, errors.New("DataServiceError")
}

/**
  Retuns incorrect structure of data
*/
type MockReturnsResultsBadStructureDataService struct{}

func (mds *MockReturnsResultsBadStructureDataService) UpdateItem(params interface{}) (result interface{}, err error) {
	return result, errors.New("DataServiceError")
}

func (mds *MockReturnsResultsBadStructureDataService) GetItem(params interface{}) (result interface{}, err error) {
	count := &[]int64{2}[0]
	// question 1
	vq := &dynamodb.AttributeValue{
		S: aws.String("ques-id"),
	}
	vqi := &dynamodb.AttributeValue{
		S: aws.String("test"),
	}
	vi := &dynamodb.AttributeValue{
		S: aws.String("question-1"),
	}
	values := []map[string]*dynamodb.AttributeValue{
		{
			"nonconformed": vq,
			"id":           vi,
			"quizId":       vqi,
		},
	}

	output := &dynamodb.QueryOutput{
		Count: count,
		Items: values,
	}
	return output, nil
}

/**
  Retuns correct data for answers
*/
type MockReturnsResultsAnswerDataService struct{}

func (mds *MockReturnsResultsAnswerDataService) UpdateItem(params interface{}) (result interface{}, err error) {
	return result, errors.New("DataServiceError")
}
func (mds *MockReturnsResultsAnswerDataService) GetItem(params interface{}) (result interface{}, err error) {
	count := &[]int64{2}[0]
	// question 1
	vq := &dynamodb.AttributeValue{
		S: aws.String("a"),
	}
	vi := &dynamodb.AttributeValue{
		S: aws.String("1"),
	}
	va := &dynamodb.AttributeValue{
		S: aws.String("alan shearer"),
	}
	vac := &dynamodb.AttributeValue{
		N: aws.String("1"),
	}
	vacom := &dynamodb.AttributeValue{
		BOOL: aws.Bool(true),
	}
	// question 2
	vq2 := &dynamodb.AttributeValue{
		S: aws.String("a"),
	}
	vi2 := &dynamodb.AttributeValue{
		S: aws.String("2"),
	}
	va2 := &dynamodb.AttributeValue{
		S: aws.String("wayne rooney"),
	}
	vac2 := &dynamodb.AttributeValue{
		N: aws.String("2"),
	}
	vacom2 := &dynamodb.AttributeValue{
		BOOL: aws.Bool(false),
	}
	values := []map[string]*dynamodb.AttributeValue{
		{
			"questionId":  vq,
			"id":          vi,
			"answer":      va,
			"answerCount": vac,
			"correct":     vacom,
		},
		{
			"questionId":  vq2,
			"id":          vi2,
			"answer":      va2,
			"answerCount": vac2,
			"correct":     vacom2,
		},
	}

	output := &dynamodb.QueryOutput{
		Count: count,
		Items: values,
	}
	return output, nil
}
