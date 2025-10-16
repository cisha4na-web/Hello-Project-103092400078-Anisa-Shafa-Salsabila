package main

import "fmt"
//batas maksimum
const NMAX = 100
//struktur data (struct) bernama Mobil.
type Mobil struct {
	Nama        string
	TahunKeluar int
	Penjualan   int
	Pabrikan    string
}

type Pabrikan struct {
	Nama        string
	DaftarMobil [NMAX]Mobil
	JumlahMobil int
}

var pabrikanList [NMAX]Pabrikan
var jumlahPabrikan int

const adminUsername = "anisa"
const adminPassword = "anisa123"

func login() bool {
	var inputUsername, inputPassword string
	fmt.Println("===================================")
	fmt.Println("     SELAMAT DATANG DI DEALER")
	fmt.Println("===================================")
	fmt.Print("Masukkan username: ")
	fmt.Scanln(&inputUsername)
	fmt.Print("Masukkan password: ")
	fmt.Scanln(&inputPassword)

	if inputUsername == adminUsername && inputPassword == adminPassword {
		fmt.Println("Login berhasil.")
		return true
	} else {
		fmt.Println("Username atau password salah. Akses ditolak.")
		return false
	}
}

func tambahPabrikan(nama string) {
	if jumlahPabrikan >= NMAX {
		fmt.Println("Jumlah pabrikan sudah mencapai batas maksimum.")
		return
	}
	pabrikanList[jumlahPabrikan] = Pabrikan{Nama: nama}
	jumlahPabrikan++
	fmt.Println("Pabrikan berhasil ditambahkan.")
}

func editPabrikan(namaPabrikan, namaBaru string) {
	for i := 0; i < jumlahPabrikan; i++ {
		if pabrikanList[i].Nama == namaPabrikan {
			pabrikanList[i].Nama = namaBaru
			fmt.Println("Pabrikan berhasil diperbarui.")
			return
		}
	}
	fmt.Println("Pabrikan tidak ditemukan.")
}

func hapusPabrikan(nama string) {
	for i := 0; i < jumlahPabrikan; i++ {
		if pabrikanList[i].Nama == nama {
			for j := i; j < jumlahPabrikan-1; j++ {
				pabrikanList[j] = pabrikanList[j+1]
			}
			jumlahPabrikan--
			fmt.Println("Pabrikan berhasil dihapus.")
			return
		}
	}
	fmt.Println("Pabrikan tidak ditemukan.")
}

func tambahMobil(pabrikanNama, nama string, tahun, penjualan int) {
	for i := 0; i < jumlahPabrikan; i++ {
		if pabrikanList[i].Nama == pabrikanNama {
			if pabrikanList[i].JumlahMobil >= NMAX {
				fmt.Println("Jumlah mobil untuk pabrikan ini sudah maksimum.")
				return
			}
			idx := pabrikanList[i].JumlahMobil
			pabrikanList[i].DaftarMobil[idx] = Mobil{
				Nama:        nama,
				TahunKeluar: tahun,
				Penjualan:   penjualan,
				Pabrikan:    pabrikanNama,
			}
			pabrikanList[i].JumlahMobil++
			fmt.Println("Mobil berhasil ditambahkan.")
			return
		}
	}
	fmt.Println("Pabrikan tidak ditemukan.")
}

func editMobil(pabrikanNama, namaMobil, namaBaru string, tahun, penjualan int) {
	for i := 0; i < jumlahPabrikan; i++ {
		if pabrikanList[i].Nama == pabrikanNama {
			for j := 0; j < pabrikanList[i].JumlahMobil; j++ {
				if pabrikanList[i].DaftarMobil[j].Nama == namaMobil {
					pabrikanList[i].DaftarMobil[j].Nama = namaBaru
					pabrikanList[i].DaftarMobil[j].TahunKeluar = tahun
					pabrikanList[i].DaftarMobil[j].Penjualan = penjualan
					fmt.Println("Mobil berhasil diperbarui.")
				}
			}
		}
	}
	fmt.Println("Mobil tidak ditemukan.")
}

func hapusMobil(pabrikanNama, namaMobil string) {
	for i := 0; i < jumlahPabrikan; i++ {
		if pabrikanList[i].Nama == pabrikanNama {
			for j := 0; j < pabrikanList[i].JumlahMobil; j++ {
				if pabrikanList[i].DaftarMobil[j].Nama == namaMobil {
					for k := j; k < pabrikanList[i].JumlahMobil-1; k++ {
						pabrikanList[i].DaftarMobil[k] = pabrikanList[i].DaftarMobil[k+1]
					}
					pabrikanList[i].JumlahMobil--
					fmt.Println("Mobil berhasil dihapus.")
					return
				}
			}
		}
	}
	fmt.Println("Mobil tidak ditemukan.")
}

