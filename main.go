package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

const msg = `
Wins: **%v**; Loses: **%v**

Win rate: **%v**
`

var winsCount = 22
var losesCount = 5

// Функция для вычисления процента успеха
func getSuccessRate() float64 {
	total := winsCount + losesCount
	if total == 0 {
		return 0.0
	}
	return (float64(winsCount) / float64(total)) * 100
}

func main() {
	pref := tele.Settings{
		Token:     os.Getenv("TOKEN"),
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeMarkdownV2,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	selector := &tele.ReplyMarkup{}
	btnWin := selector.Data("Win", "W")
	btnLose := selector.Data("Lose", "L")

	selector.Inline(
		selector.Row(btnWin, btnLose),
	)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello!", selector)
	})

	b.Handle(&btnWin, func(c tele.Context) error {
		winsCount++
		winRate := getSuccessRate()
		return c.Send(fmt.Sprintf(msg, winsCount, losesCount, winRate))
	})

	b.Handle(&btnLose, func(c tele.Context) error {
		losesCount++
		winRate := getSuccessRate()
		return c.Send(fmt.Sprintf(msg, winsCount, losesCount, winRate))
	})
	b.Start()
}
