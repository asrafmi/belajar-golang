package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`J([a-z]*)i`)

	fmt.Println(regex.MatchString("Jokowi"))
	fmt.Println(regex.MatchString("Jok0wi"))
	fmt.Println(regex.MatchString("Jokow1"))

	fmt.Println(regex.FindAllString("Joi JOi Joy joi jgo jqo jco Jyo", 10))
}
