package notification

import (
	"HeadHunter_Scaner/client"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func TriggerAlert(vacancies *[]client.Vacancy) {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🔥 НАЙДЕНЫ НОВЫЕ ВАКАНСИИ! 🔥")
	for i := range *vacancies {
		fmt.Println((*vacancies)[i].Url)
	}
	fmt.Println(time.Now().Format("15:04:05"))
	fmt.Println(strings.Repeat("=", 50))

	playSoundAndNotify()
}

func playSoundAndNotify() {
	go func() {
		for range 5 {
			if runtime.GOOS == "darwin" {
				exec.Command("osascript", "-e", `beep`).Run() // macOS
			} else {
				fmt.Print("\a") // Windows
			}
			time.Sleep(time.Microsecond * 500)
		}
	}()
}
