package notifier

import (
	"HeadHunter_Scaner/client"
	"HeadHunter_Scaner/config"
	"HeadHunter_Scaner/handler"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func TriggerAlert(vacancies *[]client.Vacancy) {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Printf("🔥 НАЙДЕНО %d НОВЫХ ВАКАНСИЙ! 🔥\n", len(*vacancies))
	for i := range *vacancies {
		fmt.Println((*vacancies)[i].Url)

		go func(text string) {
			sendNotificationToTelegram(text)
		}((*vacancies)[i].Url)

		go func() {
			playSoundNotify()
		}()
	}
	fmt.Println(time.Now().Format("15:04:05"))
	fmt.Println(strings.Repeat("=", 50))
}

func playSoundNotify() {
	if runtime.GOOS == "darwin" {
		exec.Command("osascript", "-e", `beep`).Run() // macOS
	} else {
		fmt.Print("\a") // Windows
	}
}

func sendNotificationToTelegram(text string) error {
	tempConfig, err := config.GetConfigurartion()
	if err != nil {
		return fmt.Errorf("ошибка получения конфигурации: %v", err)
	}

	params := []config.UrlParameter{
		{
			Key:   "chat_id",
			Value: tempConfig.Telegram.ChatId,
		},
		{
			Key:   "text",
			Value: text,
		},
	}

	buildedUrl := tempConfig.Telegram.BaseUrl + tempConfig.Telegram.BotToken + "/sendMessage"

	_, err = handler.Get(buildedUrl, &params)
	if err != nil {
		return fmt.Errorf("ошибка отправки уведомления: %v", err)
	}

	return nil
}
