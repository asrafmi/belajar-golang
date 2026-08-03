package main

import (
	"container/ring"
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i int, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i int, j int) {
	// temp := userSlice[i]
	// userSlice[i] = userSlice[j]
	// userSlice[j] = temp

	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func demo() {
	users := UserSlice{
		{Name: "Jokowi", Age: 100},
		{Name: "Prabowo", Age: 200},
		{Name: "Gibran", Age: 50},
		{Name: "Ma'ruf Amin", Age: 150},
		{Name: "Teddy", Age: 40},
	}

	sort.Sort(users)
	fmt.Println(users)
}

/*
=== SOAL 1: Sort Basic - Urutkan Slice Angka & String ===
Buat program yang:
1. Membuat slice angka acak: []int{5, 2, 9, 1, 5, 6}
2. Urutkan ascending dengan sort.Ints(), tampilkan hasilnya
3. Urutkan descending menggunakan sort.Sort() + sort.Reverse()
4. Membuat slice string: []string{"jeruk", "apel", "mangga", "anggur"}
5. Urutkan alphabetically dengan sort.Strings(), tampilkan hasilnya

Output contoh:
✓ Angka ascending: [1 2 5 5 6 9]
✓ Angka descending: [9 6 5 5 2 1]
✓ Buah A-Z: [anggur apel jeruk mangga]

Tips untuk ngerjain:
- Gunakan sort.Ints() untuk urutkan []int ascending
- Gunakan sort.Sort(sort.Reverse(sort.IntSlice(data))) untuk descending
- Gunakan sort.Strings() untuk urutkan []string
*/

func number1sort() {
	fmt.Println("\nSOAL NO.1")

	angka := []int{10, 2, 3, 8, 5, 6, 7, 4, 9, 1}
	sort.Ints(angka)

	fmt.Printf("✓ Angka ascending: %v\n", angka)
	sort.Sort(sort.Reverse(sort.IntSlice(angka)))
	fmt.Printf("✓ Angka descending: %v\n", angka)

	buahs := []string{"duren", "mangga", "apel"}
	sort.Strings(buahs)
	fmt.Printf("✓ Buah A-Z: %v\n", buahs)
	fmt.Println("---")
}

/*
=== SOAL 2: Sort Custom - Urutkan Slice of Struct ===
Buat program yang:
1. Menggunakan UserSlice berisi minimal 4 User dengan Name & Age berbeda
2. Implement sort.Interface (Len, Less, Swap) untuk UserSlice berdasarkan Age
3. Urutkan users dari termuda ke tertua dengan sort.Sort()
4. Tampilkan hasil urutan (Name & Age)
5. Cari tahu apakah data sudah terurut dengan sort.IsSorted()

Output contoh:
✓ Sebelum diurutkan:
Jokowi (100)
Budi (25)
Ani (30)
Citra (18)

✓ Sesudah diurutkan (termuda ke tertua):
Citra (18)
Budi (25)
Ani (30)
Jokowi (100)

✓ IsSorted: true

Bonus Challenge:
- Urutkan berdasarkan Name (alphabetical) menggunakan sort.Slice() tanpa perlu implement Len/Less/Swap

Tips untuk ngerjain:
- Lengkapi method Len(), Less(), dan Swap() pada UserSlice supaya memenuhi sort.Interface
- Gunakan sort.Sort(data) setelah interface lengkap
- Gunakan sort.IsSorted(data) untuk verifikasi
- Untuk bonus, gunakan sort.Slice(data, func(i, j int) bool {...})
*/

func number2sort() {
	fmt.Println("\nSOAL NO.2")

	users := UserSlice{
		{Name: "Jokowi", Age: 100},
		{Name: "Prabowo", Age: 200},
		{Name: "Gibran", Age: 50},
		{Name: "Ma'ruf Amin", Age: 150},
		{Name: "Teddy", Age: 40},
	}

	fmt.Println("✓ Sebelum diurutkan:")
	for i := 0; i < users.Len(); i++ {
		fmt.Printf("%s (%d)\n", users[i].Name, users[i].Age)
	}

	fmt.Println("\n")

	sort.Sort(users)
	fmt.Println("✓ Setelah diurutkan:")
	for i := 0; i < users.Len(); i++ {
		fmt.Printf("%s (%d)\n", users[i].Name, users[i].Age)
	}

	fmt.Println("\n")

	fmt.Println("✓ Urut berdasarkan nama:")
	sort.Slice(users, func(i, j int) bool {
		return users[i].Name < users[j].Name
	})
	for i := 0; i < users.Len(); i++ {
		fmt.Printf("%s (%d)\n", users[i].Name, users[i].Age)
	}
	fmt.Printf("✓ IsSorted: %v\n", sort.IsSorted(users))
	fmt.Println("---")
}

/*
=== SOAL 3: Sort + Circular (container/ring) - Giliran Bermain Terurut ===
Buat program simulasi giliran bermain game secara melingkar:
1. Siapkan slice pemain dengan nama & skor acak, minimal 5 pemain
   contoh: {"Andi", 40}, {"Budi", 10}, {"Citra", 25}, {"Dedi", 5}, {"Eka", 15}
2. Urutkan pemain berdasarkan skor tertinggi ke terendah dengan sort.Slice()
3. Masukkan urutan pemain yang sudah terurut tadi ke dalam container/ring
   sebesar jumlah pemain menggunakan ring.New()
4. Mulai dari pemain dengan skor tertinggi (head of ring), lakukan iterasi
   melingkar sebanyak 8 kali putaran giliran (boleh lebih dari jumlah
   pemain karena sifatnya melingkar/circular) sambil print nama pemain
   dan skornya di tiap giliran
5. Tampilkan total keseluruhan skor yang "dilewati" selama 8 giliran tadi

Output contoh:
✓ Ranking pemain (skor tertinggi ke terendah):
1. Andi (40)
2. Citra (25)
3. Eka (15)
4. Budi (10)
5. Dedi (5)

✓ Giliran bermain (circular, 8 putaran):
Giliran 1: Andi (40)
Giliran 2: Citra (25)
Giliran 3: Eka (15)
Giliran 4: Budi (10)
Giliran 5: Dedi (5)
Giliran 6: Andi (40)
Giliran 7: Citra (25)
Giliran 8: Eka (15)

✓ Total skor selama 8 giliran: 175

Bonus Challenge:
- Tambahkan fitur "skip" pemain dengan skor di bawah 10 (langsung lanjut ke
  pemain berikutnya tanpa dihitung sebagai giliran)

Tips untuk ngerjain:
- Gunakan sort.Slice() untuk urutkan slice pemain sebelum dimasukkan ke ring
- Gunakan ring.New(n) lalu iterasi dengan Do() atau manual Next() untuk mengisi Value
- Gunakan data.Next() untuk berpindah giliran secara melingkar
- Ingat ring bersifat circular, jadi setelah elemen terakhir akan kembali ke elemen pertama
*/

type Player struct {
	Name  string
	Score int
}

func number3sort() {
	fmt.Println("\nSOAL NO.3")

	players := []Player{
		{Name: "Andi", Score: 40},
		{Name: "Budi", Score: 10},
		{Name: "Citra", Score: 25},
		{Name: "Dedi", Score: 5},
		{Name: "Eka", Score: 15},
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Score > players[j].Score
	})

	fmt.Println("✓ Ranking pemain (skor tertinggi ke terendah):")
	for i, p := range players {
		fmt.Printf("%d. %s (%d)\n", i+1, p.Name, p.Score)
	}

	fmt.Println("\n✓ Giliran bermain (circular, 8 putaran):")
	turnRing := ring.New(len(players))
	for _, p := range players {
		turnRing.Value = p
		turnRing = turnRing.Next()
	}

	totalScore := 0
	for i := 1; i <= 8; i++ {
		p := turnRing.Value.(Player)
		fmt.Printf("Giliran %d: %s (%d)\n", i, p.Name, p.Score)
		totalScore += p.Score
		turnRing = turnRing.Next()
	}

	fmt.Printf("\n✓ Total skor selama 8 giliran: %d\n", totalScore)

	fmt.Println("\n✓ Bonus - Giliran bermain (skip skor < 10, 8 putaran):")
	skipRing := ring.New(len(players))
	for _, p := range players {
		skipRing.Value = p
		skipRing = skipRing.Next()
	}

	skipTotalScore := 0
	for turn := 1; turn <= 8; {
		p := skipRing.Value.(Player)
		if p.Score < 10 {
			fmt.Printf("Skip: %s (%d)\n", p.Name, p.Score)
			skipRing = skipRing.Next()
			continue
		}
		fmt.Printf("Giliran %d: %s (%d)\n", turn, p.Name, p.Score)
		skipTotalScore += p.Score
		skipRing = skipRing.Next()
		turn++
	}

	fmt.Printf("\n✓ Total skor selama 8 giliran (bonus): %d\n", skipTotalScore)
	fmt.Println("---")
}

func main() {
	// demo()
	// number1sort()
	// number2sort()
	number3sort()
}
