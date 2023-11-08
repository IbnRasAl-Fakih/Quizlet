package main

type QuizProxy struct {
	quiz     QuestionStrategy
	user     *User
	minScore int
}

func Proxy(quiz QuestionStrategy, user *User, minScore int) *QuizProxy {
	return &QuizProxy{
		quiz:     quiz,
		user:     user,
		minScore: minScore,
	}
}

func (q *QuizProxy) askQuestion() string {
	if q.user.score >= q.minScore {
		return q.quiz.askQuestion()
	}
	return "Error"
}

func (q *QuizProxy) evaluateAnswer() int {
	if q.user.score >= q.minScore {
		score := q.quiz.evaluateAnswer()
		if score > 0 {
			q.user.score += score
		}
		return score
	}
	return 0
}
