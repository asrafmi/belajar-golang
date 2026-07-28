package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func getById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "asraf" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := getById("asraf")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error cuy")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error cuy")
		} else {
			fmt.Println("unknown error cuy")
		}
	}
}
