package main

import (
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//Структура для пользователей, реализация позднее
type users struct {
	IdUser       int
	LoginUser    string
	PasswordUSer string
	Loggined     bool
}

type statsUser struct {
	IdUser       int
	Level        int
	AmountRaunds int
	WinnerRaunds int
}

var Player users
var PlayerStats statsUser

func createNumber(lvl int, login bool) string {
	rand.Seed(time.Now().Unix())
	b := ""
	count := lvl
	for i := 0; i < count; i++ {
		b += strconv.Itoa(rand.Intn(9))
	}

	return b
}

func main() {
	count := 0
	Num := ""
	gamestart := false
	firstplay := true
	btnStartText := "Старт игры"
	linetext := ""

	Player.LoginUser = "khisamov"
	Player.PasswordUSer = "123"
	Player.IdUser = 1
	Player.Loggined = true
	PlayerStats.IdUser = 1
	PlayerStats.Level = 1
	PlayerStats.AmountRaunds = 5
	PlayerStats.WinnerRaunds = 5

	a := app.New()
	w := a.NewWindow("Тренажер для мозга V1")

	number := canvas.NewText(linetext, color.Black)
	number.TextSize = 50
	number.Alignment = fyne.TextAlignCenter
	lable := widget.NewLabel("Введите число")

	ent := widget.NewEntry()
	btnEnter := widget.NewButton("Отправить", func() {
		if !ent.Disabled() {
			count++
			if ent.Text == Num {
				upd(number, "Верно!")
				time.Sleep(time.Second)
				if count%3 == 0 {
					PlayerStats.Level++
					upd(number, "Левел UP!")
					time.Sleep(time.Second)
					upd(number, "теперь будет еще сложнее")
					time.Sleep(time.Second)
				}

			} else {

				upd(number, "Неверно!")
				time.Sleep(time.Second)

				if PlayerStats.Level > 2 {
					PlayerStats.Level--
					upd(number, "твой уровень понижен!")
				} else {
					upd(number, "твой уровень")
					time.Sleep(time.Second)
					upd(number, "настолько мал")
					time.Sleep(time.Second)
					upd(number, "я даже не могу")
					time.Sleep(time.Second)
					upd(number, "его понизить...")
				}
				time.Sleep(time.Second)

			}

			upd(number, "Нажимай старт!")
			ent.Disable()
			ent.SetText("")
			time.Sleep(time.Second)
		} else {
			upd(number, "Хватит хитрить")
			time.Sleep(time.Second)
			upd(number, "ты не нажимал")
			time.Sleep(time.Second)
			upd(number, "на старт")
			time.Sleep(time.Second)
			upd(number, "")
		}

	})

	ent.Disable()
	btnStart := widget.NewButton(btnStartText, func() {
		if firstplay == true {
			upd(number, "Сейчас я буду")
			time.Sleep(time.Second * 2)
			upd(number, "называть числа")
			time.Sleep(time.Second * 2)
			upd(number, "твоя задача")
			time.Sleep(time.Second)
			upd(number, "запомнить их")
			time.Sleep(time.Second)
			upd(number, "поехали!")
			time.Sleep(time.Second)
			firstplay = false
		}

		Num = createNumber(PlayerStats.Level, Player.Loggined)
		for _, item := range Num {
			upd(number, string(item))
			time.Sleep(time.Millisecond * 500)
			upd(number, "")
			time.Sleep(time.Second)
		}
		upd(number, "Скорее вводи число")
		ent.Enable()
		gamestart = true
	})

	line2 := container.NewAdaptiveGrid(3, lable, ent, btnEnter)

	content := container.NewVBox(number, line2, btnStart)
	contentwrap := container.NewCenter(content)
	w.SetContent(contentwrap)
	if gamestart == true {
		btnStart.Disable()
	}
	w.ShowAndRun()

}

func upd(ndsx *canvas.Text, txt string) {
	ndsx.Text = txt
	ndsx.Refresh()
}
