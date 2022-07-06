package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/MatejaMaric/spending-tracker/models"
	"github.com/adrg/xdg"
)

func main() {
	dbPath, err := xdg.DataFile("spending-tracker/data.db")
	if err != nil {
		log.Fatal(err)
	}

	models.Connect(dbPath)

	a := app.New()
	w := a.NewWindow("Spending Tracker")

	w.SetContent(widget.NewLabel("DB path is: " + dbPath))
	w.ShowAndRun()
}
