// Package notifier отвечает за уведомления о новых вакансиях
package notifier

import (
	"fmt"
	"hhscaner/configuration"
	"hhscaner/service/httphandler"
	"strings"
	"time"
)

// TriggerAlert выводит уведомление о новых вакансиях в консоль
func TriggerAlert(vacanciesIds []int64) {

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Printf("🔥 НАЙДЕНО %d НОВЫХ ВАКАНСИЙ! 🔥\n", len(vacanciesIds))
	fmt.Println(time.Now().Format("15:04:05"))
	fmt.Println(strings.Repeat("=", 50))
}

// SendNotificationToTelegram отправляет уведомление в Telegram
func SendNotificationToTelegram(
	config *configuration.Config,
	client httphandler.HttpClient,
	text string) error {

	params := []configuration.UrlParameter{
		{
			Key:   "chat_id",
			Value: config.Telegram.ChatId,
		},
		{
			Key:   "text",
			Value: text,
		},
	}

	buildedUrl := config.Telegram.ApiUrl + config.Telegram.BotToken + "/sendMessage"

	_, err := client.Get(buildedUrl, &params)
	if err != nil {
		return fmt.Errorf("не удалось отправить сообщение в Telegram: %v", err)
	}

	return nil
}
