package services

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

	var n int64
	transactions := parser.ExtractTransactions(table)
	if len(transactions) > 0 {
		n = models.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&transactions).RowsAffected
	}

	return n, nil
}

func UpdateTransactionPersonalDescription(ID []byte, personalDescription string) error {
	return models.DB.Model(&models.Transaction{}).Where("id = ?", ID).Update("personal_description", personalDescription).Error
}
