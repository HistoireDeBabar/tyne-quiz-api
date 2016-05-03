package ctrl

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/HistoireDeBabar/tyne-quiz-api/data"
	"github.com/HistoireDeBabar/tyne-quiz-api/fixtures"
	"github.com/HistoireDeBabar/tyne-quiz-api/models"
)

func TestControllerReturnsJSONQuiz(t *testing.T) {
	controller := QuizController{
		QuizLoader: fixtures.MockQuizLoaderReturnsBasicQuiz{},
	}
	ts := httptest.NewServer(http.HandlerFunc(controller.GetQuiz))
	defer ts.Close()

	res, err := http.Get(ts.URL + "?id=test")
	if err != nil {
		log.Fatal(err)
	}
	question, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	expected := &models.Quiz{
		Id: "1",
		Questions: []models.Question{
			{
				Id:       "a",
				Question: "whats your name",
				Answers: []models.Answer{
					{
						Id:         "b",
						QuestionId: "a",
					},
				},
			},
		},
	}
	expectedString, err := json.Marshal(expected)
	if err != nil {
		t.Fatal("error:", err)
	}
	if !bytes.Equal(expectedString, question) {
		t.Fatalf("expected to eql: %v to not %v", expectedString, question)
	}
}

func TestControllerSetsHeaders(t *testing.T) {
	controller := QuizController{
		QuizLoader: fixtures.MockQuizLoaderReturnsBasicQuiz{},
	}
	ts := httptest.NewServer(http.HandlerFunc(controller.GetQuiz))
	defer ts.Close()

	res, err := http.Get(ts.URL + "?id=test")
	if err != nil {
		log.Fatal(err)
	}
	contentHeaders := res.Header["Content-Type"]
	jsonHeader := false
	for _, v := range contentHeaders {
		if v == "application/json" {
			jsonHeader = true
		}
	}
	if jsonHeader == false {
		log.Fatal("Expected application json to be in content headers")
	}
}

func TestControllerReturns404IfIdIsNotPresentInQueryString(t *testing.T) {
	controller := QuizController{
		QuizLoader: fixtures.MockQuizLoaderReturnsBasicQuiz{},
	}
	ts := httptest.NewServer(http.HandlerFunc(controller.GetQuiz))
	defer ts.Close()

	r, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal("Expected no Error")
	}
	if r.StatusCode != 400 {
		log.Fatalf("Expected Request to equal 400 Got: %v", r.StatusCode)
	}
}

func TestControllerReturns404IfNotA200Request(t *testing.T) {
	controller := QuizController{
		QuizLoader: fixtures.MockQuizLoaderReturnsBasicQuiz{},
	}
	ts := httptest.NewServer(http.HandlerFunc(controller.GetQuiz))
	defer ts.Close()

	r, err := http.Head(ts.URL)
	if err != nil {
		log.Fatal("Expected no Error")
	}
	if r.StatusCode != 405 {
		log.Fatal("Expected Request to equal 405")
	}
}

func TestControllerReturns404ForPostIfNotA200Request(t *testing.T) {
	controller := QuizController{
		QuizLoader: fixtures.MockQuizLoaderReturnsBasicQuiz{},
	}
	ts := httptest.NewServer(http.HandlerFunc(controller.PostAnswers))
	defer ts.Close()

	r, err := http.Head(ts.URL)
	if err != nil {
		log.Fatal("Expected no Error")
	}
	if r.StatusCode != 405 {
		log.Fatal("Expected Request to equal 405")
	}
}

func TestControllerUsesIdFromQueryString(t *testing.T) {
	mock := fixtures.MockQuizLoaderAccessParams{}
	controller := QuizController{
		QuizLoader: mock,
	}
	ts := httptest.NewServer(http.HandlerFunc(controller.GetQuiz))
	defer ts.Close()

	r, err := http.Get(ts.URL + "?id=test")
	if err != nil {
		log.Fatal("Expected no Error")
	}
	if r.StatusCode != 200 {
		log.Fatalf("Expected Request to equal 200 got: %v", r.StatusCode)
	}
}

