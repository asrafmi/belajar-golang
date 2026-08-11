package main

import (
	"fmt"
	"slices"
)

/*
=== SOAL 1: Cari Nilai Min dan Max ===
Deskripsi:
1. Diberikan slice of integers
2. Gunakan slices.Min() untuk cari nilai minimum
3. Gunakan slices.Max() untuk cari nilai maximum
4. Tampilkan hasil dalam format: Min: X, Max: Y

Input:
scores := []int{85, 92, 78, 95, 88, 76, 91}

Output yang diharapkan:
✓ Min: 76
✓ Max: 95

Tips untuk ngerjain:
- Gunakan slices.Min(slice) untuk nilai minimum
- Gunakan slices.Max(slice) untuk nilai maximum
- Slice tidak boleh kosong
*/
func number1slices() {
	// TODO: Tulis kode kamu di sini
}

/*
=== SOAL 2: Cek Keberadaan dan Posisi Elemen ===
Deskripsi:
1. Diberikan slice of strings (daftar nama produk)
2. Cek apakah produk tertentu ada di dalam slice menggunakan slices.Contains()
3. Jika ada, cari posisi/index menggunakan slices.Index()
4. Tampilkan hasil: "Produk ditemukan di posisi X" atau "Produk tidak ditemukan"

Input:
products := []string{"Laptop", "Mouse", "Keyboard", "Monitor", "Headset"}
target := "Keyboard"

Output yang diharapkan:
✓ Keyboard ditemukan di posisi 2

Bonus:
- Cari semua posisi jika ada duplikat menggunakan slices.Index() dalam loop

Tips untuk ngerjain:
- Gunakan slices.Contains(slice, element) untuk cek keberadaan
- Gunakan slices.Index(slice, element) untuk cari posisi (return -1 jika tidak ada)
- Gunakan if-else untuk handle ada/tidak ada
*/
func number2slices() {
	// TODO: Tulis kode kamu di sini
}

/*
=== SOAL 3: Sorting dan Reverse ===
Deskripsi:
1. Diberikan slice unsorted
2. Gunakan slices.Sort() untuk sort ascending
3. Gunakan slices.Reverse() untuk reverse order
4. Tampilkan hasil setiap step

Input:
numbers := []int{45, 12, 78, 23, 56, 34, 89, 11}

Output yang diharapkan:
✓ Original: [45 12 78 23 56 34 89 11]
✓ Sorted (Ascending): [11 12 23 34 45 56 78 89]
✓ Reversed (Descending): [89 78 56 45 34 23 12 11]

Tips untuk ngerjain:
- Gunakan slices.Sort(slice) untuk sort ascending
- Gunakan slices.Reverse(slice) untuk reverse
- slices.Reverse() bersifat in-place (mengubah slice original)
- Copy slice jika perlu preserve original
*/
func number3slices() {
	// TODO: Tulis kode kamu di sini
	numbers := []int{45, 12, 78, 23, 56, 34, 89, 11}

	fmt.Println("✓ Original: ", numbers)
	slices.Sort(numbers)

	fmt.Println("✓ Sorted (Ascending)", numbers)

	NumbersToBeReversed := slices.Clone(numbers)
	slices.Reverse(NumbersToBeReversed)
	fmt.Println("✓ Sorted (Descending)", NumbersToBeReversed)
}

/*
=== SOAL 4: Insert dan Delete Elemen ===
Deskripsi:
1. Diberikan slice awal
2. Insert elemen baru di posisi tertentu menggunakan slices.Insert()
3. Delete elemen di range tertentu menggunakan slices.Delete()
4. Tampilkan hasil setiap operasi

Input:
fruits := []string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
- Insert "Blueberry" di posisi 2
- Delete elemen di index 3 sampai 4

Output yang diharapkan:
✓ Original: [Apple Banana Cherry Date Elderberry]
✓ Setelah insert: [Apple Banana Blueberry Cherry Date Elderberry]
✓ Setelah delete: [Apple Banana Blueberry Elderberry]

Tips untuk ngerjain:
- slices.Insert(slice, index, elements...) → return slice baru
- slices.Delete(slice, start, end) → delete range [start:end)
- Insert/Delete return slice baru, jangan lupa assign kembali
- Index mulai dari 0
*/

func number4slices() {
	// TODO: Tulis kode kamu di sini
	fruits := []string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}

	fmt.Println("✓ Original: ", fruits)

	// Insert "Blueberry" di posisi 2
	fruits = slices.Insert(fruits, 1, "Blueberry")
	fmt.Println("✓ Setelah insert: ", fruits)

	// Delete elemen di index 3 sampai 4
	fruits = slices.Delete(fruits, 3, 4)
	fmt.Println("✓ Setelah delete: ", fruits)
}

/*
=== SOAL 5: Clone dan Membandingkan Slice ===
Deskripsi:
1. Clone slice menggunakan slices.Clone()
2. Modifikasi clone tanpa mempengaruhi original
3. Bandingkan 2 slice menggunakan slices.Equal()
4. Tampilkan perbandingan

Input:
original := []int{10, 20, 30, 40, 50}

Output yang diharapkan:
✓ Original: [10 20 30 40 50]
✓ Clone: [10 20 30 40 50]
✓ Sebelum modifikasi - Equal: true
✓ Modifikasi clone index 2 menjadi 99
✓ Sesudah modifikasi - Equal: false
✓ Original: [10 20 30 40 50]
✓ Clone (modified): [10 20 99 40 50]

Tips untuk ngerjain:
- slices.Clone(slice) → return shallow copy
- slices.Equal(slice1, slice2) → compare 2 slices
- Modifikasi elemen di clone tidak mempengaruhi original
- Equal() return true jika semua elemen sama
*/
func number5slices() {
	// TODO: Tulis kode kamu di sini
}

func main() {
	// Uncomment soal yang sedang dikerjakan:
	// number1slices()
	// number2slices()
	// number3slices()
	number4slices()
	// number5slices()
}
