package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/theterminalguy/writeonce/internal/entity"
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
	})
}

type ProjectRepository struct {
}

type ProjectParams struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{}
}

func (r *ProjectRepository) GetAll() []entity.Project {
	return projects
}

var ErrProjectNameNotUnique = errors.New("project name not unique")

func (r *ProjectRepository) Create(p ProjectParams) (*entity.Project, error) {
	if err := ValidateParams(p); err != nil {
		return nil, err
	}
	// validate unique name: TODO this is for demo purposes only
	for _, project := range projects {
		if project.Name == p.Name {
			return nil, ErrProjectNameNotUnique
		}
	}
	project := entity.Project{
		ID:          uuid.New(),
		Name:        p.Name,
		Description: p.Description,
	}
	projects = append(projects, project)
	return &project, nil
}

func (r *ProjectRepository) Get(id uuid.UUID) (*entity.Project, error) {
	for _, project := range projects {
		if project.ID == id {
			return &project, nil
		}
	}
	return nil, nil
}

func (r *ProjectRepository) Update(id uuid.UUID, p ProjectParams) (*entity.Project, error) {
	for i, project := range projects {
		if project.ID == id {
			projects[i].Name = p.Name
			projects[i].Description = p.Description
			return &projects[i], nil
		}
	}
	return nil, nil
}

func (r *ProjectRepository) Delete(id uuid.UUID) error {
	for i, project := range projects {
		if project.ID == id {
			projects = append(projects[:i], projects[i+1:]...)
			return nil
		}
	}
	return nil
}
