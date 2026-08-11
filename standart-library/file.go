package main

import (
	"fmt"
	"standart-library/helper"
	"strings"
)

const OUTPUT_FOLDER = "output"

/*
=== SOAL 1: Buat dan Baca File ===
Deskripsi:
1. Buat file baru dengan nama "notes.txt"
2. Tulis pesan ke file menggunakan helper.CreateNewFile()
3. Baca file menggunakan helper.ReadFile()
4. Tampilkan isi file yang terbaca
5. Handle error jika ada

Input:
- Nama file: "notes.txt"
- Pesan: "Belajar Go sangat menyenangkan!"

Output yang diharapkan:
✓ File berhasil dibuat
✓ Isi file: Belajar Go sangat menyenangkan!

Tips untuk ngerjain:
- Gunakan helper.CreateNewFile(filename, message) untuk buat file
- Gunakan helper.ReadFile(filename) untuk baca file
- helper.ReadFile() return (string, error)
- Check error sebelum menggunakan hasil
- Gunakan defer atau pastikan file ditutup
*/
func number1file() {
	_, err := helper.CheckFolderExistance(OUTPUT_FOLDER)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = helper.CreateNewFile(OUTPUT_FOLDER, "notes.txt", "Belajar Go sangat menyenangkan!")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println("✓ File berhasil dibuat")

	file, err := helper.ReadFile(OUTPUT_FOLDER, "notes.txt")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println("✓ Isi file: " + file)
}

/*
=== SOAL 2: Validasi dan Simpan Data Email ===
Deskripsi:
1. Validasi email menggunakan helper.IsEmailValid()
2. Jika valid, simpan email ke file "emails.txt"
3. Jika tidak valid, tampilkan error message
4. Baca file dan tampilkan semua email yang tersimpan
5. Handle error jika ada

Input:
- Email 1: "john@example.com" (valid)
- Email 2: "alice.invalid" (invalid - tidak ada @)
- Email 3: "bob@gmail.com" (valid)

Output yang diharapkan:
✓ john@example.com - Valid, disimpan ke file
✗ alice.invalid - Invalid: should contains '@'
✓ bob@gmail.com - Valid, disimpan ke file
---
✓ Email tersimpan di file: john@example.com

Tips untuk ngerjain:
- Gunakan helper.IsEmailValid(email) return (bool, string)
- Perhatikan error message dari helper.IsEmailValid()
- Gunakan helper.CreateNewFile() untuk simpan
- Gunakan helper.ReadFile() untuk baca
- Validasi email satu per satu
*/
func number2file() {
	_, err := helper.CheckFolderExistance(OUTPUT_FOLDER)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	emails := []string{"john@example.com", "alice.invalid", "bob@gmail.com"}

	var output string

	for _, email := range emails {
		isEmailValid, msg := helper.IsEmailValid(email)
		if !isEmailValid {
			output += "✗ " + email + " - Invalid: " + msg + "\n"
			continue
		}

		output += "✓ " + email + " - Valid, disimpan ke file\n"
	}

	err = helper.CreateNewFile(OUTPUT_FOLDER, "emails.txt", output)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println("✓ File berhasil dibuat")

	file, err := helper.ReadFile(OUTPUT_FOLDER, "emails.txt")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println(file)
}

/*
=== SOAL 3: Log Aplikasi dengan Timestamp ===
Deskripsi:
1. Simulasikan aplikasi yang membuat log entries
2. Setiap log harus diformat: "[TIMESTAMP] STATUS: message"
3. Simpan log ke file "app.log"
4. Baca dan tampilkan semua log entries
5. Hitung jumlah log berdasarkan status (INFO, ERROR, WARNING)

Input:
- Log entries dengan berbagai status
  - [2026-08-11 10:15:30] INFO: Aplikasi dimulai
  - [2026-08-11 10:16:45] WARNING: Memory usage tinggi
  - [2026-08-11 10:17:10] ERROR: Database connection failed
  - [2026-08-11 10:18:00] INFO: Retry connection
  - [2026-08-11 10:18:30] INFO: Connected to database

Output yang diharapkan:
✓ Log saved to app.log
✓ Log entries:

	[2026-08-11 10:15:30] INFO: Aplikasi dimulai
	[2026-08-11 10:16:45] WARNING: Memory usage tinggi
	[2026-08-11 10:17:10] ERROR: Database connection failed
	[2026-08-11 10:18:00] INFO: Retry connection
	[2026-08-11 10:18:30] INFO: Connected to database

---
✓ Summary:

	INFO: 3
	WARNING: 1
	ERROR: 1

Tips untuk ngerjain:
- Format log: "[TIMESTAMP] STATUS: message"
- Gunakan helper.CreateNewFile() untuk simpan semua log sekaligus
- Gunakan helper.ReadFile() untuk baca
- Split log entries berdasarkan newline untuk parsing
- Gunakan strings.Count() atau loop untuk hitung status
- Pisahkan entries dengan newline character "\n"
*/
func number3file() {
	_, err := helper.CheckFolderExistance(OUTPUT_FOLDER)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	logs := "[2026-08-11 10:15:30] INFO: Aplikasi dimulai\n" +
		"[2026-08-11 10:16:45] WARNING: Memory usage tinggi\n" +
		"[2026-08-11 10:17:10] ERROR: Database connection failed\n" +
		"[2026-08-11 10:18:00] INFO: Retry connection\n" +
		"[2026-08-11 10:18:30] INFO: Connected to database"

	err = helper.CreateNewFile(OUTPUT_FOLDER, "app.log", logs)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println("✓ Log saved to app.log")

	file, err := helper.ReadFile(OUTPUT_FOLDER, "app.log")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println("✓ Log entries:")
	fmt.Println()
	fmt.Println(file)
	fmt.Println()

	infoCount := strings.Count(file, "INFO:")
	warningCount := strings.Count(file, "WARNING:")
	errorCount := strings.Count(file, "ERROR:")

	fmt.Println("---")
	fmt.Println("✓ Summary:")
	fmt.Println()
	fmt.Printf("INFO: %d\n", infoCount)
	fmt.Printf("WARNING: %d\n", warningCount)
	fmt.Printf("ERROR: %d\n", errorCount)
}

func main() {
	// Uncomment soal yang sedang dikerjakan:
	// number1file()
	// number2file()
	number3file()
}
