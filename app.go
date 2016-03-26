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
		QuizSaver:  data.CreateDynamoDataSaver(),
	}
	http.HandleFunc("/quiz", quizController.GetQuiz)
	http.HandleFunc("/answers", quizController.PostAnswers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
