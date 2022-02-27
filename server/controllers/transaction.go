package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
	"gorm.io/gorm/clause"

	"github.com/MatejaMaric/spending-tracker/server/models"
	"github.com/MatejaMaric/spending-tracker/server/parser"
)

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	results := models.DB.Order("made_at desc").Find(&transactions)

	c.JSON(http.StatusOK, gin.H{
		"amount": results.RowsAffected,
		"data":   transactions,
	})
}

func CreateTransaction(c *gin.Context) {
	table, err := html.Parse(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed to parse data...",
			"error":  err,
		})
		return
	}

	transactions := parser.ExtractTransactions(table)
	n := models.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&transactions).RowsAffected

	c.JSON(http.StatusOK, gin.H{
		"added_rows": n,
	})
}
