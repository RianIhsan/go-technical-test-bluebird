// Package main merupakan package utama yang digunakan untuk membuat program yang dapat dijalankan.
package main

// Import adalah sintaks untuk mengakses package atau library lain yang dibutuhkan dalam program.
// Pada contoh ini, kita menggunakan package fmt untuk operasi input/output dan strconv untuk konversi string ke int.
import (
	"fmt"
	"strconv"
)

// Fungsi main adalah titik masuk utama program. Eksekusi program dimulai dari fungsi ini.
func main() {
	var input string

	// Minta pengguna memasukkan angka sebagai string
	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scan(&input)

	// Konversi input ke tipe data int
	// Fungsi strconv.Atoi digunakan untuk mengonversi string ke integer.
	// Jika konversi berhasil, err akan nil. Jika tidak, err akan berisi informasi kesalahan.
	inputAngka, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error: Masukkan angka bulat yang valid.")
		return
	}

	// Validasi agar angka yang dimasukkan lebih dari 1
	if inputAngka <= 1 {
		fmt.Println("Silakan masukkan angka yang lebih dari 1.")
		return
	}

	// Hitung dan cetak angka ganjil beserta jumlahnya
	// Fungsi hitungDanCetakBilanganGanjil digunakan untuk menghitung dan mencetak angka ganjil beserta jumlahnya.
	angkaGanjil, jumlahGanjil := hitungDanCetakBilanganGanjil(inputAngka)

	// Cetak hasil
	// Fungsi fmt.Printf digunakan untuk mencetak string dengan format tertentu.
	// %d digunakan untuk mencetak integer.
	// %v digunakan untuk mencetak slice.
	// \n digunakan untuk membuat baris baru.
	fmt.Printf("Angka ganjil dari 1 hingga %d adalah: %v\n", inputAngka, angkaGanjil)
	fmt.Printf("Jumlah angka ganjil dari 1 hingga %d adalah: %d\n", inputAngka, jumlahGanjil)
}

// Fungsi hitungDanCetakBilanganGanjil untuk menghitung dan mencetak angka ganjil beserta jumlahnya.
func hitungDanCetakBilanganGanjil(angka int) ([]int, int) {
	var angkaGanjil []int // Slice untuk menyimpan angka-angka ganjil
	jumlahGanjil := 0     // Variabel untuk menghitung jumlah angka ganjil

	// Loop dari 1 hingga angka yang dimasukkan
	for i := 1; i <= angka; i++ {
		// Disini saya menggunakan operasi modulo untuk mengecek apakah i adalah angka ganjil
		if i%2 == 1 {
			angkaGanjil = append(angkaGanjil, i) // Tambahkan angka ganjil ke slice
			jumlahGanjil++                       // Tambahkan 1 ke jumlah angka ganjil
		}
	}
	// Disini saya menggunakan return multiple values untuk mengembalikan slice angka ganjil dan jumlah angka ganjil
	return angkaGanjil, jumlahGanjil
}
