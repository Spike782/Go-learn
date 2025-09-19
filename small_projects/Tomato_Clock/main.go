package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

func formatTime(seconds int) string {
	min := seconds / 60
	sec := seconds % 60
	return fmt.Sprintf("%02d:%02d", min, sec)
}

func main() {
	myAPP := app.New()
	myWindow := myAPP.NewWindow("番茄钟")

	totalSecond := 25 * 60
	remainingSecond := totalSecond
	isrunning := false
	var timer *time.Ticker

	timeLabel := widget.NewLabel(formatTime(remainingSecond))

	var startBtn *widget.Button
	startBtn = widget.NewButton("开始", func() {
		if !isrunning {
			isrunning = true
			startBtn.SetText("暂停")
			timer = time.NewTicker(1 * time.Second)

			go func() {
				for range timer.C {
					if remainingSecond <= 0 {
						timer.Stop()
						isrunning = false
						startBtn.SetText("开始")
						break
					}
					remainingSecond--
					myWindow.Content().Refresh()
					timeLabel.SetText(formatTime(remainingSecond))
				}
			}()

		} else {
			isrunning = false
			startBtn.SetText("开始")
			timer.Stop()
		}
	})
	resetBtn := widget.NewButton("重置", func() {
		if timer != nil {
			timer.Stop()
		}
		isrunning = false
		remainingSecond = totalSecond
		timeLabel.SetText(formatTime(remainingSecond))
	})

	content := container.NewVBox(timeLabel, container.NewHBox(startBtn, resetBtn))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