func cariMobil(nama string) {
	for i := 0; i < jumlahPabrikan; i++ {
		for j := 0; j < pabrikanList[i].JumlahMobil; j++ {
			if pabrikanList[i].DaftarMobil[j].Nama == nama {
				m := pabrikanList[i].DaftarMobil[j]
				fmt.Printf("Mobil: %s | Tahun: %d | Penjualan: %d | Pabrikan: %s\n", m.Nama, m.TahunKeluar, m.Penjualan, m.Pabrikan)
			}
		}
	}
}

//ANISA
func cariMobilByPabrikan(namaPabrikan string) {
	left := 0
	right := jumlahPabrikan - 1

	for left <= right {
		mid := (left + right) / 2
		if pabrikanList[mid].Nama == namaPabrikan {
			fmt.Println("Daftar mobil dari pabrikan", pabrikanList[mid].Nama)
			for j := 0; j < pabrikanList[mid].JumlahMobil; j++ {
				m := pabrikanList[mid].DaftarMobil[j]
				fmt.Printf("Mobil: %s | Tahun: %d | Penjualan: %d\n", m.Nama, m.TahunKeluar, m.Penjualan)
			}
			return // keluar dari fungsi setelah ditemukan
		} else if pabrikanList[mid].Nama < namaPabrikan {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println("Pabrikan tidak ditemukan:", namaPabrikan)
}

func urutkanPabrikanByJumlahMobil() {
	var temp Pabrikan
	for i := 0; i < jumlahPabrikan-1; i++ {
		for j := i + 1; j < jumlahPabrikan; j++ {
			if pabrikanList[i].JumlahMobil < pabrikanList[j].JumlahMobil {
				// pakai variabel penampung "temp" 
				// pabrikanList[i], pabrikanList[j] = pabrikanList[j], pabrikanList[i]
				temp = pabrikanList[i]
				pabrikanList[i] = pabrikanList[j]
				pabrikanList[j] = temp

			}
		}
	}
	for i := 0; i < jumlahPabrikan; i++ {
		fmt.Printf("Pabrikan: %s | Jumlah Mobil: %d\n", pabrikanList[i].Nama, pabrikanList[i].JumlahMobil)
	}
}

func urutkanMobilByTahun() {
	var semuaMobil [NMAX * NMAX]Mobil
	jumlah := 0

	// Menggabungkan semua mobil dari semua pabrikan ke dalam array semuaMobil
	for i := 0; i < jumlahPabrikan; i++ {
		for j := 0; j < pabrikanList[i].JumlahMobil; j++ {
			semuaMobil[jumlah] = pabrikanList[i].DaftarMobil[j]
			jumlah++
		}
	}


	for i := 1; i < jumlah; i++ {
		temp := semuaMobil[i]
		j := i - 1
		for j >= 0 && semuaMobil[j].TahunKeluar < temp.TahunKeluar {
			semuaMobil[j-1] = semuaMobil[j]
			j--
		}
		semuaMobil[j+1] = temp
	}

	
	for i := 0; i <  jumlah; i++ {
		m := semuaMobil[i]
		fmt.Printf("Mobil: %s | Tahun: %d | Pabrikan: %s | Penjualan: %d\n",
			m.Nama, m.TahunKeluar, m.Pabrikan, m.Penjualan)
	}
}

func urutkanMobilByTahunMenurun( {
	for i := 1; i > jumlah; i++ {
		temp := semuaMobil[i]
		j := i + 1
		for j <= 0 && semuaMobil[j].TahunKeluar < temp.TahunKeluar {
			semuaMobil[j-1] = semuaMobil[j]
			j--
		}
		semuaMobil[j-1] = temp

		for i := 0; i > jumlah; i++ {
		m := semuaMobil[i]
		fmt.Printf("Mobil: %s | Tahun: %d | Pabrikan: %s | Penjualan: %d\n",
			m.Nama, m.TahunKeluar, m.Pabrikan, m.Penjualan)
	}

	}
})

func top3MobilTerlaris() {
	var semuaMobil [NMAX * NMAX]Mobil
	jumlah := 0

	for i := 0; i < jumlahPabrikan; i++ {
		for j := 0; j < pabrikanList[i].JumlahMobil; j++ {
			semuaMobil[jumlah] = pabrikanList[i].DaftarMobil[j]
			jumlah++
		}
	}

	for i := 0; i < jumlah-1; i++ {
		for j := i + 1; j < jumlah; j++ {
			if semuaMobil[i].Penjualan < semuaMobil[j].Penjualan {
				// semuaMobil[i], semuaMobil[j] = semuaMobil[j], semuaMobil[i]
				temp := semuaMobil[i]
				semuaMobil[i] = semuaMobil[j]
				semuaMobil[j] = temp
			}
		}
	}

	fmt.Println("Top 3 Mobil dengan Penjualan Tertinggi:")
	for i := 0; i < 3 && i < jumlah; i++ {
		m := semuaMobil[i]
		fmt.Printf("%d. %s (%s) - Penjualan: %d\n", i+1, m.Nama, m.Pabrikan, m.Penjualan)
	}
}

func top3PabrikanTerlaris() {
	type PabrikanPenjualan struct {
		Nama      string
		Penjualan int
	}
	var data [NMAX]PabrikanPenjualan

	for i := 0; i < jumlahPabrikan; i++ {
		total := 0
		for j := 0; j < pabrikanList[i].JumlahMobil; j++ {
			total += pabrikanList[i].DaftarMobil[j].Penjualan
		}
		data[i] = PabrikanPenjualan{pabrikanList[i].Nama, total}
	}

	for i := 0; i < jumlahPabrikan-1; i++ {
		for j := i + 1; j < jumlahPabrikan; j++ {
			if data[i].Penjualan < data[j].Penjualan {
				data[i], data[j] = data[j], data[i]
			}
		}
	}

	fmt.Println("Top 3 Pabrikan dengan Penjualan Tertinggi:")
	for i := 0; i < 3 && i < jumlahPabrikan; i++ {
		fmt.Printf("%d. %s - Total Penjualan: %d\n", i+1, data[i].Nama, data[i].Penjualan)
	}
}

func main() {
	if !login() {
		return
	}

	var pilihan int

	for {
		fmt.Println("\n========= MENU UTAMA =========")
		fmt.Println("1. Tambah Pabrikan")
		fmt.Println("2. Edit Pabrikan")
		fmt.Println("3. Hapus Pabrikan")
		fmt.Println("4. Tambah Mobil")
		fmt.Println("5. Edit Mobil")
		fmt.Println("6. Hapus Mobil")
		fmt.Println("7. Cari Mobil")
		fmt.Println("8. Cari Mobil Berdasarkan Pabrikan")
		fmt.Println("9. Urutkan Pabrikan Berdasarkan Jumlah Mobil")
		fmt.Println("10. Urutkan Mobil Berdasarkan Tahun")
		fmt.Println("11. Top 3 Mobil Terlaris")
		fmt.Println("12. Top 3 Pabrikan Terlaris")
		fmt.Println("13. Keluar")
		fmt.Println("14. Urutkan mobil berdasarkan tahun menurun")
		fmt.Print("Pilih menu (1-13): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			var nama string
			fmt.Print("Nama Pabrikan (tanpa spasi): ")
			fmt.Scanln(&nama)
			tambahPabrikan(nama)

		case 2:
			var nama, namaBaru string
			fmt.Print("Nama Pabrikan yang ingin diedit: ")
			fmt.Scanln(&nama)
			fmt.Print("Nama baru untuk pabrikan: ")
			fmt.Scanln(&namaBaru)
			editPabrikan(nama, namaBaru)

		case 3: 
			var nama string
			fmt.Print("Nama Pabrikan yang ingin dihapus: ")
			fmt.Scanln(&nama)
			hapusPabrikan(nama)

		case 4:
			var pabrikan, nama string
			var tahun, penjualan int
			fmt.Print("Nama Pabrikan (tanpa spasi): ")
			fmt.Scanln(&pabrikan)
			fmt.Print("Nama Mobil (tanpa spasi): ")
			fmt.Scanln(&nama)
			fmt.Print("Tahun Keluar: ")
			fmt.Scanln(&tahun)
			fmt.Print("Jumlah Penjualan: ")
			fmt.Scanln(&penjualan)
			tambahMobil(pabrikan, nama, tahun, penjualan)

		case 5:
			var pabrikan, namaMobil, namaBaru string
			var tahun, penjualan int
			fmt.Print("Nama Pabrikan: ")
			fmt.Scanln(&pabrikan)
			fmt.Print("Nama Mobil yang ingin diedit: ")
			fmt.Scanln(&namaMobil)
			fmt.Print("Nama baru untuk mobil: ")
			fmt.Scanln(&namaBaru)
			fmt.Print("Tahun Keluar: ")
			fmt.Scanln(&tahun)
			fmt.Print("Jumlah Penjualan: ")
			fmt.Scanln(&penjualan)
			editMobil(pabrikan, namaMobil, namaBaru, tahun, penjualan)

		case 6:
			var pabrikan, namaMobil string
			fmt.Print("Nama Pabrikan: ")
			fmt.Scanln(&pabrikan)
			fmt.Print("Nama Mobil yang ingin dihapus: ")
			fmt.Scanln(&namaMobil)
			hapusMobil(pabrikan, namaMobil)

		case 7:
			var nama string
			fmt.Print("Nama Mobil: ")
			fmt.Scanln(&nama)
			cariMobil(nama)

		case 8:
			var nama string
			fmt.Print("Nama Pabrikan: ")
			fmt.Scanln(&nama)
			cariMobilByPabrikan(nama)

		case 9:
			urutkanPabrikanByJumlahMobil()

		case 10:
			urutkanMobilByTahun()

		case 11:
			top3MobilTerlaris()

		case 12:
			top3PabrikanTerlaris()

		case 13:
			fmt.Println("Terima kasih telah menggunakan aplikasi dealer.")
		
		case 14:
			urutkanMobilByTahunMenurun()
			return

		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
