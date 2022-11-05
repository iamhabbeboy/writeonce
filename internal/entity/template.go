package entity

import "github.com/google/uuid"

type Template struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Template    string    `json:"template"`
}
