package repository

import (
	"errors"

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
	return nil, errors.New("pipe not found")
}

func (r *PipeRepository) Create(pipe entity.Pipe) (*entity.Pipe, error) {
	pipe.ID = uuid.New()
	pipes = append(pipes, pipe)
	return &pipe, nil
}

func (r *PipeRepository) Update(id uuid.UUID, pipe entity.Pipe) (*entity.Pipe, error) {
	record, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	if record == nil {
		return nil, nil
	}
	record = &pipe
	return record, nil
}
