package main

import (
	"flag"
	"fmt"
	"standart-library/helper"
	"strconv"
	"time"
)

/*
=== SOAL 1: Age Calculator ===
Buat program yang menghitung umur seseorang berdasarkan tanggal lahir:
- Flag `-birthdate` (string, required) - format "2006-01-02" (contoh: "2000-05-17")
- Gunakan time.Parse() untuk parse tanggal lahir
- Hitung umur dalam tahun (bandingkan dengan time.Now())
- Hitung juga sudah berapa hari sejak lahir (gunakan Sub() dan Hours())
- Tampilkan hari ulang tahun berikutnya (bulan & tanggal yang sama, tahun depan jika sudah lewat tahun ini)

Input: "-birthdate=2000-05-17"

Output contoh:
$ go run time.go -birthdate=2000-05-17
✓ Tanggal lahir: 17 May 2000
✓ Umur: 26 tahun
✓ Total hari hidup: 9585 hari
✓ Ulang tahun berikutnya: 17 May 2027
---

$ go run time.go -birthdate=2026-99-99
✗ Error: format tanggal tidak valid, gunakan YYYY-MM-DD

Tips:
  - Gunakan time.Parse("2006-01-02", value) untuk parsing
  - Gunakan .Year(), .Month(), .Day() untuk ambil komponen tanggal
  - Gunakan time.Since() atau now.Sub(birthdate) untuk selisih waktu
  - Hati-hati: pengurangan tahun sederhana (now.Year() - birth.Year()) belum tentu akurat,
    cek dulu apakah ulang tahun tahun ini sudah lewat atau belum
*/
func number1time() {
	birthdate := flag.String("birthdate", "", "birthdate")
	flag.Parse()

	if !helper.IsStringExist(*birthdate) {
		fmt.Println("Error: Birthdate is required!")
		return
	}

	formatter := "2006-01-02"
	parsedBirthdate, err := time.Parse(formatter, *birthdate)
	if err != nil {
		fmt.Println("Error: format tanggal tidak valid, gunakan YYYY-MM-DD")
		return
	}

	date, month, year := parsedBirthdate.Day(), parsedBirthdate.Month(), parsedBirthdate.Year()
	normalizedBirthdate := fmt.Sprintf("%s %s %s\n", strconv.Itoa(date), month.String(), strconv.Itoa(year))
	sub := time.Now().Sub(parsedBirthdate)
	totalDays := int(sub.Hours() / 24)

	thisYearBirthday := time.Date(time.Now().Year(), parsedBirthdate.Month(), parsedBirthdate.Day(), 0, 0, 0, 0, time.Local)
	isCurrentBirthdayPassed := !time.Now().Before(thisYearBirthday)
	nextBirthdayYear := time.Now().Year()
	age := time.Now().Year() - year
	if isCurrentBirthdayPassed {
		nextBirthdayYear++
	} else {
		age--
	}

	fmt.Printf("✓ Tanggal lahir: %s", normalizedBirthdate)
	fmt.Printf("✓ Umur: %d tahun\n", age)
	fmt.Printf("✓ Total hari hidup: %d\n", totalDays)
	fmt.Printf("✓ Ulang tahun berikutnya: %d %s %d\n", date, month, nextBirthdayYear)
	fmt.Println("Sudah ultah tahun ini?", isCurrentBirthdayPassed)
}

