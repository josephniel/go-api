package schema

// Sample is a schema used by the sample controller and its operations
type Sample struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
