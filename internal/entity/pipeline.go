package entity

import (
	"github.com/google/uuid"
)

type PipeInput struct {
	PipeID uuid.UUID `json:"id"`
	// Parameters is the JSON object that is sent to the pipe.
	Parameters interface{} `json:"parameters"`

	// Headers is a list of headers that are sent to the pipe.
	Headers map[string]string `json:"headers"`
}

type PipeStatus struct {
	ID         uuid.UUID `json:"id"`
	Err        error     `json:"error"`
	HTTPStatus int       `json:"http_status"`
}

func (p *PipeStatus) Unwrap() error {
	return p.Err
}
func (p *PipeStatus) Error() string {
	return p.Err.Error()
}

type Pipeline struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	Template struct {
		// TemplateID is the ID of the template that this pipeline is based on.
		ID uuid.UUID `json:"id"`

		// Variables is a list of variables that are used in the pipeline.
		Variables map[string]string `json:"variables"`
	} `json:"template"`

	// Pipes is a list of pipes that are executed in order.
	Pipes []PipeInput `json:"pipes"`
}
