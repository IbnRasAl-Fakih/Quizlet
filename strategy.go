package main

import (
	"fmt"
	"strconv"
	"strings"
)

type QuestionStrategy interface {
	askQuestion() string
	evaluateAnswer() int
}

type MultipleChoiceQuestion struct {
	Question           string
	Choices            []string
	CorrectAnswerIndex int
}

func NewMultipleChoiceQuestion(question string, choices []string, answer int) *MultipleChoiceQuestion {
	return &MultipleChoiceQuestion{
		Question:           question,
		Choices:            choices,
		CorrectAnswerIndex: answer,
	}
}

func (m *MultipleChoiceQuestion) askQuestion() string {
	question := fmt.Sprintf("Вопрос: %s\n", m.Question)

	choices := ""
	for i, choice := range q.Choices {
		choices += fmt.Sprintf("%d. %s\n", i+1, choice)
	}

	return question + choices

}

func (m *MultipleChoiceQuestion) evaluateAnswer() int {
	num, err := strconv.Atoi(takeAnswer())

	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	if num == q.CorrectAnswerIndex+1 {
		return 1
	}

	return 0
}

type TrueFalseQuestion struct {
	Question      string
	CorrectAnswer bool
}

func NewTrueFalseQuestion(question string, answer bool) *TrueFalseQuestion {
	return &TrueFalseQuestion{
		Question:      question,
		CorrectAnswer: answer,
	}
}

func (t *TrueFalseQuestion) AskQuestion() string {
	return fmt.Sprintf("Вопрос: %s (True/False)\n", t.Question)
}

func (t *TrueFalseQuestion) EvaluateAnswer(answer string) int {
	answer = strings.ToLower(strings.TrimSpace(answer))

	correctAnswer := t.CorrectAnswer

	if answer == "true" && correctAnswer {
		return 1
	} else if answer == "false" && !correctAnswer {
		return 1
	}

	return 0
}

type OpenEndedQuestion struct {
	Question      string
	CorrectAnswer string
}

func NewOpenEndedQuestion(question string, answer string) *OpenEndedQuestion {
	return &OpenEndedQuestion{
		Question:      question,
		CorrectAnswer: answer,
	}
}

func (q *OpenEndedQuestion) AskQuestion() string {
	return fmt.Sprintf("Вопрос: %s\n", q.Question)
}

func (q *OpenEndedQuestion) evaluateAnswer() int {

	userAnswer := strings.ToLower(takeAnswer())
	correctAnswer := strings.ToLower(q.CorrectAnswer)

	if userAnswer == correctAnswer {
		return 1
	}

	return 0
}
