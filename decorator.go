package main

import (
	"fmt"
	"strconv"
	"time"
)

type QuestionDecorator interface {
	QuestionStrategy
	askQuestion() string
	evaluateAnswer() int
}

type TimerDecorator struct {
	decorator QuestionStrategy
	timeLimit int
}

func (t *TimerDecorator) askQuestion() string {
	timerMessage := "У вас есть " + strconv.Itoa(t.timeLimit) + " секунд чтобы ответить на вопрос."
	return timerMessage + "\n" + t.decorator.askQuestion()
}

func (t *TimerDecorator) evaluateAnswer() int {
	t0 := time.Now()
	result := t.decorator.evaluateAnswer()
	t1 := time.Now()

	elapsed := t1.Sub(t0)
	if (elapsed) > time.Duration(t.timeLimit)*time.Second {
		fmt.Println("Вы не успели ответить вовремя")
		result = 0
	}

	if result == 1 {
		result++
	}

	return result
}
