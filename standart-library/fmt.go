package main

import "fmt"

/*
=== SOAL LATIHAN PACKAGE fmt ===

SOAL 1 — Format dasar (Printf/Sprintf)
Buat variabel untuk data seorang mahasiswa: nama (string), nim (string), ipk (float64), lulus (bool).
Tampilkan dalam satu kalimat menggunakan fmt.Printf, dengan ketentuan:
- IPK ditampilkan dengan 2 angka di belakang koma (%.2f)
- Status lulus ditampilkan sebagai %t
- Gunakan juga fmt.Sprintf untuk menyimpan hasil format ke sebuah variabel string sebelum dicetak.
*/

func checkIsPassed(status bool) string {
	if !status {
		return "tidak lulus"
	}

	return "lulus"
}

func number1() {
	name := "Jokowi"
	nim := "6702703132"
	ipk := 2.55
	ipkFormatted := fmt.Sprintf("%.2f", ipk)
	isPassed := false
	isPassedText := checkIsPassed(isPassed)

	fmt.Printf("Si %s dengan NIM %s dan IPK %s dinyatakan %s\n", name, nim, ipkFormatted, isPassedText)
}

/*
SOAL 2 — Input dari user (Scan/Scanln)
Buat program yang meminta user memasukkan nama dan umur lewat input terminal menggunakan fmt.Scan atau fmt.Scanln.
Setelah itu:
- Jika umur >= 17, cetak "<nama> sudah boleh memilih di pemilu"
- Jika belum, cetak "<nama> belum boleh memilih di pemilu"
Gunakan fmt.Sprintf untuk menyusun pesan sebelum dicetak dengan fmt.Println.
*/
func isEligible(age int) string {
	if age < 17 {
		return "belum"
	}

	return "sudah"
}
func number2() {
	fmt.Print("Masukkan nama:")
	var name string
	fmt.Scan(&name)

	fmt.Print("Masukkan umur:")
	var age int
	fmt.Scan(&age)

	isEligibleText := isEligible(age)
	message := fmt.Sprintf("%s %s boleh memilih di pemilu", name, isEligibleText)

	fmt.Println(message)
}

/*
SOAL 3 — Verb lanjutan & error formatting (%v, %+v, %T, Errorf)
Buat sebuah struct Produk dengan field Nama string, Harga int, Stok int.
Lalu:
 1. Cetak struct-nya menggunakan %v dan %+v, jelaskan lewat komentar apa bedanya.
 2. Cetak tipe data struct tersebut menggunakan %T.
 3. Buat fungsi CekStok(p Produk) error yang mengembalikan error dengan fmt.Errorf (format: "stok %s habis")
    jika Stok == 0, lalu tangani error tersebut di main dengan if err != nil.
*/

type Product struct {
	Nama  string
	Harga int
	Stok  int
}

func CekStok(p Product) error {
	if p.Stok == 0 {
		return fmt.Errorf("stok %s habis", p.Nama)
	}

	return nil
}
func number3() {
	p := Product{
		Nama:  "Shampoo",
		Harga: 10000,
		Stok:  10,
	}

	fmt.Printf("%v\n", p)  // output ringkas: {Shampoo 10000 10}
	fmt.Printf("%+v\n", p) // output dengan field name: {Nama:Shampoo Harga:10000 Stok:10}
	fmt.Printf("%T\n", p)  // output tipe: main.Product

	if err := CekStok(p); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("stok aman")
	}
}

func main() {
	fmt.Println("Hello, World!")

	firstName := "John"
	lastName := "Doe"
	age := 50

	fmt.Printf("My name is %s %s, I am %d years old\n", firstName, lastName, age)

	number1()
	number2()
	number3()
}
