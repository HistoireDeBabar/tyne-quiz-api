package models

type Quiz struct {
	Id        string     `json:"id"`
	Questions []Question `json:"questions"`
}

func (q *Quiz) IsEmpty() bool {
	if q.Questions == nil || q.Id == "" {
		return true
	}
	invalidCounter := 0
	questionLength := len(q.Questions)
	for _, p := range q.Questions {
		if p.IsValid() == false {
			invalidCounter++
		}
	}
	if questionLength == 0 || questionLength == invalidCounter {
		return true
	}
	return false
}
