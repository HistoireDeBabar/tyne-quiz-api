package data

type DataService interface {
	GetItem(params interface{}) (result interface{}, err error)
	PutItem(params interface{}) (result interface{}, err error)
}
