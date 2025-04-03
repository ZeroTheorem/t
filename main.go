package main

import (
	"fmt"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var successCount = 0
var failCount = 0

// Функция для вычисления процента успеха
func getSuccessRate() float64 {
	total := successCount + failCount
	if total == 0 {
		return 0.0
	}
	return (float64(successCount) / float64(total)) * 100
}

// Функция для формирования ответа бота
func handleMessage(message string) string {
	// Проверяем, что сообщение состоит из одного символа + или -
	switch message {
	case "+":
		successCount++
	case "-":
		failCount++
	default:
		return "Отправь только символы '+' или '-' для инкрементации успеха или провала."
	}

	// Формируем процент успеха
	successRate := getSuccessRate()

	// Возвращаем процент успеха и количество успехов и провалов
	return fmt.Sprintf("Успехов: %d, Провалов: %d\nПроцент успеха: %.2f%%", successCount, failCount, successRate)
}

func main() {
	// Создаем бота с вашим токеном
	token := "6951087608:AAGfVdrNduwKwnvUoE5JG_DbPKlc0OQLG9s"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// Получаем обновления для бота
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		// Игнорируем, если нет текста в сообщении
		if update.Message == nil {
			continue
		}

		// Получаем текст из сообщения
		msgText := update.Message.Text

		// Формируем ответ
		response := handleMessage(msgText)

		// Отправляем ответ
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		bot.Send(msg)
	}
}
