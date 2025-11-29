// Package notifier отвечает за уведомления о новых вакансиях
package notifier

import (
	"fmt"
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
	client httphandler.HttpClient,
	urlString string,
	text string) error {

	_, err := client.Get(urlString)
	if err != nil {
		return fmt.Errorf("не удалось отправить сообщение в Telegram: %v", err)
	}

	return nil
}
