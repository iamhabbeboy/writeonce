package entity

type Pipeline struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	// Pipes is a list of pipes that are executed in order.
	Pipes []PipeInput `json:"pipes"`
}
