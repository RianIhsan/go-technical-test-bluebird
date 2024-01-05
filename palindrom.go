// Package main merupakan package utama yang digunakan untuk membuat program yang dapat dijalankan.
package main

/*
Import adalah sintaks untuk mengakses package atau
library lain yang dibutuhkan dalam program.
Pada contoh ini, kita menggunakan package fmt untuk
operasi input/output.
*/
import (
	"fmt"
	"strings"
	/* Saya mengonversi string menjadi lowercase menggunakan
	fungsi strings.ToLower untuk memastikan perbandingan tidak case-sensitive.*/)

// Fungsi main sebagai titik masuk utama program.
func main() {
	// Saya mendeklarasikan variabel untuk menyimpan string yang akan dimasukkan oleh pengguna.
	var inputString string

	// Saya meminta pengguna untuk memasukkan sebuah string.
	fmt.Print("Masukkan sebuah string: ")

	// Saya menggunakan fungsi fmt.Scan untuk membaca input pengguna dan menyimpannya dalam variabel inputString.
	fmt.Scan(&inputString)

	// Saya memanggil fungsi isPalindrom untuk memeriksa apakah string tersebut adalah palindrom atau tidak.
	if isPalindrom(inputString) {
		fmt.Println("Kata tersebut adalah palindrom.")
	} else {
		fmt.Println("Kata tersebut bukan palindrom.")
	}
}

// Fungsi isPalindrom untuk memeriksa apakah sebuah string adalah palindrom atau tidak.
func isPalindrom(str string) bool {
	// Saya mengonversi string menjadi lowercase untuk memastikan perbandingan tidak case-sensitive.
	str = strings.ToLower(str)

	// Saya menggunakan loop untuk memeriksa karakter dari depan dan belakang.
	// Jika terdapat karakter yang tidak sesuai, maka string bukan palindrom.
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}

	// Jika berhasil melewati loop tanpa kembali false, maka string adalah palindrom.
	return true
}
