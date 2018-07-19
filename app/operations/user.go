package operations

import "github.com/josephniel/go-api/app/models"

// GetUser retrieves a user from the users table using its id
func GetUser(id int64) (*models.User, error) {
	user := &models.User{
		ID: id,
	}
	err := models.DB.Select(user)
	return user, err
}

// AddUser adds a user to the users table
func AddUser(user *models.User) error {
	return models.DB.Insert(user)
}
