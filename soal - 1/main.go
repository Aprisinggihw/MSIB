package main

import (
	"fmt"
	"strconv"
)

// Fungsi utama untuk menampilkan menu dan mengatur pilihan
func main() {
	for {
		fmt.Println("\nMenu Utama Aplikasi:")
		fmt.Println("1. Menampilkan 'Hello, World!'")
		fmt.Println("2. Operasi matematika sederhana")
		fmt.Println("3. Menyimpan dan menampilkan data pengguna")
		fmt.Println("4. Hitung faktorial(rekursif)")
		fmt.Println("5. Hitung rata-rata(variadic)")
		fmt.Println("6. Keluar")

		var pilihan int
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			cetakHelloWorld()
		case 2:
			operasiMatematika()
		case 3:
			mengelolaDataPengguna()
		case 4:
			var angka int
			fmt.Print("Masukkan angka untuk menghitung faktorial: ")
			fmt.Scan(&angka)
			fmt.Println("\n---------------------------------------------------")
			fmt.Printf("Faktorial dari %d adalah %d\n", angka, faktorial(angka))
			fmt.Println("\n---------------------------------------------------")
		case 5:
			rata_rata := rata_rata(1, 2, 3, 4, 5)
			fmt.Println("\n---------------------------------------------------")
			fmt.Println("Rata-rata dari 1, 2, 3, 4, 5 adalah:", rata_rata)
			fmt.Println("\n---------------------------------------------------")
		case 6:
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

// Fungsi untuk menampilkan 'Hello, World!'
func cetakHelloWorld() {
	fmt.Println("\n---------------------------------------------------")
	fmt.Println("Hello, World!")
	fmt.Println("\n---------------------------------------------------")
}

// Fungsi perkalian
func perkalian(angkaPertama, angkaKedua float64) float64 {
	return angkaPertama * angkaKedua
}

// Fungsi pembagian
func pembagian(angkaPertama, angkaKedua float64) float64 {
	if angkaKedua != 0 {
		return angkaPertama / angkaKedua
	}
	fmt.Println("\n---------------------------------------------------")
	fmt.Println("Pembagian dengan nol tidak diizinkan!")
	fmt.Println("\n---------------------------------------------------")
	return 0
}

// Fungsi untuk melakukan operasi matematika sederhana
func operasiMatematika() {
	var angkaPertama, angkaKedua float64
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&angkaPertama)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&angkaKedua)

	// Fungsi penjumlahan
	penjumlahan := func(x, y float64) float64 { return x + y }

	// Fungsi pengurangan
	pengurangan := func(x, y float64) float64 { return x - y }

	fmt.Println("\n---------------------------------------------------")
	fmt.Printf("Penjumlahan: %.2f\n", penjumlahan(angkaPertama, angkaKedua))
	fmt.Printf("Pengurangan: %.2f\n", pengurangan(angkaPertama, angkaKedua))
	fmt.Printf("Perkalian: %.2f\n", perkalian(angkaPertama, angkaKedua))
	fmt.Printf("Pembagian: %.2f\n", pembagian(angkaPertama, angkaKedua))
	fmt.Println("\n---------------------------------------------------")
}

// Fungsi filter berdasarkan umur
func filterDataPenggunaBerdasarkanUmur(users map[string][]string) {
	var umur string
	fmt.Print("Masukkan umur yang ingin dicari: ")
	fmt.Scan(&umur)
	fmt.Println("Pengguna dengan umur", umur, ":")
	fmt.Println("\n---------------------------------------------------")
	for nama, info := range users {
		if info[0] == umur {
			fmt.Printf("Nama: %s, Umur: %s, Hobi: %s\n", nama, info[0], info[1])
		}
	}
	fmt.Println("\n---------------------------------------------------")
}

// Fungsi filter berdasarkan hobi
func filterDataPenggunaBerdasarkanHobi(users map[string][]string) {
	var hobby string
	fmt.Print("Masukkan hobi yang ingin dicari: ")
	fmt.Scan(&hobby)
	fmt.Println("Pengguna dengan hobi", hobby, ":")
	fmt.Println("\n---------------------------------------------------")
	for name, info := range users {
		if info[1] == hobby {
			fmt.Printf("Nama: %s, Umur: %s, Hobi: %s\n", name, info[0], info[1])
		}
	}
	fmt.Println("\n---------------------------------------------------")
}

// Fungsi untuk mengelola data pengguna (menyimpan dan menampilkan)
func mengelolaDataPengguna() {
	// Map untuk menyimpan data pengguna, dengan nama sebagai key dan value berupa slice yang berisi umur dan hobi
	users := make(map[string][]string)
	for {
		fmt.Println("\n1. Tambah pengguna")
		fmt.Println("2. Tampilkan semua pengguna")
		fmt.Println("3. Tampilkan data pengguna berdasarkan umur")
		fmt.Println("4. Tampilkan data pengguna berdasarkan hobi")
		fmt.Println("5. Kembali ke menu utama")
		var choice int
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var nama, hobi string
			var umur int
			fmt.Print("Masukkan nama: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan umur: ")
			fmt.Scan(&umur)
			fmt.Print("Masukkan hobi: ")
			fmt.Scan(&hobi)

			// Konversi umur (int) ke string
			umurString := strconv.Itoa(umur)

			// Menyimpan umur dan hobi sebagai slice
			users[nama] = []string{umurString, hobi}
			fmt.Println("Pengguna berhasil ditambahkan!")
		case 2:
			fmt.Println("\n---------------------------------------------------")
			fmt.Println("Data Pengguna:")
			for nama, info := range users {
				fmt.Printf("Nama: %s, Umur: %s, Hobi: %s\n", nama, info[0], info[1])
			}
			fmt.Println("\n---------------------------------------------------")
		case 3:
			filterDataPenggunaBerdasarkanUmur(users)
		case 4:
			filterDataPenggunaBerdasarkanHobi(users)
		case 5:
			return
		default:
			fmt.Println("\n---------------------------------------------------")
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
			fmt.Println("\n---------------------------------------------------")
		}
	}
}

// Fungsi Rekursif untuk menghitung faktorial
func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

// Fungsi variadic untuk menghitung total
func total(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Fungsi variadic untuk menghitung rata-rata
func rata_rata(nums ...int) float64 {
	total := total(nums...)
	return float64(total) / float64(len(nums))
}
