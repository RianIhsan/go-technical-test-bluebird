// Package main merupakan package utama yang digunakan untuk membuat program yang dapat dijalankan.
package main

/*
Import adalah sintaks untuk mengakses package atau
library lain yang dibutuhkan dalam program.
Pada contoh ini, kita menggunakan package fmt untuk
operasi input/output.
*/

// pada line 4 sampai 9 merupakan contoh multiple comment
import (
	"fmt"
)

// Fungsi main sebagai titik masuk utama program.
func main() {
	// Saya mendeklarasikan variabel untuk menyimpan suhu dalam Celsius yang akan dimasukkan oleh pengguna.
	var suhuCelsius float64

	// Saya meminta pengguna untuk memasukkan suhu dalam Celsius.
	fmt.Print("Masukkan suhu dalam Celsius: ")

	// Saya menggunakan fungsi fmt.Scan untuk membaca input pengguna dan menyimpannya dalam variabel suhuCelsius.
	fmt.Scan(&suhuCelsius)

	// Saya memanggil fungsi konversiSuhuCtoF untuk mengonversi suhu dari Celsius ke Fahrenheit.
	suhuFahrenheit := konversiSuhuCtoF(suhuCelsius)

	// Saya mencetak hasil konversi suhu dalam Fahrenheit.
	fmt.Printf("Suhu setara dalam Fahrenheit: %.2f°F\n", suhuFahrenheit)
}

// Fungsi konversiSuhuCtoF untuk mengonversi suhu dari Celsius ke Fahrenheit.
func konversiSuhuCtoF(suhuCelsius float64) float64 {
	// Saya menggunakan rumus sesuai ketentuan konversi suhu dari Celsius ke Fahrenheit: F = (C * 9/5) + 32.
	suhuFahrenheit := (suhuCelsius * 9 / 5) + 32

	// Saya mengembalikan nilai suhu dalam Fahrenheit.
	return suhuFahrenheit
}
