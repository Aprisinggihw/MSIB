package view

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Interface Pesanan untuk mendefinisikan metode umum
type Pesanan interface {
	GetNama() string
	GetHarga() float64
	GetQty() int
	SetQty(int)
}

// Implementasi MenuItem terhadap interface Pesanan
type MenuItem struct {
	Nama  string
	Harga float64
	Qty   int
}

func (m *MenuItem) GetNama() string {
	return m.Nama
}

func (m *MenuItem) GetHarga() float64 {
	return m.Harga
}

func (m *MenuItem) GetQty() int {
	return m.Qty
}

func (m *MenuItem) SetQty(qty int) {
	m.Qty = qty
}

// Interface kosong untuk pemrosesan pesanan berbagai tipe data
func prosesData(data interface{}) {
	switch v := data.(type) {
	case string:
		fmt.Printf("Data string: %s\n", v)
	case int:
		fmt.Printf("Data int: %d\n", v)
	default:
		fmt.Printf("Tipe data tidak dikenal\n")
	}
}

var menu []*MenuItem

func tampilkanItemMenu(menu []*MenuItem) {
	fmt.Println("\n=== Menu Restoran ===")
	for i, item := range menu {
		fmt.Printf("%d. %s - Rp%.2f\n", i+1, item.Nama, item.Harga)
	}
	fmt.Println("\n========================")
}

// Fungsi untuk memproses pesanan menggunakan goroutine
func prosesPesanan(pesanan Pesanan, wg *sync.WaitGroup, ch chan Pesanan) {
	defer wg.Done()
	time.Sleep(2 * time.Second) // Simulasi pemrosesan pesanan
	ch <- pesanan
}

// Validasi harga dengan regexp
func isValidPrice(input string) bool {
	re := regexp.MustCompile(`^\d+(\.\d{1,2})?$`) // Hanya menerima angka dengan maksimal 2 digit desimal
	return re.MatchString(input)
}

// Validasi nama item menggunakan error interface
type ValidasiError struct {
	Message string
}

func (e *ValidasiError) Error() string {
	return e.Message
}

func validasiNamaItem(nama string) error {
	if strings.TrimSpace(nama) == "" {
		return &ValidasiError{Message: "Nama item tidak boleh kosong"}
	}
	return nil
}

func tanganiPanic() {
	if r := recover(); r != nil {
		fmt.Println("Error:", r)
	}
}


// Konversi harga dari string ke float64 menggunakan strconv
func konversiHarga(input string) (float64, error) {
	defer tanganiPanic() // Memastikan recover dijalankan jika terjadi panic

	if isValidPrice(input) {
		harga, err := strconv.ParseFloat(input, 64)
		if err != nil {
			panic("Kesalahan konversi: harga tidak valid")
		}
		return harga, nil
	} else {
		return 0, fmt.Errorf("format harga tidak valid. Silakan masukkan angka yang benar")
	}
}


// Encode semua pesanan
func encodedSemuaPesanan(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Fungsi Exported 
func TampilkanMenuUtama() {
	var pilihan int
	// Insialisasi item menu awal
	menu = append(menu, &MenuItem{Nama: "Nasi Goreng", Harga: 20000})
	menu = append(menu, &MenuItem{Nama: "Mie Goreng", Harga: 18000})
	menu = append(menu, &MenuItem{Nama: "Ayam Bakar", Harga: 25000})

	for {
		tampilkanItemMenu(menu)

		fmt.Println("1. Tambah Menu\n2. Hapus Menu\n3. Pesan\n4. Keluar")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahItemMenu()
		case 2:
			hapusItemMenu()
		case 3:
			menuPesan()
		case 4:
			return
		default:
			fmt.Printf("Inputan tidak valid!")
		}
	}
}

