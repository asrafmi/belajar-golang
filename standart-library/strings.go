package main

import (
	"flag"
	"fmt"
	"standart-library/helper"
	"strings"
)

/*
=== SOAL 1: Email Validator ===
Buat program validasi email dengan ketentuan:
- Email harus mengandung "@"
- Email harus mengandung "."
- Email tidak boleh dimulai atau diakhiri dengan "."
- Email tidak boleh ada spasi
- Tampilkan pesan valid/invalid

Input: "user@example.com", "invalid.email@", "user @email.com", "user@email"

Output contoh:
$ go run email_validator.go -email="user@example.com"
✓ Email valid: user@example.com

$ go run email_validator.go -email="user @email.com"
✗ Email invalid: user @email.com (contains space)
*/
func number1string() {
	email := flag.String("email", "", "email")
	flag.Parse()

	if !helper.IsStringExist(*email) {
		fmt.Printf("✗ Email required\n")
		return
	}
	valid, msg := helper.IsEmailValid(*email)
	if !valid {
		fmt.Printf("✗ Email invalid: user %s (%s)\n", *email, msg)
		return
	}

	fmt.Printf("✓ Email valid: %s\n", *email)
}

/*
=== SOAL 2: Password Strength Checker ===
Buat program cek kekuatan password:
- Minimal 8 karakter
- Harus ada huruf besar (A-Z)
- Harus ada huruf kecil (a-z)
- Harus ada angka (0-9)
- Harus ada special character (!@#$%^&*)
- Tampilkan detail requirement apa yang kurang

Input: "Pass123!", "password", "PASSWORD123!", "Pass123"

Output contoh:
$ go run password_checker.go -password="Pass123!"
✓ Password strong!
- Length: 8+ ✓
- Uppercase: ✓
- Lowercase: ✓
- Number: ✓
- Special char: ✓

$ go run password_checker.go -password="password"
✗ Password weak!
- Length: 8+ ✓
- Uppercase: ✗ (missing uppercase letter)
- Lowercase: ✓
- Number: ✗ (missing number)
- Special char: ✗ (missing special character)
*/

