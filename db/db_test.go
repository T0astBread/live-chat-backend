package db_test

import (
	"fmt"
	"testing"

	"t0ast.cc/symflower-live-chat/db"
	"t0ast.cc/symflower-live-chat/graph/model"
)

func TestInsertUser(t *testing.T) {
	db.Reset()
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
	db.Reset()
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
	db.Reset()
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

func TestInsertMessage(t *testing.T) {
	db.Reset()
	msg := "This is my message"
	user, err := db.InsertUser(model.NewUser{
		Name: "dummy",
	})
	if err != nil {
		t.Fatal(fmt.Errorf("Error inserting user: %w", err))
	}
	message, err := db.InsertMessage(model.NewMessage{
		PosterID: user.ID,
		Content:  msg,
	})
	if err != nil {
		t.Fatal(fmt.Errorf("Error inserting message: %w", err))
	}
	if message.Content != msg {
		t.Fatalf("Unexpected message content:\n  Expected: %s\n  Actual: %s\n", msg, message.Content)
	}
}

func TestInsertMessageNonExistantPoster(t *testing.T) {
	db.Reset()
	posterId := 1
	_, err := db.InsertMessage(model.NewMessage{
		PosterID: posterId,
		Content:  "This is my message",
	})
	if err == nil {
		t.Fatal("No error when supplying non-existant poster for new message")
	}
	if err.Error() != fmt.Sprintf("User with ID %d does not exist", posterId) {
		t.Fatalf("Unexpected error message: %s", err.Error())
	}
}

func TestGetMessages(t *testing.T) {
	db.Reset()
	user, err := db.InsertUser(model.NewUser{
		Name: "dummy",
	})
	if err != nil {
		t.Fatal(fmt.Errorf("Error inserting user: %w", err))
	}

	messageContents := [10]string{}

	for i := 0; i < 10; i++ {
		messageContents[i] = fmt.Sprintf("Message #%d", i)

		_, err := db.InsertMessage(model.NewMessage{
			PosterID: user.ID,
			Content:  messageContents[i],
		})
		if err != nil {
			t.Fatal(fmt.Errorf("Error inserting message #%d: %w", i, err))
		}

		msgs := db.GetMessages()
		if expected, actual := i+1, len(msgs); expected != actual {
			t.Fatalf("Expected message count to be %d but got %d", expected, actual)
		}
		for j, msg := range msgs {
			if expected, actual := user.ID, msg.Poster.ID; expected != actual {
				t.Fatalf("Message #%d (in check #%d) has unexpected poster ID:\n  Expected: %d\n  Actual: %d", j, i, expected, actual)
			}
			if expected, actual := messageContents[j], msg.Content; expected != actual {
				t.Fatalf("Message #%d (in check #%d) has unexpected content:\n  Expected: %s\n  Actual: %s", j, i, expected, actual)
			}
		}
	}
}

func TestMessageSubscription(t *testing.T) {
	// Set up database
	db.Reset()
	user, err := db.InsertUser(model.NewUser{
		Name: "dummy",
	})
	if err != nil {
		t.Fatal(fmt.Errorf("Error inserting user: %w", err))
	}

	// Register subscription
	sub := make(chan *model.Message, 1)
	subHandle := db.RegisterMessageSubscription(sub)

	// Post some messages and ensure that the subscription channel
	// gets them
	for i := 0; i < 10; i++ {
		content := fmt.Sprintf("Message #%d", i)
		_, err := db.InsertMessage(model.NewMessage{
			PosterID: user.ID,
			Content:  content,
		})
		if err != nil {
			t.Fatal(fmt.Errorf("Error inserting message #%d: %w", i, err))
		}
		select {
		case subMsg := <-sub:
			if subMsg.Content != content {
				t.Fatalf("Unexpected message content in recieved message\n  Expected: %s\n  Actual: %s", content, subMsg.Content)
			}
		}
	}

	// Unregister the subscription
	db.UnregisterMessageSubscription(subHandle)

	// Post some messages and ensure that the subscription channel
	// doesn't get them
	for i := 0; i < 10; i++ {
		_, err := db.InsertMessage(model.NewMessage{
			PosterID: user.ID,
			Content:  fmt.Sprintf("Un-subscribed message #%d", i),
		})
		if err != nil {
			t.Fatal(fmt.Errorf("Error inserting message #%d: %w", i, err))
		}
		select {
		case subMsg := <-sub:
			t.Fatalf("Message in unregistered subscription channel: %s", subMsg.Content)
		default:
		}
	}
}
