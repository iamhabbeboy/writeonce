package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/theterminalguy/writeonce/internal/entity"
)

var projects []entity.Project

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