func hasUppercase(password string) bool {
	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func hasLowercase(password string) bool {
	for _, char := range password {
		if char >= 'a' && char <= 'z' {
			return true
		}
	}
	return false
}

func hasNumber(password string) bool {
	for _, char := range password {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func hasSpecialChar(password string) bool {
	specialChars := "!@#$%^&*"
	return strings.ContainsAny(password, specialChars)
}

func checkPasswordStrength(password string) {
	if !helper.IsStringExist(password) {
		fmt.Printf("✗ Password required\n")
		return
	}

	isLengthValid := len(password) >= 8
	hasUpper := hasUppercase(password)
	hasLower := hasLowercase(password)
	hasNum := hasNumber(password)
	hasSpecial := hasSpecialChar(password)

	isStrong := isLengthValid && hasUpper && hasLower && hasNum && hasSpecial

	if isStrong {
		fmt.Printf("✓ Password strong!\n")
	} else {
		fmt.Printf("✗ Password weak!\n")
	}

	fmt.Printf("- Length: 8+ %s\n", boolToCheck(isLengthValid))
	fmt.Printf("- Uppercase: %s", boolToCheck(hasUpper))
	if !hasUpper {
		fmt.Print(" (missing uppercase letter)")
	}
	fmt.Printf("\n")

	fmt.Printf("- Lowercase: %s", boolToCheck(hasLower))
	if !hasLower {
		fmt.Print(" (missing lowercase letter)")
	}
	fmt.Printf("\n")

	fmt.Printf("- Number: %s", boolToCheck(hasNum))
	if !hasNum {
		fmt.Print(" (missing number)")
	}
	fmt.Printf("\n")

	fmt.Printf("- Special char: %s", boolToCheck(hasSpecial))
	if !hasSpecial {
		fmt.Print(" (missing special character)")
	}
	fmt.Printf("\n")
}

func boolToCheck(value bool) string {
	if value {
		return "✓"
	}
	return "✗"
}

func number2string() {
	password := flag.String("password", "", "password")
	flag.Parse()
	checkPasswordStrength(*password)
}

/*
=== SOAL 3: Text Filter & Formatter ===
Buat program yang bisa:
1. Capitalize first letter setiap kata
2. Hapus extra spaces (lebih dari 1 spasi jadi 1)
3. Hapus special characters (hanya keep alphanumeric + space)
4. Count jumlah kata, huruf, dan spasi
5. Reverse text

Flag:
- -text (required, teks yang akan diproses)
- -capitalize (bool, capitalize setiap kata)
- -remove-spaces (bool, hapus extra spaces)
- -remove-special (bool, hapus special characters)
- -count (bool, hitung kata/huruf/spasi)
- -reverse (bool, reverse text)

Output contoh:
$ go run text_formatter.go -text="hello  world!!!" -capitalize -remove-spaces -remove-special -count
✓ Original: hello  world!!!
✓ Formatted: Hello World
✓ Word count: 2
✓ Character count: 10
✓ Space count: 1

$ go run text_formatter.go -text="programming is fun" -reverse
✓ Original: programming is fun
✓ Reversed: nuf si gnimmargorp

Tips untuk ngerjain:
- Gunakan strings.Contains(), strings.Index(), strings.Split()
- Gunakan strings.ToUpper(), strings.ToLower(), strings.Title()
- Gunakan strings.ReplaceAll(), strings.Trim(), strings.TrimSpace()
- Gunakan strings.FieldsFunc() untuk custom split
- Untuk reverse, convert ke []rune, reverse loop, convert balik ke string
- Untuk capitalize per kata, split dulu terus process satu per satu
*/

func capitalizeWords(text string) string {
	words := strings.Fields(text)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

func removeExtraSpaces(text string) string {
	text = strings.TrimSpace(text)
	for strings.Contains(text, "  ") {
		text = strings.ReplaceAll(text, "  ", " ")
	}
	return text
}

func removeSpecialChars(text string) string {
	result := ""
	for _, char := range text {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') || char == ' ' {
			result += string(char)
		}
	}
	return result
}

func reverseText(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func countTextStats(text string) (int, int, int) {
	words := strings.Fields(text)
	wordCount := len(words)
	charCount := len(strings.ReplaceAll(text, " ", ""))
	spaceCount := strings.Count(text, " ")
	return wordCount, charCount, spaceCount
}

func number3string() {
	text := flag.String("text", "", "text to process")
	capitalize := flag.Bool("capitalize", false, "capitalize words")
	removeSpaces := flag.Bool("remove-spaces", false, "remove extra spaces")
	removeSpecial := flag.Bool("remove-special", false, "remove special characters")
	count := flag.Bool("count", false, "count words/chars/spaces")
	reverse := flag.Bool("reverse", false, "reverse text")
	flag.Parse()

	if !helper.IsStringExist(*text) {
		fmt.Printf("✗ Text required\n")
		return
	}

	fmt.Printf("✓ Original: %s\n", *text)

	result := *text

	if *removeSpecial {
		result = removeSpecialChars(result)
	}
	if *removeSpaces {
		result = removeExtraSpaces(result)
	}
	if *capitalize {
		result = capitalizeWords(result)
	}
	if *reverse {
		result = reverseText(result)
	}

	if *removeSpecial || *removeSpaces || *capitalize || *reverse {
		fmt.Printf("✓ Formatted: %s\n", result)
	}

	if *count {
		wordCount, charCount, spaceCount := countTextStats(result)
		fmt.Printf("✓ Word count: %d\n", wordCount)
		fmt.Printf("✓ Character count: %d\n", charCount)
		fmt.Printf("✓ Space count: %d\n", spaceCount)
	}
}

func cobacoba() {
	fmt.Println(strings.Contains("Joko Widodo", "dodo"))
	fmt.Println(strings.Split("Joko Widodo", " "))
	fmt.Println(strings.ToLower("Joko Widodo"))
	fmt.Println(strings.ToUpper("Joko Widodo"))
	fmt.Println(strings.Trim("  Joko Widodo do  ", " "))
	fmt.Println(strings.ReplaceAll("avukuvu suvukavu mavukanvu navusivugovurengvu", "vu", ""))
}
func main() {
	// number1string()
	// number2string()
	number3string()
}
