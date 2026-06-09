package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Komentar struct {
	ID   int
	Teks string
}

var dataKomentar []Komentar
var idSekarang = 1

func tambahKomentar(teks string) {
	komentarBaru := Komentar{
		ID:   idSekarang,
		Teks: teks,
	}

	dataKomentar = append(dataKomentar, komentarBaru)
	idSekarang++
}

func updateKomentar(id int, teksBaru string) {
	for i := 0; i < len(dataKomentar); i++ {
		if dataKomentar[i].ID == id {
			dataKomentar[i].Teks = teksBaru
			fmt.Println("Komentar berhasil diubah!")
			return
		}
	}

	fmt.Println("ID tidak ditemukan!")
}

func hapusKomentar(id int) {
	for i := 0; i < len(dataKomentar); i++ {
		if dataKomentar[i].ID == id {
			dataKomentar = append(dataKomentar[:i], dataKomentar[i+1:]...)
			fmt.Println("Komentar berhasil dihapus!")
			return
		}
	}

	fmt.Println("ID tidak ditemukan!")
}

func tampilkanSemua() {
	fmt.Println("\n=== DAFTAR KOMENTAR ===")

	if len(dataKomentar) == 0 {
		fmt.Println("Data masih kosong.")
		return
	}

	for _, k := range dataKomentar {
		fmt.Printf("ID: %d | Teks: %s\n", k.ID, k.Teks)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== MENU APLIKASI ===")
		fmt.Println("1. Tambah Komentar")
		fmt.Println("2. Update Komentar")
		fmt.Println("3. Hapus Komentar")
		fmt.Println("4. Tampilkan Semua")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih Menu: ")

		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {

		case "1":
			fmt.Print("Masukkan Komentar: ")
			scanner.Scan()
			teks := scanner.Text()

			tambahKomentar(teks)
			fmt.Println("Komentar berhasil ditambahkan!")

		case "2":
			fmt.Print("Masukkan ID yang mau diubah: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Masukkan Komentar Baru: ")
			scanner.Scan()
			teksBaru := scanner.Text()

			updateKomentar(id, teksBaru)

		case "3":
			fmt.Print("Masukkan ID yang mau dihapus: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			hapusKomentar(id)

		case "4":
			tampilkanSemua()

		case "0":
			fmt.Println("Program selesai. Sampai jumpa!")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
