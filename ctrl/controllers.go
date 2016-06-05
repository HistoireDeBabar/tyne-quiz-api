package ctrl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/HistoireDeBabar/tyne-quiz-api/data"
	"github.com/HistoireDeBabar/tyne-quiz-api/models"
)

type QuizController struct {
	QuizLoader data.QuizLoader
	QuizSaver  data.QuizSaver
}

const GET = "GET"
const POST = "POST"
const OPTIONS = "OPTIONS"

func (qc *QuizController) PostAnswers(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("Request::%s::Path::%s", r.Method, r.URL))
	// For CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	if r.Method == OPTIONS {
		return
	}
	// Only accepts POST Methods
	if r.Method != POST {
		http.Error(w, "Status Method Not Allow", http.StatusMethodNotAllowed)
		return
	}
	answer, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		http.Error(w, "Invalid Body", http.StatusInternalServerError)
		return
	}

	var answered models.AnsweredQuiz
	err = json.Unmarshal(answer, &answered)
	if err != nil || answered.IsValid() == false {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	qc.QuizSaver.Save(&answered)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{}")
}

func (qc *QuizController) GetQuiz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println(fmt.Sprintf("Request::%s::Path::%s", r.Method, r.URL))
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
	if result.IsEmpty() {
		http.Error(w, "Resource Not Found", http.StatusNotFound)
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