/*
=== SOAL 2: Meeting Scheduler ===
Buat program yang mengecek apakah sebuah jadwal meeting bentrok dengan jadwal lain:
- Flag `-start1`, `-end1` (string, required) - format "15:04" (jam:menit), jadwal meeting pertama
- Flag `-start2`, `-end2` (string, required) - format "15:04", jadwal meeting kedua
- Gunakan time.Parse() dengan layout "15:04"
- Dua jadwal dianggap bentrok jika ada irisan waktu (overlap)
- Gunakan method Before() dan After() untuk membandingkan
- Tampilkan durasi masing-masing meeting (gunakan Sub())

Input: "-start1=09:00 -end1=10:30 -start2=10:00 -end2=11:00"

Output contoh:
$ go run time.go -start1=09:00 -end1=10:30 -start2=10:00 -end2=11:00
✓ Meeting 1: 09:00 - 10:30 (durasi: 1h30m0s)
✓ Meeting 2: 10:00 - 11:00 (durasi: 1h0m0s)
✗ Jadwal BENTROK! Overlap dari 10:00 sampai 10:30

$ go run time.go -start1=09:00 -end1=10:00 -start2=10:00 -end2=11:00
✓ Meeting 1: 09:00 - 10:00 (durasi: 1h0m0s)
✓ Meeting 2: 10:00 - 11:00 (durasi: 1h0m0s)
✓ Jadwal AMAN, tidak ada bentrok

Bonus Challenge:
- Tambahkan flag `-buffer` (menit) untuk jarak minimal antar meeting (misal butuh jeda 15 menit)

Tips:
- Gunakan time.Parse("15:04", value)
- Overlap terjadi jika start1 < end2 DAN start2 < end1
- Durasi = end.Sub(start), hasilnya time.Duration yang sudah punya format String() bagus
*/
func number2time() {
	// TODO: Tulis kode kamu di sini
}

/*
=== SOAL 3: Countdown & Timezone Converter ===
Buat program event countdown yang mendukung banyak timezone:
- Flag `-event` (string, required) - waktu event format RFC3339 (contoh: "2026-12-31T23:59:59+07:00")
- Flag `-timezone` (default: "Asia/Jakarta") - timezone tujuan untuk ditampilkan (contoh: "America/New_York", "UTC")
- Parse waktu event dengan time.Parse(time.RFC3339, ...)
- Konversi waktu event ke timezone tujuan pakai time.LoadLocation() dan In()
- Hitung selisih waktu dari sekarang ke event (jika event di masa depan) menggunakan time.Until()
- Breakdown selisih waktu menjadi hari, jam, menit, detik (bukan cuma total Duration mentah)
- Jika event sudah lewat, tampilkan pesan "Event sudah berlalu"

Input: "-event=2026-12-31T23:59:59+07:00 -timezone=America/New_York"

Output contoh:
$ go run time.go -event=2026-12-31T23:59:59+07:00 -timezone=America/New_York
✓ Event (WIB)      : 31 Dec 2026 23:59:59 +0700
✓ Event (New York)  : 31 Dec 2026 11:59:59 -0500
✓ Countdown: 149 hari 3 jam 12 menit 5 detik lagi
---

$ go run time.go -event=2020-01-01T00:00:00Z -timezone=UTC
✗ Event sudah berlalu (6 tahun yang lalu)

$ go run time.go -event=2026-12-31T23:59:59+07:00 -timezone=Mars/Phobos
✗ Error: timezone tidak dikenali: Mars/Phobos

Bonus Challenge:
- Tambahkan opsi format output custom lewat flag `-format` menggunakan reference time Go (Mon Jan 2 15:04:05 MST 2006)

Tips:
  - Gunakan time.LoadLocation(name) lalu t.In(location) untuk convert timezone, jangan lupa handle error jika timezone invalid
  - Gunakan time.Until(event) untuk dapat Duration ke masa depan (negatif kalau sudah lewat)
  - Breakdown Duration: gunakan operasi div/mod terhadap time.Hour, time.Minute, time.Second,
    atau field d.Hours(), d.Minutes() lalu di-truncate
*/
func number3time() {
	// TODO: Tulis kode kamu di sini
}

func demoCobaTime() {
	now := time.Now()

	fmt.Println("now", now.Local())

	utc := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.UTC)
	fmt.Println("utc", utc)

	parse, _ := time.Parse(time.RFC3339, "2026-01-02T15:04:05Z")
	fmt.Println("parse", parse)

	formatter := "2006-01-02 15:04:05"
	value := "2020-10-10 23:59:59"

	parsed, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Parsed:", parsed)
	}

	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond
	var duration3 time.Duration = duration1 + duration2
	fmt.Println("duration1", duration1)
	fmt.Println("duration2", duration2)
	fmt.Printf("duration3 %d\n", duration3)
}

func main() {
	// Pilih soal yang mau dijalankan:
	number1time() // Soal 1: Age Calculator
	// number2time() // Soal 2: Meeting Scheduler
	// number3time() // Soal 3: Countdown & Timezone Converter

	// demoCobaTime()
}
