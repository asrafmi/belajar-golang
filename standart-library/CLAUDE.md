# Belajar Golang - Standard Library

Dokumentasi format dan panduan untuk membuat soal di repository ini.

## Format Pembuatan Soal

### Struktur File

Setiap file soal harus memiliki struktur:
1. **Package declaration**
2. **Import statements**
3. **Soal dalam comment block**
4. **Function untuk setiap soal**
5. **Main function** untuk menjalankan soal

### Template Soal

```go
/*
=== SOAL N: Judul Soal ===
Deskripsi singkat tentang apa yang harus dikerjakan:
1. Point pertama
2. Point kedua
3. Point ketiga

Input: contoh input
Output contoh:
$ go run file.go -flag="value"
✓ Output 1
✓ Output 2

Bonus Challenge:
- Challenge tambahan (opsional)

Tips untuk ngerjain:
- Tips/hint 1
- Tips/hint 2
- Gunakan method X untuk Y
*/

func numberNfunction() {
	// TODO: Tulis kode kamu di sini
}
```

### Konvensi Naming

- **File**: `<topic>.go` (contoh: `strings.go`, `list.go`, `math.go`)
- **Function**: `number<N><topic>()` (contoh: `number1string()`, `number2list()`)
- **Variable**: snake_case untuk local, camelCase untuk package-level

### Output Format

Gunakan emoji dan format standar:
- `✓` untuk output berhasil / informasi
- `✗` untuk error / invalid
- `---` sebagai separator antar soal
- Newline sebelum soal dimulai

**Contoh:**
```
✓ List awal (10 elemen): 1 2 3 4 5 6 7 8 9 10
✓ List akhir (9 elemen): 1 2 4 5 5.5 6 8 9 10
---
```

### Expected Output

- Tulis dengan akurat apa yang **seharusnya** dihasilkan
- Include semua output dari program, bukan hanya snippet
- Jika ada bonus challenge, include outputnya juga
- Update expected output jika ada fix/perubahan

### Main Function

```go
func main() {
	number1function()
	number2function()
	number3function()
}
```

Uncomment sesuai soal yang sedang dikerjakan.

## Tips Umum

1. **Jangan bikin soal terlalu kompleks** - Fokus pada 1 konsep utama per soal
2. **Include test case yang jelas** - Berikan input konkret dan expected output
3. **Gunakan helper functions** - Abstrak logic kompleks ke fungsi terpisah
4. **Comment adalah dokumentasi** - Jelaskan WHY, bukan WHAT
5. **Verifikasi sebelum commit** - Pastikan kode berjalan dan outputnya sesuai

## Variasi Tingkat Kesulitan

- **Mudah (Soal 1)**: Basic usage dari package, single concept
- **Sedang (Soal 2)**: Combine multiple concepts, problem solving
- **Sulit (Soal 3)**: Advanced patterns, design implementation
- **Bonus**: Optional challenge untuk yang ingin lebih dalam
