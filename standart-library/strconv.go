package main

import (
	"fmt"
	"strconv"
)

func main() {
	// parsebool
	dataBool, errBool := strconv.ParseBool("true")
	if errBool != nil {
		fmt.Printf("Error Parsing Boolean: %s", errBool.Error())
		return
	}
	fmt.Printf("Boolean aman: %v\n", dataBool)

	// atoi
	dataAtoi, errAtoi := strconv.Atoi("100")
	if errAtoi != nil {
		fmt.Printf("Error Atoi: %s", errAtoi.Error())
		return
	}
	fmt.Printf("Atoi aman: %v\n", dataAtoi)

	// Format int
	dataFormatInt := strconv.FormatInt(900, 2)
	fmt.Printf("FormatInt aman: %v\n", dataFormatInt) // binary

	// itoa
	dataItoa := strconv.Itoa(900)
	fmt.Printf("Itoa aman: %v\n", dataItoa) // string
}
