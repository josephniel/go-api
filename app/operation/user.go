package operation

import "github.com/josephniel/go-api/app/model"

// GetUser retrieves a user from the users table using its id
func GetUser(id int64) (*model.User, error) {
	user := &model.User{
		ID: id,
	}
	err := model.DB.Select(user)
	return user, err
}

// AddUser adds a user to the users table
func AddUser(user *model.User) error {
	return model.DB.Insert(user)
}
