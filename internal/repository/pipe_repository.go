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
	return nil, entity.ErrPipeNotFound
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
	record.Name = pipe.Name
	record.Description = pipe.Description
	record.Endpoint = pipe.Endpoint
	record.Schema = pipe.Schema
	record.Headers = pipe.Headers
	return record, nil
}
