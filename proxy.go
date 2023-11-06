package main

type IProxy interface {
	QuestionStrategy
}

type OpenEndedQuestionProxy struct {
	question *OpenEndedQuestion
	user     *User
	minScore int
}

func NewOpenEndedQuestionProxy(question *OpenEndedQuestion, user *User) *OpenEndedQuestionProxy {
	return &OpenEndedQuestionProxy{
		question: question,
		user:     user,
		minScore: 1,
	}
}

func (o *OpenEndedQuestionProxy) AskQuestion() string {
	if o.user.score > o.minScore {
		return o.question.AskQuestion()
	}
	return "Извините, у вас недостаточно баллов, чтобы пройти тест. Пожалуйста, наберите больше баллов, чтобы продолжить."
}

func (o *OpenEndedQuestionProxy) EvaluateAnswer(answer string) int {
	if o.user.score > o.minScore {
		score := o.question.EvaluateAnswer(answer)
		if score > 0 {
			o.user.score += score
		}
		return score
	}
	return 0
}

type MultipleChoiceQuestionProxy struct {
	question *MultipleChoiceQuestion
	user     *User
	minScore int
}

func NewMultipleChoiceQuestionProxy(question *MultipleChoiceQuestion, user *User) *MultipleChoiceQuestionProxy {
	return &MultipleChoiceQuestionProxy{
		question: question,
		user:     user,
		minScore: 1,
	}
}

func (m *MultipleChoiceQuestionProxy) AskQuestion() string {
	if m.user.score > m.minScore {
		return m.question.AskQuestion()
	}
	return "Извините, у вас недостаточно баллов, чтобы пройти тест. Пожалуйста, наберите больше баллов, чтобы продолжить."

}

func (m *MultipleChoiceQuestionProxy) EvaluateAnswer(answer string) int {
	if m.user.score > m.minScore {
		score := m.question.EvaluateAnswer(answer)
		if score > 0 {
			m.user.score += score
		}
		return score
	}
	return 0
}
