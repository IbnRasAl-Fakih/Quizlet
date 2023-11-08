package main

type classicalQuiz struct {
	questionType QuestionStrategy
}

func (c *classicalQuiz) setQuestionType(questionType QuestionStrategy) {
	c.questionType = questionType
}

func (c *classicalQuiz) askQuestion() string {
	return c.questionType.askQuestion()
}

func (c *classicalQuiz) evaluateAnswer() int {
	return c.questionType.evaluateAnswer()
}
