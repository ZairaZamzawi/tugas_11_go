package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input1, input2 string
	fmt.Printf("Input Angka Pertama : ")
	fmt.Scanln(&input1)

	fmt.Printf("Input Angka Kedua: ")
	fmt.Scanln(&input2)

	var number1, number2 int
	var err1, err2 error
	number1, err1 = strconv.Atoi(input1)
	number2, err2 = strconv.Atoi(input2)

	if err1 == nil {
		defer hitung(number1, number2)
	} else {
		fmt.Println(input1, "Ini Bukan Angka")
		// Show Error
		fmt.Println(err1.Error())
		fmt.Println(err2.Error())
		fmt.Println("Silahkan Isi Angka")
	}
}

func hitung(x1 int, x2 int) {
	var tambah int = x1 + x2
	fmt.Println("Hasil Penjumlahan :", tambah)

	var kurang int = x1 - x2
	fmt.Println("Hasil Pengurangan :", kurang)

	var kali int = x1 * x2
	fmt.Println("Hasil Perkalian :", kali)

	var bagi int = x1 / x2
	fmt.Println("Hasil Pembagian :", bagi)
}
