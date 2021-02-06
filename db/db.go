package db

import (
	"fmt"
	"math/rand"

	"t0ast.cc/symflower-live-chat/graph/model"
)

type database struct {
	Users    []*model.User
	Messages []*model.Message
}

var db database
var messageSubscriptions map[int]chan *model.Message

// Reset sets up a fresh, empty database
func Reset() {
	db = database{
		Users:    []*model.User{},
		Messages: []*model.Message{},
	}
	messageSubscriptions = make(map[int]chan *model.Message)
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

// GetMessages returns all messages in the database
func GetMessages() []*model.Message {
	return db.Messages
}

// InsertMessage inserts a new message into the database or fails if
// the specified poster doesn't exist
func InsertMessage(newMessage model.NewMessage) (*model.Message, error) {
	var poster *model.User
	for _, u := range db.Users {
		if u.ID == newMessage.PosterID {
			poster = u
			break
		}
	}
	if poster == nil {
		return nil, fmt.Errorf("User with ID %d does not exist", newMessage.PosterID)
	}

	inserted := &model.Message{
		ID:      len(db.Messages),
		Content: newMessage.Content,
		Poster:  poster,
	}
	db.Messages = append(db.Messages, inserted)

	for _, s := range messageSubscriptions {
		s <- inserted
	}

	return inserted, nil
}

// RegisterMessageSubscription registers the provided message channel
// as a subscriber for new messages and returns a subscription handle
// used for unregistering
func RegisterMessageSubscription(sub chan *model.Message) int {
	var subHandle int
	for {
		subHandle = rand.Int()
		if _, handleIsTaken := messageSubscriptions[subHandle]; !handleIsTaken {
			break
		}
	}
	messageSubscriptions[subHandle] = sub
	return subHandle
}

// UnregisterMessageSubscription stops a subscription from recieving
// further message updates
func UnregisterMessageSubscription(subHandle int) {
	delete(messageSubscriptions, subHandle)
}
