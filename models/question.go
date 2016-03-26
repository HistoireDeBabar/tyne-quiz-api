package models

type Question struct {
	Id       string   `json:"id"`
	Answer   string   `json:"answer"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
}

func (q *Question) IsValid() bool {
	if q.Id == "" || q.Answer == "" || q.Question == "" {
		return false
	}
	if q.Options == nil || len(q.Options) == 0 {
		return false
	}
	return true
}