func TestReturns404IfQuizIsntValid(t *testing.T) {
	mock := fixtures.MockQuizLoaderAccessParamsEmpty{}
	controller := QuizController{
		QuizLoader: mock,
	}
	ts := httptest.NewServer(http.HandlerFunc(controller.GetQuiz))
	defer ts.Close()

	r, err := http.Get(ts.URL + "?id=test")
	if err != nil {
		log.Fatal("Expected no Error")
	}
	if r.StatusCode != 404 {
		log.Fatalf("Expected Request to equal 404 got: %v", r.StatusCode)
	}
}

func TestControllerReturnsErrorFromServer(t *testing.T) {
	mock := fixtures.MockError{}
	controller := QuizController{
		QuizLoader: mock,
	}
	ts := httptest.NewServer(http.HandlerFunc(controller.GetQuiz))
	defer ts.Close()

	r, err := http.Get(ts.URL + "?id=test")
	if err != nil {
		log.Fatal("Expected no Error")
	}
	if r.StatusCode != 500 {
		log.Fatalf("Expected Request to equal 500 got: %v", r.StatusCode)
	}
}

func BenchmarkQuizEndpoint(b *testing.B) {
	quizController := QuizController{
		QuizLoader: data.CreateDynamoDataLoader(),
	}
	ts := httptest.NewServer(http.HandlerFunc(quizController.GetQuiz))
	defer ts.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		http.Get(ts.URL + "/quiz?id=test")
	}
}

func TestControllerReturnsErrorIfJsonParseFailed(t *testing.T) {
	mock := &fixtures.MockSaver{}
	controller := QuizController{
		QuizSaver: mock,
	}
	ts := httptest.NewServer(http.HandlerFunc(controller.PostAnswers))
	defer ts.Close()

	r, err := http.Post(ts.URL, "appliction/json", strings.NewReader(""))
	if err != nil {
		log.Fatal("Expected no Error")
	}
	if r.StatusCode != 400 {
		log.Fatal("Expected Request to equal 400")
	}
}

func TestControllerReturnsErrorIfResponseIsNotValid(t *testing.T) {
	mock := &fixtures.MockSaver{}
	controller := QuizController{
		QuizSaver: mock,
	}
	ts := httptest.NewServer(http.HandlerFunc(controller.PostAnswers))
	defer ts.Close()

	r, err := http.Post(ts.URL, "appliction/json", strings.NewReader("{\"id\":\"hello\"}"))
	if err != nil {
		log.Fatal("Expected no Error")
	}
	if r.StatusCode != 400 {
		log.Fatal("Expected Request to equal 400")
	}
}

func TestControllerReturnsSuccessCallingSaveWithValidBody(t *testing.T) {
	mock := &fixtures.MockSaver{}
	controller := QuizController{
		QuizSaver: mock,
	}
	ts := httptest.NewServer(http.HandlerFunc(controller.PostAnswers))
	defer ts.Close()

	r, err := http.Post(ts.URL, "appliction/json", strings.NewReader("{\"id\":\"hello\", \"answers\": [{\"questionId\": \"question1\", \"answer\": \"shearer\"}]}"))
	if err != nil {
		log.Fatal("Expected no Error")
	}
	if r.StatusCode != 200 {
		log.Fatal("Expected Request to equal 200")
	}
	if mock.Params == nil {
		log.Fatal("Expected params to not equal nil")
	}
	if mock.Params.Id != "hello" {
		log.Fatal("Expected Id to equal hello got %v", mock.Params.Id)
	}
	if len(mock.Params.Answers) != 1 {
		log.Fatal("Expected To Have 1 Answer")
	}
}

func BenchmarkAnswerEndpoint(b *testing.B) {
	quizController := QuizController{
		QuizSaver: data.CreateDynamoDataSaver(),
	}
	ts := httptest.NewServer(http.HandlerFunc(quizController.PostAnswers))
	defer ts.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		http.Post(ts.URL, "appliction/json", strings.NewReader("{\"id\":\"hello\", \"answers\": [{\"questionId\": \"question1\", \"answer\": \"shearer\"}]}"))
	}
}
