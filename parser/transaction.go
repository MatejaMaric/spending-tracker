package parser

import (
	"log"
	"strings"
	"time"

	"golang.org/x/net/html"

	"github.com/MatejaMaric/spending-tracker/models"
)

func ExtractTransactions(n *html.Node) []models.Transaction {
	var transactions []models.Transaction
	table := ExtractText(n)

	if len(table) == 0 {
		log.Println("No text extracted...")
		return transactions
	}

	for _, row := range table[1 : len(table)-1] {
		var t models.Transaction

		if len(row) < 6 {
			log.Printf("Skipped extracted row... Number of row items is: %d\n", len(row))
			log.Println(row)
			continue
		}

		var err error
		t.MadeAt, err = time.Parse("02.01.2006.", row[0])
		t.ProcessedAt, err = time.Parse("02.01.2006.", row[1])
		t.Description = strings.TrimSpace(row[2])
		t.Paid, err = ParseCommaFloat(row[3])
		t.Received, err = ParseCommaFloat(row[4])
		t.Balance, err = ParseCommaFloat(row[5])
		if err != nil {
			log.Println("Error parsing values...")
			continue
		}

		t.GenerateHash()

		transactions = append(transactions, t)
	}

	return transactions
}
