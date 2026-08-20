package helper

import (
	"container/list"
	"fmt"
	"os"
	"path/filepath"
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

type InsertListFromArrayInterface interface {
	Numeric | Float | string
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

func InsertListFromArray[T InsertListFromArrayInterface](array []T) *list.List {
	var data *list.List = list.New()
	for _, e := range array {
		data.PushBack(e)
	}

	return data
}

func CreateNewFile(targetPath, filename, message string) error {
	filePath := filepath.Join(targetPath, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(targetPath, filename string) (string, error) {
	filePath := filepath.Join(targetPath, filename)
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func AddToFile(filename, message string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(message)
	return nil
}

func IsFolderExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func CreateFolder(name string) error {
	err := os.Mkdir(name, 0755)
	if err != nil {
		fmt.Println("Error creating folder: ", err.Error())
		return err
	}

	return nil
}

func CheckFolderExistance(name string) (bool, error) {
	isFolderExists := IsFolderExists(name)
	if !isFolderExists {
		err := CreateFolder(name)
		if err != nil {
			return false, err
		}
	}

	fmt.Printf("✓ Folder '%s' telah dibuat\n", name)
	return true, nil
}
