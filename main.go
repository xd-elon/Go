package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Estrutura do item
type Item struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Preco float64 `json:"preco"`
}

var items = []Item{
	{ID: "1", Nome: "Aqw", Preco: 10.5},
	{ID: "2", Nome: "Asa", Preco: 20.0},
}

func main() {
	r := gin.Default()

	// Endpoint para pegar todos os itens
	r.GET("/items", func(c *gin.Context) {
		c.JSON(http.StatusOK, items)
	})

	// Endpoint para criar um item novo
	r.POST("/items", func(c *gin.Context) {
		var newItem Item
		if err := c.ShouldBindJSON(&newItem); err == nil {
			items = append(items, newItem)
			c.JSON(http.StatusCreated, newItem)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// Iniciando o servidor
	r.Run(":8080")
}
