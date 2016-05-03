package models

type Question struct {
	Id       string    `json:"id"`
	Question string    `json:"question"`
	Answers  []*Answer `json: "answers`
}

func (q *Question) IsValid() bool {
	if q.Id == "" || q.Question == "" {
		return false
	}
	return true
}
