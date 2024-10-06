package main

import (
	"fmt"
	"soal_golang_2/view"
)

func main() {
	// Penggunaan defer
	defer func() {
		fmt.Println("Program selesai")
	}()
	// Fungsi Exported 
	view.TampilkanMenuUtama()
}