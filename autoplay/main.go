package main

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var firstClick bool

func main() {
	rows := 16
	columns := 16
	mines := 40

	rand.Seed(time.Now().UnixNano())
	myApp := app.New()
	myWindow := myApp.NewWindow("MINES")

	board := EmptyBoard(rows, columns)
	content := container.NewGridWithColumns(len(board[0]))
	board.SetButtonLog(content, mines)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
