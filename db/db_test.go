package db_test

import (
	"testing"

	"t0ast.cc/symflower-live-chat/db"
	"t0ast.cc/symflower-live-chat/graph/model"
)

func TestInsertUser(t *testing.T) {
	db.ResetDatabase()
	username := "MyUser"
	user, err := db.InsertUser(model.NewUser{
		Name: username,
	})
	if err != nil {
		t.Fatalf("Err was not nil: %s", err.Error())
	}
	if user.Name != username {
		t.Fatalf("Username of inserted user does not match input:\n  Inserted: %s\n  Input: %s", user.Name, username)
	}
}

func TestInsertExistingUser(t *testing.T) {
	db.ResetDatabase()
	username := "MyUser"
	user, err := db.InsertUser(model.NewUser{
		Name: username,
	})
	if err != nil {
		t.Fatalf("Err was not nil: %s", err.Error())
	}
	if user.Name != username {
		t.Fatalf("Username of inserted user does not match input:\n  Inserted: %s\n  Input: %s", user.Name, username)
	}
	_, err = db.InsertUser(model.NewUser{
		Name: username,
	})
	if err == nil {
		t.Fatalf("Err was nil when an err was expected: %s", err.Error())
	}
	if err.Error() != "User "+username+" already exists" {
		t.Fatalf("Unexpected error message: %s", err.Error())
	}
}

func TestGetUsers(t *testing.T) {
	db.ResetDatabase()
	usernames := []string{
		"MyUser", "MyUser2",
	}
	for _, u := range usernames {
		db.InsertUser(model.NewUser{
			Name: u,
		})
	}

	users := db.GetUsers()
	if expected, actual := len(usernames), len(users); actual != expected {
		t.Fatalf("Returned user count is not %d: %d", expected, actual)
	}
	for i, expected := range usernames {
		if actual := users[i].Name; actual != expected {
			t.Fatalf("User at position %d is not named as expected:\n  Expected: %s\n  Actual: %s", i, expected, actual)
		}
	}
}
