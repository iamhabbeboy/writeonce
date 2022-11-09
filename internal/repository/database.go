package repository

import (
	"github.com/google/uuid"
	"github.com/theterminalguy/writeonce/internal/entity"
	"github.com/theterminalguy/writeonce/internal/template"
)

// this is fake data for demo purposes only

var projects []entity.Project
var templates []entity.Template
var pipes []entity.Pipe

func init() {
	projects = append(projects, entity.Project{
		ID:          uuid.MustParse("cf4e5443-7f24-4014-8a52-64a53201d1c2"),
		Name:        "Project 1",
		Description: "Description 1",
	})
	templates = append(templates, entity.Template{
		ID:          uuid.MustParse("928afa3e-e0aa-4eb0-86dd-e70fa4ff961f"),
		Name:        "Template 1",
		Description: "Description 1",
		Template:    "Hello {{.Name}}",
		ProjectID:   uuid.MustParse("cf4e5443-7f24-4014-8a52-64a53201d1c2"),
		Placeholders: []template.Placeholder{
			{
				Name: "Name",
				Pos:  8,
			},
		},
	})
	pipes = []entity.Pipe{
		{
			ID:          uuid.MustParse("3e42bd93-a627-4a13-8a84-e8e80b6b3763"),
			Name:        "Pipe 1",
			Description: "Description 1",
			Endpoint:    "http://localhost:4567/generate-pdf",
			Schema: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name": map[string]interface{}{
						"type": "string",
					}},
				"required": []string{
					"name",
				},
			},
			Headers: []entity.HTTPHeader{
				{
					Name:     "Authorization",
					Required: true,
				},
			},
		},
		{
			ID:          uuid.MustParse("4e42bd93-a627-4a13-8a84-e8e80b6b3763"),
			Name:        "Pipe 2",
			Description: "Description 2",
			Endpoint:    "https://jsonplaceholder.typicode.com/todos",
		},
	}
}
