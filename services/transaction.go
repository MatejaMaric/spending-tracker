package controllers

import (
	"strings"

	"golang.org/x/net/html"
	"gorm.io/gorm/clause"

	"github.com/MatejaMaric/spending-tracker/models"
	"github.com/MatejaMaric/spending-tracker/parser"
)

func GetTransactions() []models.Transaction {
	var transactions []models.Transaction
	models.DB.Order("made_at desc").Find(&transactions)
	return transactions
}

func CreateTransactions(text string) (addedRows int64, err error) {
	table, err := html.Parse(strings.NewReader(text))
	if err != nil {
		return 0, err
	}

	transactions := parser.ExtractTransactions(table)
	n := models.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&transactions).RowsAffected

	return n, nil
}
