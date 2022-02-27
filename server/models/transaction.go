package models

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Transaction struct {
	ID          []byte    `gorm:"primaryKey"`
	MadeAt      time.Time `gorm:"type:date not null"`
	ProcessedAt time.Time `gorm:"type:date not null"`
	Description string    `gorm:"type:text not null"`
	Paid        float64   `gorm:"type:numeric(15,2) not null"`
	Received    float64   `gorm:"type:numeric(15,2) not null"`
	Balance     float64   `gorm:"type:numeric(15,2) not null"`
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
