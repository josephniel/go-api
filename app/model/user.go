package model

// User is the map for `users` DB
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name" validate:"required"`
}
