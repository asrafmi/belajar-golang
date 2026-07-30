package main

import (
	"flag"
	"fmt"
	"standart-library/helper"
)

/*
=== SOAL 1: Command-line Calculator ===
Buat program yang menerima operasi matematika via flag:
- Flag `-operation` untuk jenis operasi (add, subtract, multiply, divide)
- Flag `-a` untuk angka pertama (default: 0)
- Flag `-b` untuk angka kedua (default: 0)
- Flag `-verbose` untuk menampilkan detail kalkulasi

Output contoh:
$ go run calc.go -operation=add -a=10 -b=5
Result: 15

$ go run calc.go -operation=multiply -a=4 -b=3 -verbose
Operation: multiply
Operand A: 4
Operand B: 3
Result: 12
*/

func number1flag() {
	operation := flag.String("operation", "add", "math operation")
	a := flag.Int("a", 0, "operand A")
	b := flag.Int("b", 0, "operand B")
	isVerbose := flag.Bool("verbose", false, "verbose")
	flag.Parse()

	if *isVerbose {
		fmt.Printf("Operation: %s\nOperand A: %d\nOperand B: %d\n", *operation, *a, *b)
	}

	helper.Operate(operation, a, b)
}

/*
=== SOAL 2: File Configuration Reader ===
Buat program yang membaca konfigurasi server dari flag:
- Flag `-app-name` (default: "MyApp")
- Flag `-port` (default: 8080)
- Flag `-env` untuk environment (development/production, default: "development")
- Flag `-debug` untuk mode debug (boolean)
- Flag `-timeout` dalam detik (default: 30)

Output contoh:
$ go run config.go -app-name=WebServer -port=3000 -env=production -debug -timeout=60
App Name: WebServer
Port: 3000
Environment: production
Debug Mode: true
Timeout: 60 seconds
*/
func number2flag() {
	app := flag.String("app-name", "MyApp", "app")
	port := flag.Int("port", 8080, "port")
	env := flag.String("env", "development", "env")
	debug := flag.Bool("debug", false, "debug")
	timeout := flag.Int("timeout", 30, "timeout")
	flag.Parse()

	fmt.Printf("App Name: %s\nPort: %d\nEnvironment: %s\nDebug Mode: %v\nTimeout: %d seconds\n", *app, *port, *env, *debug, *timeout)
}

/*
=== SOAL 3: User Registration CLI ===
Buat program registrasi user dengan validasi:
- Flag `-name` (required, string)
- Flag `-email` (required, string)
- Flag `-age` (default: 18)
- Flag `-country` (default: "Indonesia")
- Flag `-subscribe` untuk newsletter (boolean)

Requirements:
- Validasi email mengandung `@`
- Validasi age antara 13-120
- Tampilkan pesan error jika ada yang tidak valid
- Tampilkan summary registrasi jika valid

Output contoh:
$ go run register.go -name="Asraf Muhammad" -email="asraf@example.com" -age=25 -subscribe
✓ Registration successful!
Name: Asraf Muhammad
Email: asraf@example.com
Age: 25
Country: Indonesia
Newsletter: Subscribed

$ go run register.go -name="John" -age=10
✗ Error: Age must be between 13-120

Tips untuk ngerjain:
- Gunakan `flag.String()`, `flag.Int()`, `flag.Bool()` sesuai tipe data
- Jangan lupa `flag.Parse()` sebelum akses nilai
- Gunakan `flag.NFlag()` atau cek string kosong untuk validasi required fields
*/
func number3flag() {
	name := flag.String("name", "", "name")
	email := flag.String("email", "", "email")
	age := flag.Int("age", 18, "age")
	country := flag.String("country", "Indonesia", "country")
	subscribe := flag.Bool("subscribe", false, "subscribe")
	flag.Parse()

	if !helper.IsStringExist(*name) {
		fmt.Println("Error: Name is required!")
		return
	}
	if !helper.IsStringExist(*email) {
		fmt.Println("Error: Email is required!")
		return
	}
	if !helper.IsAgeValid(*age) {
		fmt.Println("Error: Age must be between 13-120")
		return
	}

	fmt.Printf("✓ Registration successful!\nName: %s\nEmail: %s\nAge: %d\nCountry: %s\n", *name, *email, *age, *country)
	if *subscribe {
		fmt.Println("Newsletter: Subscribed\n")
	}
}
func dbConfig() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "password", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *string = flag.String("port", "5432", "database port")
	flag.Parse()

	fmt.Printf("Username: %s\nPassword: %s \nHost: %s\nPort: %s\n", *username, *password, *host, *port)
}

func main() {
	// dbConfig()
	// number1flag()
	// number2flag()
	number3flag()
}
