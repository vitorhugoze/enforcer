package rules

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

var FaultHandler func(string)

func RequiredRule(fieldName, fieldVal string, params ...string) {
	if len(fieldVal) == 0 {
		handleFault(fmt.Sprintf("field '%v' is empty, breaking the 'required' rule", fieldName))
	}
}

func MinRule(fieldName, fieldVal string, params ...string) {
	min := must(strconv.Atoi(params[1]))
	if len(fieldVal) < min {
		handleFault(fmt.Sprintf("lenght of '%v' is less than 'min' value(%v)", fieldName, min))
	}
}

func MaxRule(fieldName, fieldVal string, params ...string) {
	max := must(strconv.Atoi(params[1]))
	if len(fieldVal) > max {
		handleFault(fmt.Sprintf("lenght of '%v' is bigger than 'max' value(%v)", fieldName, max))
	}
}

func EmailRule(fieldName, fieldVal string, params ...string) {
	if !must(regexp.Match("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$", []byte(fieldVal))) {
		handleFault(fmt.Sprintf("'%v' is not an valid e-mail", fieldVal))
	}
}

func PasswordRule(fieldName, fieldVal string, params ...string) {
	number := false
	symbol := false
	lowerCase := false
	upperCase := false

	for _, c := range fieldVal {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsSymbol(c) || unicode.IsPunct(c):
			symbol = true
		case unicode.IsLower(c):
			lowerCase = true
		case unicode.IsUpper(c):
			upperCase = true
		}
	}

	if !number || !symbol || !lowerCase || !upperCase {
		handleFault(fmt.Sprintf(`'%v' is not a valid password, a password shoud have:
									At least one lower case letter
									At least one upper case letter
									At least one number
									At least one symbol`, fieldVal))
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		handleFault(err.Error())
	}

	return v
}

func handleFault(reason string) {
	if FaultHandler != nil {
		FaultHandler(reason)
	} else {
		panic(reason)
	}
}
