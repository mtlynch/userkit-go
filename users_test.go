package userkit

import (
	"strings"
	"testing"

	"github.com/workpail/userkit-go/utils"
)

var isSetup = false
var uk Client
var user1 *User
var user1PW string
var user1Sess *Session

func setupOnce() {
	if isSetup {
		return
	}
	isSetup = true

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

/* User tests */

func TestUserCreate(t *testing.T) {
	setupOnce()
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

func TestUpdateUser(t *testing.T) {
	setupOnce()
	data := map[string]string{"name": "Jane Smith"}
	user, err := uk.Users.Update(user1.ID, data)
	if err != nil {
		t.Errorf("API Error: %s", err)
	}

	if user.Name != data["name"] {
		t.Errorf("Expected name: %s, but got %s", data["name"], user.Name)
	}
}

func TestGetUser(t *testing.T) {
	setupOnce()
	user, err := uk.Users.Get(user1.ID)
	if err != nil {
		t.Errorf("API Error: %s", err)
	}

	if user.ID != user1.ID {
		t.Errorf("Expected user with id: %s, got %s", user1.ID, user.ID)
	}
}

func TestLoginUser(t *testing.T) {
	setupOnce()
	session, err := uk.Users.Login(user1.Email, user1PW, "")
	if err != nil {
		t.Errorf("API Error: %s", err)
	}

	if session.Token == "" {
		t.Error("Expected a session with token value, got empty string")
	}
}

func TestGetUserBySession(t *testing.T) {
	setupOnce()
	user, err := uk.Users.GetUserBySession(user1Sess.Token)
	if err != nil {
		t.Errorf("API Error: %s", err)
	}

	if user.ID != user1.ID {
		t.Errorf("Expected user with ID: %s, got %s instead", user1.ID, user.ID)
	}
}

/* Errors tests */

func TestCreateUserBadUsernameAndEmail(t *testing.T) {
	setupOnce()
	badData := map[string]string{"username": "bad us ername",
		"email": "bademail", "password": utils.RandStr(14)}
	_, err := uk.Users.Create(badData)

	if err == nil {
		t.Error("Expected error, but got nil instead")
	}

	ukErr := err.(Error)

	// ensure it's the right type of error
	if ukErr.Type != "invalid_request_error" {
		t.Errorf("Expected 'invalid_request_error' but got %s", ukErr.Type)
	}

	// ensure it has multiple Errors
	errLength := len(ukErr.Errors)
	if errLength < 2 {
		t.Errorf("Expected at least 2 errors, got: %d", errLength)
		t.Errorf("Errors: %v", ukErr.Errors)
	}
}

func TestGetUserError(t *testing.T) {
	setupOnce()
	_, err := uk.Users.Get("BadUserID")

	if err == nil {
		t.Errorf("Expected error, but got nil instead")
	}

	ukErr := err.(Error)

	// ensure only one error returned
	if len(ukErr.Errors) != 0 {
		t.Errorf("Expected error.Errors to be empty list, but got: %+v", ukErr.Errors)
	}
}
