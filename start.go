package main

import (
	"fmt"
	"github.com/lib/pq"
	"math/rand"
)

func quizStart() {
	questions := []MultipleChoiceQuestion{}
	trueFalseQuestions := []TrueFalseQuestion{}
	openEndedQuestions := []OpenEndedQuestion{}

	db := GetDatabaseInstance()
	defer db.connection.Close()

	rows, err1 := db.connection.Query(`SELECT * FROM openended`)
	if err1 != nil {
		panic(err1)
	}
	for rows.Next() {
		var question string
		var answer string
		if err := rows.Scan(&question, &answer); err != nil {
			panic(err)
		}
		openEndedQuestions = append(openEndedQuestions, OpenEndedQuestion{question, answer})
	}
	defer rows.Close()

	rows2, err2 := db.connection.Query(`SELECT * FROM mcq`)
	if err2 != nil {
		panic(err2)
	}
	for rows2.Next() {
		var question string
		var choices pq.StringArray
		var answer int
		if err := rows2.Scan(&question, &choices, &answer); err != nil {
			panic(err)
		}
		questions = append(questions, MultipleChoiceQuestion{question, []string(choices), answer})
	}
	defer rows2.Close()

	rows3, err3 := db.connection.Query(`SELECT * FROM truefalse`)
	if err3 != nil {
		panic(err3)
	}
	for rows3.Next() {
		var question string
		var answer bool
		if err := rows3.Scan(&question, &answer); err != nil {
			panic(err)
		}
		trueFalseQuestions = append(trueFalseQuestions, TrueFalseQuestion{question, answer})
	}
	defer rows3.Close()

	fmt.Print("Добро пожаловать в приложение исторических знаний! \nПожалуйста, представьтесь, чтобы мы могли обращаться к вам по имени: ")
	var username string
	fmt.Scanln(&username)
	user := User{username, 0}
	fmt.Printf("Здравствуйте, %s! Здесь вы найдете увлекательные исторические факты, события и персоналии. \nИсследуйте разные эпохи и культуры, изучайте важные события и знаковые личности.\nМы гарантируем, что вы узнаете много нового и увлекательного о мировой истории.", user.name)
	for {
		fmt.Println("Выберите разновидность теста: \n1. Классика \n2. С временем \n3. Счет \n4. Выйти")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
		LOOP1:
			for {
				fmt.Println()
				fmt.Println("Выберите тип теста: \n1. True-false \n2. Вопрос-ответ \n3. MCQ \n4. Счет \n5. Выйти")
				var choice2 int
				fmt.Scanln(&choice2)
				quiz := &classicalQuiz{}
				switch choice2 {
				case 1:
					randomNum := rand.Intn(len(trueFalseQuestions))
					quiz.setQuestionType(NewTrueFalseQuestion(trueFalseQuestions[randomNum].Question, trueFalseQuestions[randomNum].CorrectAnswer))
					fmt.Println()
					fmt.Print(quiz.askQuestion())
					score := quiz.evaluateAnswer()
					fmt.Println("Заработанные баллы", score)
					user.score += score
				case 2:
					randomNum := rand.Intn(len(openEndedQuestions))
					quiz.setQuestionType(NewOpenEndedQuestion(openEndedQuestions[randomNum].Question, openEndedQuestions[randomNum].CorrectAnswer))
					quizProxy := Proxy(quiz, &user, 5)
					fmt.Println()
					res := quizProxy.askQuestion()
					if res == "Error" {
						fmt.Println("Извините, у вас недостаточно баллов, чтобы пройти тест. Пожалуйста, наберите больше баллов, чтобы продолжить.")
						break
					}
					fmt.Print(res)
					score := quizProxy.evaluateAnswer()
					fmt.Println("Заработанные баллы", score)
				case 3:
					randomNum := rand.Intn(len(questions))
					quiz.setQuestionType(NewMultipleChoiceQuestion(questions[randomNum].Question, questions[randomNum].Choices, questions[randomNum].CorrectAnswerIndex))
					quizProxy := Proxy(quiz, &user, 3)
					fmt.Println()
					res := quizProxy.askQuestion()
					if res == "Error" {
						fmt.Println("Извините, у вас недостаточно баллов, чтобы пройти тест. Пожалуйста, наберите больше баллов, чтобы продолжить.")
						break
					}
					fmt.Print(res)
					score := quizProxy.evaluateAnswer()
					fmt.Println("Заработанные баллы", score)
				case 4:
					fmt.Println("У вас на счету", user.score, "баллов")
				case 5:
					break LOOP1
				default:
					fmt.Println("Вы ввели неправильное значение")
				}
			}
		case 2:
		LOOP2:
			for {
				fmt.Println()
				fmt.Println("Выберите тип теста: \n1. True-false \n2. Вопрос-ответ \n3. MCQ \n4. Счет \n5. Выйти")
				var choice2 int
				fmt.Scanln(&choice2)
				quiz := &classicalQuiz{}
				switch choice2 {
				case 1:
					randomNum := rand.Intn(len(trueFalseQuestions))
					quiz.setQuestionType(NewTrueFalseQuestion(trueFalseQuestions[randomNum].Question, trueFalseQuestions[randomNum].CorrectAnswer))
					quizWithTime := TimerDecorator{quiz, 7}
					fmt.Println()
					fmt.Print(quizWithTime.askQuestion())
					score := quizWithTime.evaluateAnswer()
					fmt.Println("Заработанные баллы", score)
					user.score += score
				case 2:
					randomNum := rand.Intn(len(openEndedQuestions))
					quiz.setQuestionType(NewOpenEndedQuestion(openEndedQuestions[randomNum].Question, openEndedQuestions[randomNum].CorrectAnswer))
					quizWithTime := TimerDecorator{quiz, 12}
					quizProxy := Proxy(&quizWithTime, &user, 3)
					fmt.Println()
					res := quizProxy.askQuestion()
					if res == "Error" {
						fmt.Println("Извините, у вас недостаточно баллов, чтобы пройти тест. Пожалуйста, наберите больше баллов, чтобы продолжить.")
						break
					}
					fmt.Print(res)
					score := quizProxy.evaluateAnswer()
					fmt.Println("Заработанные баллы", score)
				case 3:
					randomNum := rand.Intn(len(questions))
					quiz.setQuestionType(NewMultipleChoiceQuestion(questions[randomNum].Question, questions[randomNum].Choices, questions[randomNum].CorrectAnswerIndex))
					quizWithTime := TimerDecorator{quiz, 10}
					quizProxy := Proxy(&quizWithTime, &user, 5)
					fmt.Println()
					res := quizProxy.askQuestion()
					if res == "Error" {
						fmt.Println("Извините, у вас недостаточно баллов, чтобы пройти тест. Пожалуйста, наберите больше баллов, чтобы продолжить.")
						break
					}
					fmt.Print(res)
					score := quizProxy.evaluateAnswer()
					fmt.Println("Заработанные баллы", score)
				case 4:
					fmt.Println("У вас на счету", user.score, "баллов")
				case 5:
					break LOOP2
				default:
					fmt.Println("Вы ввели неправильное значение")
				}
			}
		case 3:
			fmt.Println("У вас на счету", user.score, "баллов")
		case 4:
			return
		default:
			fmt.Println("Вы ввели неправильное значение")
		}
	}
}
