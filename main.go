package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Komentar struct {
	ID       int
	Teks     string
	Sentimen string
}

// Menggunakan variabel global 
var dataKomentar []Komentar
var idSekarang = 1

func cekSentimen(teks string) string {
	kataPositif := []string{"bagus", "keren", "mantap", "baik"}
	kataNegatif := []string{"jelek", "buruk", "parah", "benci"}

	// Ubah teks ke huruf kecil dan pecah per kata
	teks = strings.ToLower(teks)
	kataKata := strings.Split(teks, " ")
	
	skor := 0

	// Pencocokan kata utuh 
	for _, kata := range kataKata {
		// Cek daftar positif
		for _, pos := range kataPositif {
			if kata == pos {
				skor++
			}
		}
		// Cek daftar negatif
		for _, neg := range kataNegatif {
			if kata == neg {
				skor--
			}
		}
	}

	if skor > 0 {
		return "positif"
	} else if skor < 0 {
		return "negatif"
	}
	
	return "netral"
}

func tambahKomentar(teks string) {
	sentimen := cekSentimen(teks)
	komentarBaru := Komentar{
		ID:       idSekarang,
		Teks:     teks,
		Sentimen: sentimen,
	}
	dataKomentar = append(dataKomentar, komentarBaru)
	idSekarang++
}

func updateKomentar(id int, teksBaru string) {
	for i := 0; i < len(dataKomentar); i++ {
		if dataKomentar[i].ID == id {
			dataKomentar[i].Teks = teksBaru
			dataKomentar[i].Sentimen = cekSentimen(teksBaru)
			fmt.Println("Komentar berhasil diubah!")
			return
		}
	}
	fmt.Println("Gagal: ID tidak ditemukan")
}

func hapusKomentar(id int) {
	for i := 0; i < len(dataKomentar); i++ {
		if dataKomentar[i].ID == id {
			// Cara memotong slice yang biasa diajarkan
			dataKomentar = append(dataKomentar[:i], dataKomentar[i+1:]...)
			fmt.Println("Komentar berhasil dihapus!")
			return
		}
	}
	fmt.Println("Gagal: ID tidak ditemukan")
}

// Sequential Search
func cariSequential(keyword string) {
	ketemu := false
	keyword = strings.ToLower(keyword)
	
	for _, k := range dataKomentar {
		if strings.Contains(strings.ToLower(k.Teks), keyword) {
			fmt.Printf("ID: %d | Sentimen: %-7s | Teks: %s\n", k.ID, k.Sentimen, k.Teks)
			ketemu = true
		}
	}
	
	if !ketemu {
		fmt.Println("Komentar dengan kata kunci tersebut tidak ditemukan.")
	}
}

// Binary Search
func cariBinary(target string) {
	left := 0
	right := len(dataKomentar) - 1
	target = strings.ToLower(target)

	for left <= right {
		mid := (left + right) / 2
		teksMid := strings.ToLower(dataKomentar[mid].Teks)

		if teksMid == target {
			fmt.Printf("Ditemukan -> ID: %d | Sentimen: %-7s | Teks: %s\n", dataKomentar[mid].ID, dataKomentar[mid].Sentimen, dataKomentar[mid].Teks)
			return
		} else if teksMid < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("Komentar tidak ditemukan (Catatan: Binary Search butuh data yang disortir secara alfabet).")
}

// Selection Sort 
func urutkanPanjangTeks() {
	n := len(dataKomentar)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if len(dataKomentar[j].Teks) < len(dataKomentar[minIndex].Teks) {
				minIndex = j
			}
		}
		// Tukar posisi nya lek
		temp := dataKomentar[i]
		dataKomentar[i] = dataKomentar[minIndex]
		dataKomentar[minIndex] = temp
	}
	fmt.Println("Berhasil disortir dari teks terpendek ke terpanjang.")
}

// Fungsi untuk Insertion Sort
func angkaSentimen(s string) int {
	if s == "positif" {
		return 1
	} else if s == "netral" {
		return 2
	}
	return 3 // negatif
}

// Insertion Sort
func urutkanSentimen() {
	n := len(dataKomentar)
	for i := 1; i < n; i++ {
		key := dataKomentar[i]
		j := i - 1

		for j >= 0 && angkaSentimen(dataKomentar[j].Sentimen) > angkaSentimen(key.Sentimen) {
			dataKomentar[j+1] = dataKomentar[j]
			j--
		}
		dataKomentar[j+1] = key
	}
	fmt.Println("Berhasil disortir (Positif -> Netral -> Negatif).")
}

func tampilkanStatistik() {
	pos, net, neg := 0, 0, 0
	for _, k := range dataKomentar {
		if k.Sentimen == "positif" {
			pos++
		} else if k.Sentimen == "netral" {
			net++
		} else if k.Sentimen == "negatif" {
			neg++
		}
	}

	fmt.Println("\n=== STATISTIK SENTIMEN ===")
	fmt.Println("Positif :", pos)
	fmt.Println("Netral  :", net)
	fmt.Println("Negatif :", neg)
}

func tampilkanSemua() {
	fmt.Println("\n=== DAFTAR KOMENTAR ===")
	if len(dataKomentar) == 0 {
		fmt.Println("Data masih kosong.")
		return
	}
	for _, k := range dataKomentar {
		fmt.Printf("ID: %d | Sentimen: %-7s | Teks: %s\n", k.ID, k.Sentimen, k.Teks)
	}
}

func main() {
	// Menggunakan Scanner agar jauh lebih mudah membaca spasi dibanding Reader
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== MENU APLIKASI ===")
		fmt.Println("1. Tambah Komentar")
		fmt.Println("2. Update Komentar")
		fmt.Println("3. Hapus Komentar")
		fmt.Println("4. Tampilkan Semua")
		fmt.Println("5. Search (Sequential - Cari per Kata)")
		fmt.Println("6. Sorting (Panjang Teks - Selection Sort)")
		fmt.Println("7. Sorting (Sentimen - Insertion Sort)")
		fmt.Println("8. Tampilkan Statistik")
		fmt.Println("9. Binary Search (Pencarian Teks Sama Persis)")
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
			fmt.Println("Komentar ditambahkan!")

		case "2":
			fmt.Print("Masukkan ID yang mau diubah: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Komentar Baru: ")
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

		case "5":
			fmt.Print("Masukkan Kata Kunci: ")
			scanner.Scan()
			keyword := scanner.Text()
			
			cariSequential(keyword)

		case "6":
			urutkanPanjangTeks()
			tampilkanSemua()

		case "7":
			urutkanSentimen()
			tampilkanSemua()

		case "8":
			tampilkanStatistik()

		case "9":
			fmt.Print("Masukkan teks *persis/utuh* yang mau dicari: ")
			scanner.Scan()
			target := scanner.Text()
			
			cariBinary(target)

		case "0":
			fmt.Println("Program selesai. Sampai jumpa!")
			return

		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}