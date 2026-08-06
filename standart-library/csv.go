package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func CsvReader() {
	csvString := "joko,widodo,subagyo\n" +
		"prabowo,subianto,mbg\n" +
		"gibran,rakabuming,raka\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}

func CsvWriter() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Joko", "Widodo", "Subagyo"})
	_ = writer.Write([]string{"Prabowo", "Subianto", "Mbg"})
	_ = writer.Write([]string{"Gibran", "Rakabuming", "Raka"})

	writer.Flush()
}

/*
=== SOAL 1: Validasi Email dari CSV ===
Deskripsi:
1. Baca data CSV berisi nama, email, dan telepon
2. Gunakan regexp untuk validasi format email (format: username@domain.ext)
3. Filter hanya baris yang memiliki email valid
4. Tampilkan nama dan email yang valid

Input CSV:
john,john@example.com,08123456789
alice,alice.invalid.email,08987654321
bob,bob@gmail.com,08111222333
charlie,charlie@test,08555666777
diana,diana@yahoo.co.id,08999888777

Output yang diharapkan:
✓ john - john@example.com
✓ bob - bob@gmail.com
✓ diana - diana@yahoo.co.id

Tips untuk ngerjain:
- Gunakan regexp.MustCompile untuk pattern email
- Pattern email: `^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
- Gunakan MatchString() untuk validasi
- Looping setiap record dari CSV
*/
func number1csv() {
	// TODO: Tulis kode kamu di sini
	csvString := "john,john@example.com,08123456789\n" +
		"alice,alice.invalid.email,08987654321\n" +
		"bob,bob@gmail.com,08111222333\n" +
		"charlie,charlie@test,08555666777\n" +
		"diana,diana@yahoo.co.id,08999888777"

	reader := csv.NewReader(strings.NewReader(csvString))
	emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		isEmailValid := emailPattern.MatchString(record[1])
		if isEmailValid {
			fmt.Printf("✓ %s - %s\n", record[0], record[1])
		}
	}
}

/*
=== SOAL 2: Ekstrak dan Transformasi Nama dari CSV ===
Deskripsi:
1. Baca CSV dengan format: fullname, username, status
2. Gunakan regexp untuk ekstrak nama depan dan belakang dari fullname
3. Hanya ambil data yang username mengandung angka
4. Tampilkan: nama depan, nama belakang, username

Input CSV:
John Doe,john.doe123,active
Alice Smith,alice_smith,inactive
Bob Johnson,bob456,active
Charlie Brown,charlie.brown789,active
Diana Prince,diana_prince,active

Output yang diharapkan:
✓ John | Doe | john.doe123
✓ Bob | Johnson | bob456
✓ Charlie | Brown | charlie.brown789

Tips untuk ngerjain:
- Pattern untuk split nama: `^([A-Za-z]+)\s+([A-Za-z]+)$`
- Gunakan FindStringSubmatch() untuk ekstrak group
- Pattern untuk cek angka: `\d`
- MatchString() untuk check apakah ada angka di username
*/
func number2csv() {
	// TODO: Tulis kode kamu di sini
	csvString :=
		"John Doe,john.doe123,active\n" +
			"Alice Smith,alice_smith,inactive\n" +
			"Bob Johnson,bob456,active\n" +
			"Charlie Brown,charlie.brown789,active\n" +
			"Diana Prince,diana_prince,active\n"

	reader := csv.NewReader(strings.NewReader(csvString))
	splitNamePattern := regexp.MustCompile(`^([A-Za-z]+)\s+([A-Za-z]+)$`)
	includeNumberPattern := regexp.MustCompile(`\d`)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		splitted := splitNamePattern.FindStringSubmatch(record[0])
		firstName, lastName := splitted[1], splitted[2]
		email := record[1]

		isIncludeNumber := includeNumberPattern.MatchString(record[1])
		if !isIncludeNumber {
			continue
		}

		fmt.Printf("✓ %s | %s | %s\n", firstName, lastName, email)
	}
}

/*
=== SOAL 3: Filter dan Validasi Data Produk ===
Deskripsi:
1. Baca CSV berisi: product_name, sku, price, category
2. Gunakan regexp untuk validasi:
  - SKU harus format: 3 huruf + 4 angka (contoh: ABC1234)
  - Price hanya angka dan optional decimal (contoh: 1000 atau 99.99)
  - Category hanya huruf dan underscore

3. Tampilkan produk yang VALID dengan format: nama - SKU - price - category
4. Hitung dan tampilkan total produk valid

Input CSV:
Laptop,LAP1234,15000000,elektronik
Mouse,MSE123,250000,elektronik_aksesoris
Keyboard,KBD5678,500000,elektronik aksesoris
Monitor,MON9999,3500000.50,display
Headset,HED1234,750000,audio_device

Output yang diharapkan:
✓ Laptop - LAP1234 - 15000000 - elektronik
✓ Monitor - MON9999 - 3500000.50 - display
✓ Headset - HED1234 - 750000 - audio_device
---
✓ Total produk valid: 3

Bonus Challenge:
- Tambahkan fitur untuk menampilkan error detail (SKU invalid / Price invalid / Category invalid)
- Hitung total harga semua produk yang valid

Tips untuk ngerjain:
- SKU pattern: `^[A-Z]{3}\d{4}$`
- Price pattern: `^\d+(\.\d{2})?$`
- Category pattern: `^[a-zA-Z_]+$`
- Gunakan counter untuk hitung produk valid
*/
func number3csv() {
	csvString := "Laptop,LAP1234,15000000,elektronik\n" +
		"Mouse,MSE123,250000,elektronik_aksesoris\n" +
		"Keyboard,KBD5678,500000,elektronik aksesoris\n" +
		"Monitor,MON9999,3500000.50,display\n" +
		"Headset,HED1234,750000,audio_device\n"

	skuPattern := regexp.MustCompile(`^[A-Z]{3}\d{4}$`)
	pricePattern := regexp.MustCompile(`^\d+(\.\d{2})?$`)
	categoryPattern := regexp.MustCompile(`^[a-zA-Z_]+$`)

	reader := csv.NewReader(strings.NewReader(csvString))
	validCount := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if len(record) >= 4 {
			productName := record[0]
			sku := record[1]
			price := record[2]
			category := record[3]

			isSkuValid := skuPattern.MatchString(sku)
			isPriceValid := pricePattern.MatchString(price)
			isCategoryValid := categoryPattern.MatchString(category)

			if isSkuValid && isPriceValid && isCategoryValid {
				fmt.Printf("✓ %s - %s - %s - %s\n", productName, sku, price, category)
				validCount++
			}
		}
	}

	fmt.Println("---")
	fmt.Printf("✓ Total produk valid: %d\n", validCount)
}

func main() {
	// CsvReader()
	// CsvWriter()

	// Uncomment soal yang sedang dikerjakan:
	// number1csv()
	// number2csv()
	number3csv()
}
