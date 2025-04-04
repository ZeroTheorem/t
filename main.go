package main

import (
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v4"
)

const msg = `
Wins: <b>%v</b>; Loses: <b>%v</b>

Win rate: <b>%.2f%%</b>
`

var winsCount = 24
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
		Token:     "6951087608:AAGfVdrNduwKwnvUoE5JG_DbPKlc0OQLG9s",
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeHTML,
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
		selector.Row(btnWin),
		selector.Row(btnLose),
	)

	b.Handle("/start", func(c tele.Context) error {
		winRate := getSuccessRate()
		return c.Send(fmt.Sprintf(msg, winsCount, losesCount, winRate), selector)
	})

	b.Handle(&btnWin, func(c tele.Context) error {
		winsCount++
		winRate := getSuccessRate()
		return c.Send(fmt.Sprintf(msg, winsCount, losesCount, winRate), selector)
	})

	b.Handle(&btnLose, func(c tele.Context) error {
		losesCount++
		winRate := getSuccessRate()
		return c.Send(fmt.Sprintf(msg, winsCount, losesCount, winRate), selector)
	})
	b.Start()
}
