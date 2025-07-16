package entity

import (
	"time"
)

type ListItem struct {
	ID            string    // Identificador único do item
	ShoppingListID string    // ID da lista de compras
	Name          string    // Nome do item
	Quantity      int       // Quantidade
	CategoryID    string    // ID da categoria
	Status        string    // Status: pendente, comprado, etc.
	CreatedAt     time.Time // Data de criação
} 