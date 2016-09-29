package userkit

import (
	"strings"
	"testing"

	"github.com/workpail/userkit-go/utils"
)

var uk UserKit
var user1 *User
var user1PW string
var user1Sess *SessionToken

func init() {
	// create client
	uk = NewUserKit(utils.GetTestKey())

	// create a default user1
	user1PW = utils.RandStr(14)
	data := map[string]string{
		"email":    utils.RandEmail(),
		"password": user1PW}

	var err error
	user1, err = uk.Users.Create(data)
	if err != nil {
		panic("Could not create default test user")
	}

	// create a default user1 session
	user1Sess, err = uk.Users.Login(user1.Email, user1PW, "")
	if err != nil {
		panic("Could not create default test user session")
	}
}

func TestUserCreate(t *testing.T) {
	data := map[string]string{
		"email":    utils.RandEmail(),
		"password": utils.RandStr(14)}

	user, err := uk.Users.Create(data)
	if err != nil {
		t.Errorf("API Error: %s", err)
	}

	if user.Email != strings.ToLower(data["email"]) {
		t.Errorf("Expected email: %s, but got %s", data["email"], user.Email)
	}
}

func TestGetUser(t *testing.T) {
	user, err := uk.Users.Get(user1.ID)
	if err != nil {
		t.Errorf("API Error: %s", err)
	}

	if user.ID != user1.ID {
		t.Errorf("Expected user with id: %s, got %s", user1.ID, user.ID)
	}
}

func TestLoginUser(t *testing.T) {
	session, err := uk.Users.Login(user1.Email, user1PW, "")
	if err != nil {
		t.Errorf("API Error: %s", err)
	}

	if session.Token == "" {
		t.Error("Expected a session with token value, got empty string")
	}
}

func TestGetCurrentUser(t *testing.T) {
	user, err := uk.Users.GetCurrentUser(user1Sess.Token)
	if err != nil {
		t.Errorf("API Error: %s", err)
	}

	if user.ID != user1.ID {
		t.Errorf("Expected user with ID: %s, got %s instead", user1.ID, user.ID)
	}
}
