package entity

import (
	"github.com/google/uuid"
	"github.com/theterminalguy/writeonce/internal/template"
)

type Template struct {
	ID           uuid.UUID              `json:"id"`
	ProjectID    uuid.UUID              `json:"project_id"`
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	Template     string                 `json:"template"`
	Placeholders []template.Placeholder `json:"placeholders"`
}
