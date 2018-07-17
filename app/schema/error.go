package schema

// Error is a generic structure for failed HTTP responses
type Error struct {
	Message string `json:"message"`
}
