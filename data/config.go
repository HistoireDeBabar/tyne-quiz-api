package data

const (
	QuestionTableName      = "Question"
	AnswerTableName        = "Answer"
	QuizQuestionQuery      = "quizId = :quizId"
	AnswerQuestionIdQuery  = "questionId = :questionId"
	AnswerUpdateExpression = "SET answerCount = answerCount + :inc"
)
