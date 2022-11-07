package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/theterminalguy/writeonce/internal/entity"
	"github.com/theterminalguy/writeonce/internal/template"
)

type TemplateRepository struct {
	ProjectRepo *ProjectRepository
}

type TemplateParams struct {
	ProjectID    uuid.UUID `json:"project_id" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	Template     string    `json:"template" validate:"required"`
	Placeholders []template.Placeholder
}

func NewTemplateRepository() *TemplateRepository {
	return &TemplateRepository{
		ProjectRepo: NewProjectRepository(),
	}
}

func (r *TemplateRepository) GetAll() []entity.Template {
	return templates
}

var ErrTemplateNameNotUnique = errors.New("template name not unique")

func (r *TemplateRepository) Create(p TemplateParams) (*entity.Template, error) {
	if err := ValidateParams(p); err != nil {
		return nil, err
	}
	if _, err := r.ProjectRepo.Get(p.ProjectID); err != nil {
		return nil, err
	}
	// validate unique name: TODO this is for demo purposes only
	for _, template := range templates {
		if template.Name == p.Name {
			return nil, ErrTemplateNameNotUnique
		}
	}
	template := entity.Template{
		ID:           uuid.New(),
		Name:         p.Name,
		Description:  p.Description,
		Template:     p.Template,
		ProjectID:    p.ProjectID,
		Placeholders: p.Placeholders,
	}
	templates = append(templates, template)
	return &template, nil
}

func (r *TemplateRepository) Get(id uuid.UUID) (*entity.Template, error) {
	for _, template := range templates {
		if template.ID == id {
			return &template, nil
		}
	}
	return nil, nil
}

func (r *TemplateRepository) Update(id uuid.UUID, p TemplateParams) (*entity.Template, error) {
	for i, template := range templates {
		if template.ID == id {
			templates[i].Name = p.Name
			templates[i].Description = p.Description
			return &templates[i], nil
		}
	}
	return nil, nil
}

func (r *TemplateRepository) Delete(id uuid.UUID) error {
	for i, template := range templates {
		if template.ID == id {
			templates = append(templates[:i], templates[i+1:]...)
			return nil
		}
	}
	return nil
}
