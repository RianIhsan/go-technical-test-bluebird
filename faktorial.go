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
)

// Fungsi main sebagai titik masuk utama program.
func main() {
	// Saya mendeklarasikan variabel untuk menyimpan bilangan bulat yang akan dimasukkan oleh pengguna.
	var angka int

	// Saya meminta pengguna untuk memasukkan sebuah bilangan bulat.
	fmt.Print("Masukkan sebuah bilangan bulat positif: ")

	// Saya menggunakan fungsi fmt.Scan untuk membaca input pengguna dan menyimpannya dalam variabel angka.
	_, err := fmt.Scan(&angka)

	// Validasi agar input adalah bilangan bulat positif.
	if err != nil || angka < 0 {
		fmt.Println("Error: Silakan masukkan bilangan bulat positif.")
		return
	}

	// Saya memanggil fungsi hitungFaktorial untuk menghitung faktorial dari angka yang dimasukkan.
	faktorial := hitungFaktorial(angka)

	// Saya mencetak hasil perhitungan faktorial.
	fmt.Printf("Faktorial dari %d adalah: %d\n", angka, faktorial)
}

// Fungsi hitungFaktorial untuk menghitung faktorial dari sebuah bilangan bulat positif.
func hitungFaktorial(n int) int {
	// Saya menggunakan variabel result untuk menyimpan hasil perhitungan faktorial.
	result := 1

	// Saya menggunakan loop for untuk mengalikan semua bilangan bulat positif dari 1 hingga n.
	for i := 1; i <= n; i++ {
		result *= i
	}

	// Saya mengembalikan nilai hasil perhitungan faktorial.
	return result
}
