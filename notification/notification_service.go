package notification

import (
	"HeadHunter_Scaner/model"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func TriggerAlert(vacancies *[]model.Vacancy) {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🔥 НАЙДЕНЫ НОВЫЕ ВАКАНСИИ! 🔥")
	for i := range *vacancies {
		fmt.Println((*vacancies)[i].Url)
	}
	fmt.Println(time.Now().Format("15:04:05"))
	fmt.Println(strings.Repeat("=", 50))

	if runtime.GOOS == "darwin" {
		playSoundAndNotify() // macOS
	} else {
		fmt.Print("\a") // Windows
	}
}

func playSoundAndNotify() {
	go func() {
		exec.Command("osascript", "-e", `beep`).Run()
	}()
	time.Sleep(time.Millisecond * 100)
}
