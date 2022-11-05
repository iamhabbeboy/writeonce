package repository

import "github.com/theterminalguy/writeonce/internal/entity"

type ProjectRepository struct {
}

type ProjectParams struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{}
}

func (r *ProjectRepository) Create(p ProjectParams) (*entity.Project, error) {
	if err := ValidateParams(p); err != nil {
		return nil, err
	}
	return nil, nil
}
