package main

type QuizProxy struct {
	quiz     QuestionStrategy
	user     *User
	minScore int
}

func (q *QuizProxy) Proxy(quiz QuestionStrategy, user *User, minscore int) {
	q.quiz = quiz
	q.user = user
	q.minScore = minscore
}

func (q *QuizProxy) askQuestion() string {
	if q.user.score > q.minScore {
		return q.quiz.askQuestion()
	}
	return "Извините, у вас недостаточно баллов, чтобы пройти тест. Пожалуйста, наберите больше баллов, чтобы продолжить."
}

func (q *QuizProxy) evaluateAnswer() int {
	if q.user.score > q.minScore {
		score := q.quiz.evaluateAnswer()
		if score > 0 {
			q.user.score += score
		}
		return score
	}
	return 0
}
