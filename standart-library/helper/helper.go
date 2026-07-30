package helper

import (
	"fmt"
	"strconv"
	"strings"
)

const resultText = "Result: "

// Robust numeric interface
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Integer constraint for operations that require integer-only operators (like %)
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

func Add[T Numeric](a, b T) T {
	return a + b
}

func Minus[T Numeric](a, b T) T {
	return a - b
}

func Multiply[T Numeric](a, b T) T {
	return a * b
}

func Divide[T Numeric](a, b T) T {
	if b == 0 {
		fmt.Println("Error cannot divide by zero")
		return 0
	}
	return a / b
}

func Modulo[T Integer](a, b T) T {
	return a % b
}

func OperateInt[T Integer](operation *string, a *T, b *T) {
	switch *operation {
	case "add":
		fmt.Println(resultText, Add(*a, *b))
	case "minus":
		fmt.Println(resultText, Minus(*a, *b))
	case "multiply":
		fmt.Println(resultText, Multiply(*a, *b))
	case "divide":
		fmt.Println(resultText, Divide(*a, *b))
	case "modulo":
		fmt.Println(resultText, Modulo(*a, *b))
	default:
		fmt.Printf("Error: Operation %s not found\n", *operation)
	}
}

func OperateFloat[T Float](operation *string, a *T, b *T) {
	switch *operation {
	case "add":
		fmt.Println(resultText, Add(*a, *b))
	case "minus":
		fmt.Println(resultText, Minus(*a, *b))
	case "multiply":
		fmt.Println(resultText, Multiply(*a, *b))
	case "divide":
		fmt.Println(resultText, Divide(*a, *b))
	case "modulo":
		fmt.Println("Error: Modulo operation is not available for float types")
	default:
		fmt.Printf("Error: Operation %s not found\n", *operation)
	}
}

func Operate[T Numeric](operation *string, a *T, b *T) {
	switch *operation {
	case "add":
		fmt.Println(resultText, Add(*a, *b))
	case "minus":
		fmt.Println(resultText, Minus(*a, *b))
	case "multiply":
		fmt.Println(resultText, Multiply(*a, *b))
	case "divide":
		fmt.Println(resultText, Divide(*a, *b))
	case "modulo":
		fmt.Println("Error: Modulo operation is only available for integer types")
	default:
		fmt.Printf("Error: Operation %s not found\n", *operation)
	}
}

func IsFloat(number string) bool {
	_, err := strconv.ParseFloat(number, 64)
	return err == nil
}

func IsInt(number string) bool {
	_, err := strconv.Atoi(number)
	return err == nil
}

func IsStringExist(value string) bool {
	if value == "" {
		return false
	}

	return true
}

func IsEmail(email string) bool {
	if !strings.Contains(email, "@") {
		return false
	}

	return true
}

func IsEmailValid(email string) (bool, string) {
	if !strings.Contains(email, "@") {
		return false, "should contains '@'"
	}
	if !strings.Contains(email, ".") {
		return false, "should contains ."
	}
	if strings.Contains(email, " ") {
		return false, "contains space"
	}
	if strings.HasPrefix(email, ".") || strings.HasSuffix(email, ".") {
		return false, "can not contains '.' at beginning or last email"
	}

	return true, ""
}

func IsAgeValid(age int) bool {
	if age < 13 || age > 120 {
		return false
	}

	return true
}
