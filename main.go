package main

import (
	"log"

	"github.com/MatejaMaric/spending-tracker/gui"
	"github.com/MatejaMaric/spending-tracker/models"
	"github.com/adrg/xdg"
)

func main() {
	dbPath, err := xdg.DataFile("spending-tracker/data.db")
	if err != nil {
		log.Fatal(err)
	}

	models.Connect(dbPath)

	gui.Run()
}
