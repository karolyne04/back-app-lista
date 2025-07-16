package entity

import (
	"time"
)

type ShoppingList struct {
	ID         string      // Identificador único da lista
	UserID     string      // ID do usuário dono da lista
	Name       string      // Nome da lista
	Items      []*ListItem // Itens da lista
	CreatedAt  time.Time   // Data de criação
	UpdatedAt  time.Time   // Data de atualização
	FinishedAt *time.Time  // Data de finalização (opcional, para histórico)
} 