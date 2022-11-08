package entity

import (
	"github.com/google/uuid"
)

type HTTPHeader struct {
	Name     string `json:"name"`
	Required bool   `json:"required"`
}

type Pipe struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Tags        []string  `json:"tags"`
	Description string    `json:"description"`

	// Endpoint is the URL of the pipe. This is where the request is sent to.
	Endpoint string `json:"url"`

	// Schema is the JSON schema of the parameters.
	Schema string `json:"schema"`

	// Headers is a list of headers that the pipe accepts.
	Headers []HTTPHeader `json:"headers"`
}
