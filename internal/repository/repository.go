package repository

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/theterminalguy/writeonce/internal/entity"
	"github.com/theterminalguy/writeonce/internal/template"
)

var projects []entity.Project
var templates []entity.Template

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
}

func ValidateParams(s interface{}, fields ...string) error {
	validate := validator.New()
	if len(fields) > 0 {
		if err := validate.StructPartial(s, fields...); err != nil {
			return err
		}
	} else {
		if err := validate.Struct(s); err != nil {
			return err
		}
	}
	return nil
}
