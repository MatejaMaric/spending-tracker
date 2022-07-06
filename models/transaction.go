package models

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Transaction struct {
	ID                  []byte       `gorm:"primaryKey"`
	MadeAt              time.Time    `gorm:"type:date not null"`
	ProcessedAt         time.Time    `gorm:"type:date not null"`
	Description         string       `gorm:"type:text not null"`
	Paid                float64      `gorm:"type:numeric(15,2) not null"`
	Received            float64      `gorm:"type:numeric(15,2) not null"`
	Balance             float64      `gorm:"type:numeric(15,2) not null"`
	PersonalDescription string       `gorm:"type:text"`
	OrderNumber         int          ``
	Link                *Transaction `gorm:"-"`
}

func (t Transaction) String() string {
	return fmt.Sprintf(
		"%s %s %s %g %g %g",
		t.MadeAt.Format("2006-01-02"),
		t.ProcessedAt.Format("2006-01-02"),
		t.Description,
		t.Paid,
		t.Received,
		t.Balance,
	)
}

func (t *Transaction) GenerateHash() {
	bytes := []byte(t.String())
	sum := sha256.Sum256(bytes)
	t.ID = sum[:]
}

func SortTransactions(transactions []Transaction) ([]Transaction, error) {

	newBalance := map[float64]*Transaction{}
	for i := range transactions {
		t := &transactions[i]
		if _, exists := newBalance[t.Paid+t.Balance]; exists {
			return nil, fmt.Errorf("duplicate balance")
		}
		newBalance[t.Paid+t.Balance] = t
	}

	var t *Transaction
	for i := range transactions {
		o := &transactions[i]
		n, ok := newBalance[o.Balance]
		if !ok {
			t = o
			continue
		}
		n.Link = o
	}

	for i := 1; t != nil; i++ {
		t.OrderNumber = i
		t = t.Link
	}

	return transactions, nil
}
