package ctrl

import (
	"encoding/json"
	"fmt"
	"github.com/HistoireDeBabar/tyne-quiz-api/data"
	"net/http"
)

type QuizController struct {
	QuizLoader data.QuizLoader
}

const GET = "GET"

func (qc *QuizController) GetQuiz(w http.ResponseWriter, r *http.Request) {
	// Only accepts GET Methods
	if r.Method != GET {
		http.Error(w, "Status Method Not Allow", http.StatusMethodNotAllowed)
		return
	}
	// Returns 404 if id is not in query string
	url := r.URL
	queryString := url.Query()
	id := queryString["id"]

	if id == nil || len(id) == 0 {
		http.Error(w, "Id Required", http.StatusBadRequest)
		return
	}
	// Loads the Quiz from the id
	testId := id[0]
	result, err := qc.QuizLoader.Load(testId)
	// If an error occurs return a 500
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonResult[:]))
}
