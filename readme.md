<h1 align="center">enforcer</h1>
<img src="https://img.shields.io/badge/go-%2300ADD8.svg?&style=for-the-badge&logo=go&logoColor=white"  align="right" position="absolute">


## Instalation

```
go get https://github.com/vitorhugoze/enforcer
```

## Description

**Enforcer** is a library used to annotate structs with rules that must be followed, otherwise, when validating an instance of that struct a panic will occur.

### Project Structure

```bash
├───internal
│   └───rules
└───pkg
    └───enforcer
```

## Usage

#### Create a struct with annotations defining the rules:

```go
type User struct {
	Name        string `rules:"required, min=5, max=30"`
	Email       string `rules:"required, email"`
	Credentials UserCredentials
}

type UserCredentials struct {
	StudentNumer string `rules:"min=10, max=30"`
	PassWord     string `rules:"password"`
}
```

#### Use an enforcer to verify the instance:
```go
	//New enforcer used to validate the rules
	enforcer := enforcer.GetEnforcer()
	//Creates an example user
	usr := User{
		Name:  "testing",
		Email: "test@mail.com",
		Credentials: UserCredentials{
			StudentNumer: "0123456789",
			PassWord:     "aA@1234567",
		},
	}
	//Validates and panics if rules were not followed
	enforcer.ValidateRules(usr)
```