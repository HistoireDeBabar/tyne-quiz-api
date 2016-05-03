package data

type DataService interface {
	GetItem(params interface{}) (result interface{}, err error)
	UpdateItem(params interface{}) (result interface{}, err error)
}
