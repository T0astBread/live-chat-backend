package db

import (
	"fmt"

	"t0ast.cc/symflower-live-chat/graph/model"
)

type database struct {
	Users    []*model.User
	Messages []*model.Message
}

var db = database{
	Users:    []*model.User{},
	Messages: []*model.Message{},
}

// ResetDatabase sets up a fresh, empty database
func ResetDatabase() {
	db = database{
		Users:    []*model.User{},
		Messages: []*model.Message{},
	}
}

// GetUsers returns all users in the database
func GetUsers() []*model.User {
	return db.Users
}

// InsertUser inserts a new user into the database or fails if a user
// with the specified name already exists
func InsertUser(newUser model.NewUser) (*model.User, error) {
	for _, existing := range db.Users {
		if existing.Name == newUser.Name {
			return nil, fmt.Errorf("User %s already exists", newUser.Name)
		}
	}
	inserted := &model.User{
		ID:   len(db.Users),
		Name: newUser.Name,
	}
	db.Users = append(db.Users, inserted)
	return inserted, nil
}
