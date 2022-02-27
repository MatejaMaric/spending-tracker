package parser

import (
	"strings"
	"time"

	"golang.org/x/net/html"

	"github.com/MatejaMaric/spending-tracker/server/models"
)

func ExtractTransactions(n *html.Node) []models.Transaction {
	var transactions []models.Transaction
	table := ExtractText(n)

	for _, row := range table[1 : len(table)-1] {
		var t models.Transaction

		t.MadeAt, _ = time.Parse("02.01.2006.", row[0])
		t.ProcessedAt, _ = time.Parse("02.01.2006.", row[1])
		t.Description = strings.TrimSpace(row[2])
		t.Paid, _ = ParseCommaFloat(row[3])
		t.Received, _ = ParseCommaFloat(row[4])
		t.Balance, _ = ParseCommaFloat(row[5])
		t.GenerateHash()

		transactions = append(transactions, t)
	}

	return transactions
}
