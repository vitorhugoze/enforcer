package enforcer

import (
	"testing"
)

type User struct {
	Name        string `rules:"required, min=5, max=30"`
	Email       string `rules:"required, email"`
	Credentials UserCredentials
}

type UserCredentials struct {
	StudentNumer string `rules:"min=10, max=30"`
	PassWord     string `rules:"password"`
}

func TestValidUser(t *testing.T) {
	usr := getUserForTesting()
	ruleEnforcer.ValidateRules(usr)
}

func TestRecursiveChecking(t *testing.T) {
	usr := getUserForTesting()
	usr.Credentials.PassWord = "012345"

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("should have panicked, password lenght is only 5")
		}
	}()
	ruleEnforcer.ValidateRules(usr)
}

func TestInvalidEmail(t *testing.T) {
	usr := getUserForTesting()
	usr.Email = "test@mail."

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("should have panicked, e-mail is invalid")
		}
	}()
	ruleEnforcer.ValidateRules(usr)
}

func TestInvalidPassword(t *testing.T) {
	usr := getUserForTesting()
	usr.Credentials.PassWord = "aA@"

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("should have panicked, password does not contains numbers")
		}
	}()
	ruleEnforcer.ValidateRules(usr)
}

func getUserForTesting() User {
	return User{
		Name:  "testing",
		Email: "test@mail.com",
		Credentials: UserCredentials{
			StudentNumer: "0123456789",
			PassWord:     "aA@1234567",
		},
	}
}
