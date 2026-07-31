package main

import (
	"container/list"
	"fmt"
	"standart-library/helper"
)

/*
=== SOAL 1: Basic List Operations ===
Buat program yang:
1. Membuat sebuah linked list baru
2. Menambahkan 5 nama bulan (Januari - Mei) ke list
3. Menampilkan semua nama bulan dari depan ke belakang
4. Menampilkan semua nama bulan dari belakang ke depan
5. Menghitung dan tampilkan jumlah elemen dalam list

Output contoh:
✓ Dari depan ke belakang:
Januari
Februari
Maret
April
Mei

✓ Dari belakang ke depan:
Mei
April
Maret
Februari
Januari

✓ Total elemen: 5

Tips untuk ngerjain:
- Gunakan list.New() untuk membuat list baru
- Gunakan PushBack() untuk menambah elemen di akhir
- Gunakan Front() untuk mulai dari elemen pertama
- Gunakan Back() untuk mulai dari elemen terakhir
- Gunakan Next() dan Prev() untuk traverse
- Gunakan Len() untuk menghitung jumlah elemen
*/

func number1list() {
	fmt.Println("\nSOAL NO.1")
	monthsList := list.New()
	monthsSlice := []string{"Januari", "Februari", "Maret", "April", "Mei"}

	monthsList = helper.InsertListFromArray(monthsSlice)

	fmt.Println("✓ Dari depan ke belakang:")
	for e := monthsList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("\n")

	fmt.Println("✓ Dari belakang ke depan:")
	for e := monthsList.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	fmt.Printf("✓ Total elemen: %d\n", monthsList.Len())
	fmt.Println("---")
}

/*
=== SOAL 2: List Manipulation ===
Buat program yang:
1. Membuat list berisi 10 angka (1-10)
2. Cari elemen dengan value 5
3. Jika ditemukan, insert "5.5" setelah elemen 5
4. Hapus elemen dengan value 3 dan 7
5. Tampilkan list sebelum dan sesudah manipulasi dengan jumlah elemennya

Output contoh:
✓ List awal (10 elemen): 1 2 3 4 5 6 7 8 9 10
✓ List akhir (9 elemen): 1 2 4 5 5.5 6 8 9 10

Bonus Challenge:
- Cari indeks dari setiap elemen (bukan pointer, tapi nomor urutan)
- Hitung total value dari semua elemen

Tips untuk ngerjain:
- Gunakan PushBack() untuk menambah banyak elemen
- Gunakan Remove() untuk menghapus elemen
- Gunakan InsertAfter() untuk insert setelah elemen tertentu
- Simpan pointer element saat traverse untuk digunakan Remove()
*/

func number2list() {
	fmt.Println("\nSOAL NO.2")
	numbersSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbersList := list.New()

	numbersList = helper.InsertListFromArray(numbersSlice)

	fmt.Print("✓ List awal (")
	fmt.Print(numbersList.Len())
	fmt.Print(" elemen): ")
	for e := numbersList.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	var numberToRemove []*list.Element
	var numberToInsertAfter []*list.Element

	for e := numbersList.Front(); e != nil; e = e.Next() {
		if e.Value == 3 || e.Value == 7 {
			numberToRemove = append(numberToRemove, e)
		}
		if e.Value == 5 {
			numberToInsertAfter = append(numberToInsertAfter, e)
		}
	}

	for _, e := range numberToRemove {
		numbersList.Remove(e)
	}

	for _, e := range numberToInsertAfter {
		numbersList.InsertAfter("5.5", e)
	}

	fmt.Println()

	fmt.Print("✓ List Akhir (")
	fmt.Print(numbersList.Len())
	fmt.Print(" elemen): ")
	for e := numbersList.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println("\n---")
}

/*
=== SOAL 3: Queue & Stack Implementation ===
Buat program menggunakan list untuk:
1. Implement Queue (FIFO - First In First Out)
   - Enqueue: tambah elemen di belakang
   - Dequeue: ambil elemen dari depan

2. Implement Stack (LIFO - Last In First Out)
   - Push: tambah elemen di belakang
   - Pop: ambil elemen dari belakang

Buat 2 fungsi terpisah untuk Queue dan Stack dengan test case:

Queue test:
- Enqueue: A, B, C
- Dequeue 2 kali
- Enqueue: D
- Dequeue sampai kosong
Expected output: A, B, C, D

Stack test:
- Push: 1, 2, 3
- Pop 2 kali
- Push: 4
- Pop sampai kosong
Expected output: 3, 2, 4, 1

Output contoh:
✓ Queue operations:
Dequeue: A
Dequeue: B
Dequeue: D

✓ Stack operations:
Pop: 3
Pop: 2
Pop: 4

Tips untuk ngerjain:
- Queue gunakan PushBack() untuk enqueue dan Front() + Remove() untuk dequeue
- Stack gunakan PushBack() untuk push dan Back() + Remove() untuk pop
- Cek IsEmpty dengan menggunakan Len() == 0
*/

func queueTest() {
	fmt.Println("\nSOAL NO.3 - Queue Test (FIFO)")
	queue := list.New()

	// Enqueue: A, B, C
	fmt.Println("✓ Enqueue: A, B, C")
	queue.PushBack("A")
	queue.PushBack("B")
	queue.PushBack("C")

	// Dequeue 2 kali
	fmt.Println("✓ Dequeue 2 kali:")
	if queue.Len() > 0 {
		e := queue.Front()
		fmt.Printf("Dequeue: %v\n", e.Value)
		queue.Remove(e)
	}
	if queue.Len() > 0 {
		e := queue.Front()
		fmt.Printf("Dequeue: %v\n", e.Value)
		queue.Remove(e)
	}

	// Enqueue: D
	fmt.Println("✓ Enqueue: D")
	queue.PushBack("D")

	// Dequeue sampai kosong
	fmt.Println("✓ Dequeue sampai kosong:")
	for queue.Len() > 0 {
		e := queue.Front()
		fmt.Printf("Dequeue: %v\n", e.Value)
		queue.Remove(e)
	}
	fmt.Println("---")
}

func stackTest() {
	fmt.Println("\nSOAL NO.3 - Stack Test (LIFO)")
	stack := list.New()

	// Push: 1, 2, 3
	fmt.Println("✓ Push: 1, 2, 3")
	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)

	// Pop 2 kali
	fmt.Println("✓ Pop 2 kali:")
	if stack.Len() > 0 {
		e := stack.Back()
		fmt.Printf("Pop: %v\n", e.Value)
		stack.Remove(e)
	}
	if stack.Len() > 0 {
		e := stack.Back()
		fmt.Printf("Pop: %v\n", e.Value)
		stack.Remove(e)
	}

	// Push: 4
	fmt.Println("✓ Push: 4")
	stack.PushBack(4)

	// Pop sampai kosong
	fmt.Println("✓ Pop sampai kosong:")
	for stack.Len() > 0 {
		e := stack.Back()
		fmt.Printf("Pop: %v\n", e.Value)
		stack.Remove(e)
	}
	fmt.Println("---")
}

func number3list() {
	queueTest()
	stackTest()
}

func demoCobaList() {
	var data *list.List = list.New()

	data.PushBack("Joko")
	data.PushBack("Widodo")
	data.PushBack("Subianto")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {
	// demoCobaList()
	number1list()
	number2list()
	number3list()
}
