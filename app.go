package main

import (
	"log"
	"net/http"

	"github.com/HistoireDeBabar/tyne-quiz-api/ctrl"
	"github.com/HistoireDeBabar/tyne-quiz-api/data"
)

const port = ":80"

func main() {
	log.Println("Initialising app")
	quizController := ctrl.QuizController{
		QuizLoader: data.CreateDynamoDataLoader(),
		QuizSaver:  data.CreateDynamoDataSaver(),
	}
	http.HandleFunc("/quiz", quizController.GetQuiz)
	http.HandleFunc("/answers", quizController.PostAnswers)
	log.Println("Listening on Port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
