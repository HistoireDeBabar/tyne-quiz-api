package models

type Answer struct {
	Id          string `json:"id"`
	QuestionId  string `json: "questionId"`
	Correct     bool   `json: "correct"`
	AnswerCount int    `json: "answerCount"`
	Answer      string `json: "answer"`
}

func CreateAnswer(questionId string, id string) *Answer {
	return &Answer{
		QuestionId: questionId,
		Id:         id,
	}
}
