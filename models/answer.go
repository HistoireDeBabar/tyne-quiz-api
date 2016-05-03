package models

type Answer struct {
	Id          string `json:"id"`
	QuestionId  string `json:"questionId"`
	Correct     bool   `json:"correct"`
	AnswerCount int    `json:"answerCount"`
	Answer      string `json:"answer"`
}

func (a *Answer) IsValid() (valid bool) {
	if a.Id == "" {
		return false
	}
	if a.QuestionId == "" {
		return false
	}
	return true
}

func CreateAnswer(questionId string, id string) Answer {
	return Answer{
		QuestionId:  questionId,
		Id:          id,
		AnswerCount: 0,
	}
}
