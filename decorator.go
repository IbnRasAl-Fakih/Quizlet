package main

import (
	"time"
)

type QuestionDecorator interface {
	AskQuestion() string
	EvaluateAnswer(answer string) int
}

type MultipleChoiceQuestionDecorator struct {
	question  *MultipleChoiceQuestion
	timeLimit int
}

func NewMultipleChoiceQuestionDecorator(question *MultipleChoiceQuestion, timeLimit int) *MultipleChoiceQuestionDecorator {
	return &MultipleChoiceQuestionDecorator{
		question:  question,
		timeLimit: timeLimit,
	}
}

func (m *MultipleChoiceQuestionDecorator) AskQuestion() string {
	return m.question.AskQuestion()
}

func (m *MultipleChoiceQuestionDecorator) EvaluateAnswer(answer string) int {
	startTime := time.Now()
	result := m.question.EvaluateAnswer(answer)
	endTime := time.Now()

	if endTime.Sub(startTime) > time.Duration(m.timeLimit)*time.Second {
		result = 0
	}

	return result
}

type OpenEndedQuestionDecorator struct {
	question  *OpenEndedQuestion
	timeLimit int
}

func NewOpenEndedQuestionDecorator(question *OpenEndedQuestion, timeLimit int) *OpenEndedQuestionDecorator {
	return &OpenEndedQuestionDecorator{
		question:  question,
		timeLimit: timeLimit,
	}
}

func (o *OpenEndedQuestionDecorator) AskQuestion() string {
	return o.question.AskQuestion()
}

func (o *OpenEndedQuestionDecorator) EvaluateAnswer(answer string) int {
	startTime := time.Now()
	result := o.question.EvaluateAnswer(answer)
	endTime := time.Now()

	if endTime.Sub(startTime) > time.Duration(o.timeLimit)*time.Second {
		result = 0
	}

	return result
}

type TrueFalseQuestionDecorator struct {
	question  *TrueFalseQuestion
	timeLimit int
}

func NewTrueFalseQuestionDecorator(question *TrueFalseQuestion, timeLimit int) *TrueFalseQuestionDecorator {
	return &TrueFalseQuestionDecorator{
		question:  question,
		timeLimit: timeLimit,
	}
}

func (t *TrueFalseQuestionDecorator) AskQuestion() string {
	return t.question.AskQuestion()
}

func (t *TrueFalseQuestionDecorator) EvaluateAnswer(answer string) int {
	startTime := time.Now()
	result := t.question.EvaluateAnswer(answer)
	endTime := time.Now()

	if endTime.Sub(startTime) > time.Duration(t.timeLimit)*time.Second {
		result = 0
	}

	return result
}
