package repository

import (
	"github.com/google/uuid"
	"github.com/theterminalguy/writeonce/internal/entity"
)

type PipeRepository struct {
}

func NewPipeRepository() *PipeRepository {
	return &PipeRepository{}
}

func (r *PipeRepository) GetAll() []entity.Pipe {
	return pipes
}

func (r *PipeRepository) Get(id uuid.UUID) (*entity.Pipe, error) {
	for _, pipe := range pipes {
		if pipe.ID == id {
			return &pipe, nil
		}
	}
	return nil, nil
}

func (r *PipeRepository) Create(pipe entity.Pipe) (*entity.Pipe, error) {
	pipes = append(pipes, pipe)
	return &pipe, nil
}
