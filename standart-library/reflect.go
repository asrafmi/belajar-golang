package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode/utf8"
)

type Sample struct {
	Name string
}

type Person struct {
	Name   string `required:"true" max:"10"`
	Addres string `required:"false" max:"25"`
	Age    int    `required:"true" max:"100"`
}

func readField(value any) {
	fmt.Printf("===\n")
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var value reflect.StructField = valueType.Field(i)
		fmt.Printf("Field -> %s\n", value.Name)
		fmt.Printf("Type -> %s\n", value.Type)
		fmt.Printf("---\n")
	}
}

func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			isInt := reflect.ValueOf(value).Field(i).CanInt()
			if isInt {
				result = data != 0
			} else {
				result = data != ""
			}

			if !result {
				return result
			}
		}

		fType := f.Type.String()
		if fType == "string" {
			data := reflect.ValueOf(value).Field(i).Interface()
			s, _ := data.(string)
			charLength := utf8.RuneCountInString(s)
			charMax, _ := strconv.Atoi(f.Tag.Get("max"))
			result = charLength <= charMax
			if !result {
				return result
			}
		}
		if fType == "int" {
			data := reflect.ValueOf(value).Field(i).Interface()
			num, _ := data.(int)
			numMax, _ := strconv.Atoi(f.Tag.Get("max"))
			result = num <= numMax
			if !result {
				return result
			}
		}
	}

	return result
}

func main() {
	// readField(Sample{"joko"})
	// readField(Person{"joko", "solo", 12})

	person := Person{
		Name:   "Joko Widod",
		Addres: "Solo",
		Age:    100,
	}

	isValid := IsValid(person)
	fmt.Println("is person valid ?", isValid)
}
