package main

import (
	"flag"
	"fmt"
	"math"
	"standart-library/helper"
	"strconv"
)

/*
=== SOAL 1: Calculator with Strconv ===
Buat program kalkulator yang menerima input string dan convert ke number:
- Flag `-num1` (string, required) - angka pertama
- Flag `-num2` (string, required) - angka kedua
- Flag `-operation` (default: "add") - operasi (add, subtract, multiply, divide)
- Gunakan strconv untuk convert string ke int/float64
- Handle error saat parsing
- Tampilkan hasil dengan 2 decimal places

Input: "-num1=10.5 -num2=3.2 -operation=add"

Output contoh:
$ go run math.go -num1=10.5 -num2=3.2 -operation=add
Input: 10.50 + 3.20
Result: 13.70

$ go run math.go -num1=abc -num2=5
✗ Error: invalid number format for num1: abc

Tips:
- Gunakan strconv.ParseFloat() untuk float
- Gunakan strconv.Atoi() atau strconv.ParseInt() untuk integer
- Handle error dengan if err != nil
*/
func handleIntOperation(operation *string, num1 *string, num2 *string) {
	parseNum1, errParseNum1 := strconv.Atoi(*num1)
	if errParseNum1 != nil {
		fmt.Printf("Error: invalid number format for num1: %s\n", *num1)
		return
	}

	parseNum2, errParseNum2 := strconv.Atoi(*num2)
	if errParseNum2 != nil {
		fmt.Printf("Error: invalid number format for num2: %s\n", *num2)
		return
	}

	helper.OperateInt(operation, &parseNum1, &parseNum2)
}

func handleFloatOperation(operation *string, num1 *string, num2 *string) {
	parseNum1, errParseNum1 := strconv.ParseFloat(*num1, 64)
	if errParseNum1 != nil {
		fmt.Printf("Error: invalid number format for num1: %s\n", *num1)
		return
	}

	parseNum2, errParseNum2 := strconv.ParseFloat(*num2, 64)
	if errParseNum2 != nil {
		fmt.Printf("Error: invalid number format for num2: %s\n", *num2)
		return
	}

	helper.OperateFloat(operation, &parseNum1, &parseNum2)
}
func number1math() {
	num1 := flag.String("num1", "", "num1")
	num2 := flag.String("num2", "", "num2")
	operation := flag.String("operation", "add", "operation")
	flag.Parse()

	if !helper.IsStringExist(*num1) || !helper.IsStringExist(*num2) {
		fmt.Println("Error: num1 and num2 are required!")
		return
	}

	if helper.IsInt(*num1) && helper.IsInt(*num2) {
		handleIntOperation(operation, num1, num2)
	} else {
		handleFloatOperation(operation, num1, num2)
	}
}

/*
=== SOAL 2: BMI Calculator ===
Buat program hitung BMI (Body Mass Index):
- Flag `-weight` (string, required) - berat badan dalam kg
- Flag `-height` (string, required) - tinggi badan dalam meter (contoh: 1.75)
- Formula: BMI = weight / (height * height)
- Kategori:
  * BMI < 18.5 = Underweight
  * 18.5 <= BMI < 25 = Normal weight
  * 25 <= BMI < 30 = Overweight
  * BMI >= 30 = Obese
- Gunakan math.Round() untuk pembulatan
- Output kategori dengan emoji

Input: "-weight=70 -height=1.75"

Output contoh:
$ go run math.go -weight=70 -height=1.75
Weight: 70.00 kg
Height: 1.75 m
BMI: 22.86
Category: 😊 Normal weight

$ go run math.go -weight=90 -height=1.70
Weight: 90.00 kg
Height: 1.70 m
BMI: 31.03
Category: ⚠️ Obese

Tips:
- Gunakan strconv.ParseFloat() dengan bitSize 64
- Gunakan math.Round() dan format dengan %.2f
*/

/*
=== SOAL 3: GPA Calculator with Math Operations ===
Buat program hitung GPA (Grade Point Average):
- Input: flag untuk multiple subjects (format: -subject1=name:grade:credit)
- Grade: A=4.0, B=3.0, C=2.0, D=1.0, F=0.0
- Formula GPA: (sum of (grade * credit)) / (sum of credit)
- Gunakan math.Abs() untuk nilai absolute
- Tampilkan detail setiap subject dan total GPA
- Validasi grade hanya A-F

Input: "-subject1=Math:A:3 -subject2=English:B:2 -subject3=Science:A:4"

Output contoh:
$ go run math.go -subject1="Math:A:3" -subject2="English:B:2" -subject3="Science:A:4"
Subject GPA Breakdown:
- Math (Credit: 3) → Grade A (4.00) = 12.00
- English (Credit: 2) → Grade B (3.00) = 6.00
- Science (Credit: 4) → Grade A (4.00) = 16.00

Total Credit: 9
Total Grade Points: 34.00
GPA: 3.78 ✓

$ go run math.go -subject1="Math:X:3"
✗ Error: Invalid grade 'X' for Math (must be A-F)

Tips:
- Gunakan strings.Split() untuk parse format "name:grade:credit"
- Gunakan strconv.ParseInt() untuk credit
- Gunakan strconv.ParseFloat() untuk numeric calculations
- Gunakan math.Round() untuk pembulatan GPA
- Validasi credit harus > 0
*/

func demoCoba() {
	fmt.Println("math.Ceil dari 1.40 adalah", math.Ceil(1.40))
	fmt.Println("math.Floor dari 1.40 adalah", math.Floor(1.40))
	fmt.Println("math.Round dari 1.40 adalah", math.Round(1.40))
	fmt.Println("math.Max dari 1.40 dan 1.41 adalah", math.Max(1.40, 1.41))
	fmt.Println("math.Min dari 1.40 dan 1.39 adalah", math.Min(1.40, 1.39))
}

func main() {
	// Pilih soal yang mau dijalankan:
	number1math() // Soal 1: Calculator with Strconv
	// number2math()  // Soal 2: BMI Calculator
	// number3math()  // Soal 3: GPA Calculator

	// demoCoba()
}
