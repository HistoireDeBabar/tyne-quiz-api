package models

type AnsweredQuiz struct {
	Id      string    `json:"id"`
	Answers []*Answer `json:"answers"`
}

func (aq *AnsweredQuiz) IsValid() bool {
	if aq.Id == "" {
		return false
	}
	if aq.Answers == nil || len(aq.Answers) == 0 {
		return false
	}
	return true
}
