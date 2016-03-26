package main

import (
	"github.com/HistoireDeBabar/tyne-quiz-api/ctrl"
	"github.com/HistoireDeBabar/tyne-quiz-api/data"
	"log"
	"net/http"
)

func main() {
	quizController := ctrl.QuizController{
		QuizLoader: data.CreateDynamoDataLoader(),
	}
	http.HandleFunc("/quiz", quizController.GetQuiz)
	log.Fatal(http.ListenAndServe(":80", nil))
}