func menuPesan() {
	scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup
	ch := make(chan Pesanan, 5)
	var semuaPesanan []Pesanan

	tampilkanItemMenu(menu)

	// Input pesanan dari pengguna
	for {
		fmt.Println("\nMasukkan nama menu yang ingin dipesan (atau 'q' untuk selesai dan hitung total):")
		scanner.Scan()
		input := scanner.Text()

		if input == "q" {
			break
		}

		var menuIndex int
		cari := false
		for i, item := range menu {
			if strings.EqualFold(input, item.Nama) { // Pencarian case-insensitive
				menuIndex = i
				cari = true
				break
			}
		}

		if !cari {
			fmt.Println("Menu tidak ditemukan, silakan coba lagi.")
			continue
		}

		itemPilihan := menu[menuIndex]

		fmt.Printf("Masukkan jumlah pesanan untuk %s: ", itemPilihan.Nama)
		scanner.Scan()
		qtyInput := scanner.Text()

		qty, err := strconv.Atoi(qtyInput)
		if err != nil || qty <= 0 {
			fmt.Println("Jumlah pesanan tidak valid. Silakan coba lagi.")
			continue
		}
		itemPilihan.SetQty(qty)
		wg.Add(1)
		go prosesPesanan(itemPilihan, &wg, ch)
	}

	// Menunggu semua goroutine selesai
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Menghitung total harga dan menyimpan semua pesanan
	var totalHarga float64
	for pesanan := range ch {
		totalHarga += pesanan.GetHarga() * float64(pesanan.GetQty())
		semuaPesanan = append(semuaPesanan, pesanan)
	}

	// Menampilkan daftar semua pesanan
	fmt.Println("\n=== Daftar Semua Pesanan ===")
	for _, pesanan := range semuaPesanan {
		fmt.Printf("%d x %s @ Rp%.2f = Rp%.2f\n", pesanan.GetQty(), pesanan.GetNama(), pesanan.GetHarga(), pesanan.GetHarga()*float64(pesanan.GetQty()))
	}

	// Melakukan Encoded semua pesanan
	var pesananStr string // Menyimpan string dari semua pesanan
	pesananStr += "Pesanan:["

	for _, pesanan := range semuaPesanan {
		pesananStr += fmt.Sprintf("{%s %d %d} ", pesanan.GetNama(), int(pesanan.GetHarga()), pesanan.GetQty())
	}
	pesananStr = strings.TrimSpace(pesananStr) // Hapus spasi berlebih di akhir
	pesananStr += fmt.Sprintf("], Total: Rp%.2f", totalHarga) // Tambahkan total harga

	encodedSemuaPesanan := encodedSemuaPesanan(pesananStr)
	fmt.Printf("Encoded: %s\n", encodedSemuaPesanan)
	fmt.Printf("\n==============================================\n")
	fmt.Printf("\nTotal yang harus dibayar: Rp%.2f\n", totalHarga)
	fmt.Printf("\n==============================================\n")
	// Meminta pengguna memasukkan jumlah pembayaran
	for {
		fmt.Println("Masukkan jumlah uang yang dibayarkan:")
		scanner.Scan()
		nilaiPembayaran := scanner.Text()

		numPembayaran, err := strconv.ParseFloat(nilaiPembayaran, 64)
		if err != nil || numPembayaran <= 0 {
			fmt.Println("Input tidak valid. Silakan coba lagi.")
			continue
		}

		// Cek apakah uang cukup
		if numPembayaran < totalHarga {
			fmt.Printf("Uang yang dibayarkan kurang. Anda harus membayar minimal Rp%.2f\n", totalHarga)
			continue
		}

		// Hitung kembalian
		kembalian := numPembayaran - totalHarga
		fmt.Printf("Pembayaran diterima. Kembalian: Rp%.2f\n", kembalian)
		break
	}
}

func tambahItemMenu() {
	scanner := bufio.NewScanner(os.Stdin)

	// Tambahkan beberapa item menu awal dengan validasi input harga dan nama
	for {
		fmt.Println("Masukkan nama item menu (atau 'q' untuk selesai): ")
		scanner.Scan()
		nama := scanner.Text()
		if nama == "q" {
			break
		}

		// Validasi nama item
		if err := validasiNamaItem(nama); err != nil {
			fmt.Println(err.Error())
			continue
		}

		// Gunakan fungsi prosesData untuk memproses nama
		prosesData(nama)

		var harga float64
		var masukanHarga string
		var err error

		// Input harga dengan validasi menggunakan regexp dan strconv
		for {
			fmt.Printf("Masukkan harga untuk %s: \n", nama)
			scanner.Scan()
			masukanHarga = scanner.Text()

			harga, err = konversiHarga(masukanHarga)
			if err != nil {
				// Kesalahan sudah ditangani oleh panic dan recover
				fmt.Println(err)
			}else{
				break
			}
		}

		menuItem := &MenuItem{Nama: nama, Harga: harga}
		// Menambahkan item ke menu
		menu = append(menu, menuItem)
	}
}


func hapusItemMenu() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Masukkan nomor menu yang ingin dihapus:")
		scanner.Scan()
		input := scanner.Text()

		menuIndex, err := strconv.Atoi(input)
		if err != nil || menuIndex < 1 || menuIndex > len(menu) {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		} else {
			index := menuIndex - 1
			menu = append(menu[:index], menu[index+1:]...)
			break
		}
	}
	fmt.Println("Item berhasil dihapus.")
}
