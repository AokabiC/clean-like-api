package domain

import (
	"errors"
	"regexp"
)

type Username string

func NewUsername(s string) (Username, error) {
	var IsValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]{1,16}$`).MatchString
	if !IsValidUsername(s) {
		return "", errors.New("username must consist of alphanumeric characters and underscore, and must be less than 16 characters long")
	}

	return Username(s), nil
}
