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

func (mds *MockParamsReturnDataService) PutItem(params interface{}) (result interface{}, err error) {
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

func (mds *MockErrorDataService) PutItem(params interface{}) (result interface{}, err error) {
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

func (mds *MockReturnsEmptyResultsDataService) PutItem(params interface{}) (result interface{}, err error) {
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

func (mds *MockReturnsWrongTypedResultsDataService) PutItem(params interface{}) (result interface{}, err error) {
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
	va := &dynamodb.AttributeValue{
		S: aws.String("shearer"),
	}
	vqi := &dynamodb.AttributeValue{
		S: aws.String("test"),
	}
	vo := &dynamodb.AttributeValue{
		L: []*dynamodb.AttributeValue{
			{
				S: aws.String("rooney"),
			},
			{
				S: aws.String("les ferdinand"),
			},
			{
				S: aws.String("beardsley"),
			},
		},
	}
	// question 2
	vq2 := &dynamodb.AttributeValue{
		S: aws.String("town in which the great north run ends"),
	}
	vi2 := &dynamodb.AttributeValue{
		S: aws.String("question-2"),
	}
	va2 := &dynamodb.AttributeValue{
		S: aws.String("south shields"),
	}
	vqi2 := &dynamodb.AttributeValue{
		S: aws.String("test"),
	}
	vo2 := &dynamodb.AttributeValue{
		L: []*dynamodb.AttributeValue{
			{
				S: aws.String("gateshead"),
			},
			{
				S: aws.String("boldon"),
			},
			{
				S: aws.String("sunderland"),
			},
		},
	}
	values := []map[string]*dynamodb.AttributeValue{
		{
			"question": vq,
			"id":       vi,
			"answer":   va,
			"quizId":   vqi,
			"options":  vo,
		},
		{
			"question": vq2,
			"id":       vi2,
			"answer":   va2,
			"quizId":   vqi2,
			"options":  vo2,
		},
	}

	output := &dynamodb.QueryOutput{
		Count: count,
		Items: values,
	}
	return output, nil
}

func (mds *MockReturnsResultsDataService) PutItem(params interface{}) (result interface{}, err error) {
	return result, errors.New("DataServiceError")
}

/**
  Retuns incorrect structure of data
*/
type MockReturnsResultsBadStructureDataService struct{}

func (mds *MockReturnsResultsBadStructureDataService) PutItem(params interface{}) (result interface{}, err error) {
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
