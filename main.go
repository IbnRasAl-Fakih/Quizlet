package main

import (
	"fmt"
	"math/rand"
)

func main() {
	questions := []MultipleChoiceQuestion{
		{
			Question:           "Какой газ является основным составляющим атмосферы Земли?",
			Choices:            []string{"Кислород", "Азот", "Углекислый газ", "Аргон"},
			CorrectAnswerIndex: 1, // Индекс правильного ответа: Азот
		},
		{
			Question:           "Какое животное является символом США?",
			Choices:            []string{"Орёл", "Лев", "Сова", "Белка"},
			CorrectAnswerIndex: 0, // Индекс правильного ответа: Орёл
		},
		{
			Question:           "Какое из этих известных произведений принадлежит Шекспиру?",
			Choices:            []string{"Преступление и наказание", "Гамлет", "Война и мир", "Гарри Поттер"},
			CorrectAnswerIndex: 1, // Индекс правильного ответа: Гамлет
		},
	}

	trueFalseQuestions := []TrueFalseQuestion{
		{
			Question:      "Солнце вращается вокруг Земли.",
			CorrectAnswer: false, // Ложь
		},
		{
			Question:      "Вода состоит из двух атомов водорода и одного атома кислорода.",
			CorrectAnswer: true, // Истина
		},
		{
			Question:      "Москва - столица Франции.",
			CorrectAnswer: false, // Ложь
		},
	}

	openEndedQuestions := []OpenEndedQuestion{
		{
			Question:      "Как называется самая большая планета в Солнечной системе?",
			CorrectAnswer: "Юпитер",
		},
		{
			Question:      "Что такое формула Эйнштейна E=mc² представляет?",
			CorrectAnswer: "Эквивалент массы и энергии",
		},
		{
			Question:      "Как называется столица Японии?",
			CorrectAnswer: "Токио",
		},
	}

	fmt.Print("Добро пожаловать в приложение исторических знаний! \nПожалуйста, представьтесь, чтобы мы могли обращаться к вам по имени: ")
	var username string
	fmt.Scanln(&username)
	user := User{username, 0}
	fmt.Println("Здравствуйте, ", user.name, "! Здесь вы найдете увлекательные исторические факты, события и персоналии. \nИсследуйте разные эпохи и культуры, изучайте важные события и знаковые личности.\nМы гарантируем, что вы узнаете много нового и увлекательного о мировой истории.")

	for {
		fmt.Println("Выберите разновидность теста: \n1. Классика \n2. С временем \n ")
		var name = fmt.Println("Выберите тип теста: \n1. True-false \n2. Вопрос-ответ \n3. MCQ \n4. Счет \n5. Выйти")
		var choice int
		fmt.Scanln(&choice)
		quiz := &classicalQuiz{}
		switch choice {
		case 1:
			randomNum := rand.Intn(len(trueFalseQuestions))
			quiz.setQuestionType(NewTrueFalseQuestion(trueFalseQuestions[randomNum].Question, trueFalseQuestions[randomNum].CorrectAnswer))
			fmt.Print(quiz.AskQuestion())
			var answer string
			fmt.Scanln(&answer)
			score := quiz.EvaluateAnswer(answer)
			fmt.Println("Заработанные баллы", score)
			user.score += score
		case 2:
			randomNum := rand.Intn(len(openEndedQuestions))
			quiz.setQuestionType(NewOpenEndedQuestionProxy(&openEndedQuestions[randomNum], &user))
			res := quiz.AskQuestion()
			fmt.Print(res)
			if res == "Извините, у вас недостаточно баллов, чтобы пройти тест. Пожалуйста, наберите больше баллов, чтобы продолжить." {
				break
			}
			var answer string
			fmt.Scanln(&answer)
			score := quiz.EvaluateAnswer(answer)
			fmt.Println("Заработанные баллы", score)
		case 3:
			randomNum := rand.Intn(len(questions))
			quiz.setQuestionType(NewMultipleChoiceQuestionProxy(&questions[randomNum], &user))
			res := quiz.AskQuestion()
			fmt.Print(res)
			if res == "Извините, у вас недостаточно баллов, чтобы пройти тест. Пожалуйста, наберите больше баллов, чтобы продолжить." {
				break
			}
			var answer string
			fmt.Scanln(&answer)
			score := quiz.EvaluateAnswer(answer)
			fmt.Println("Заработанные баллы", score)
		case 4:
			fmt.Println("У вас на счету", user.score, "баллов")
		case 5:
			return
		default:
			fmt.Println("Вы ввели неправильное значение")
		}
	}
}
