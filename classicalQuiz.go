package main

type classicalQuiz struct {
	questionType QuestionStrategy
}

func (c *classicalQuiz) setQuestionType(questionType QuestionStrategy) {
	c.questionType = questionType
}

func (c *classicalQuiz) AskQuestion() string {
	return c.questionType.AskQuestion()
}

func (c *classicalQuiz) EvaluateAnswer(answer string) int {
	return c.questionType.EvaluateAnswer(answer)
}
