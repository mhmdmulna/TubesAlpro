package main

import (
	"fmt"
)

const NMAX int = 1000

type sampah struct {
	jenis        string
	jumlah       int
	berat, total float64
	daurUlang    bool
}

type dataSampah [NMAX]sampah

func main() {
	var data dataSampah
	var jumlahData int
	var pilihan int

	for {
		fmt.Println("\n----- Aplikasi Pengelolaan Sampah -----")
		fmt.Println("1. Tambah Data Sampah")
		fmt.Println("2. Ubah Data Sampah")
		fmt.Println("3. Hapus Data Sampah")
		fmt.Println("4. Cari Data Sampah")
		fmt.Println("5. Urutkan Data Sampah")
		fmt.Println("6. Tampilkan Semua Data")
		fmt.Println("0. Keluar\n")

		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahData(&data, &jumlahData)
		case 2:
			ubahData(&data, jumlahData)
		case 3:
			hapusData(&data, &jumlahData)
		case 4:
			var key string
			fmt.Print("Masukkan jenis sampah yang dicari: ")
			fmt.Scan(&key)
			index := cariData(data, jumlahData, key)
			if index != -1 {
				fmt.Printf("Data ditemukan: Jenis: %s, Jumlah: %d, Berat: %.2f, Total: %.2f\n",
					data[index].jenis, data[index].jumlah, data[index].berat, data[index].total)
			} else {
				fmt.Println("Data tidak ditemukan.")
			}

		case 5:
			urutkanData(&data, jumlahData)
		case 6:
			tampilkanData(data, jumlahData)
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
func tampilkanData(data dataSampah, n int) {
	if n == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	fmt.Printf("\n%-4s %-15s %-10s %-10s %-10s\n", "No", "Jenis", "Jumlah", "Berat", "Total")
	for i := 0; i < n; i++ {
		fmt.Printf("%-4d %-15s %-10d %-10.2f %-10.2f\n", i+1, data[i].jenis, data[i].jumlah, data[i].berat, data[i].total)
	}
}

func tambahData(A *dataSampah, n *int) {
	if *n >= NMAX {
		fmt.Println("Kapasitas penyimpanan data sampah penuh.")
		return
	}

	fmt.Print("Masukkan jenis sampah: ")
	fmt.Scan(&A[*n].jenis)
	fmt.Print("Masukkan jumlah sampah: ")
	fmt.Scan(&A[*n].jumlah)
	fmt.Print("Masukkan berat per item (kg): ")
	fmt.Scan(&A[*n].berat)

	A[*n].total = float64(A[*n].jumlah) * A[*n].berat
	fmt.Println("Data sampah berhasil ditambahkan.")
	*n = *n + 1
	fmt.Println("Jumlah data saat ini:", *n)
}

func ubahData(data *dataSampah, n int) {
	if n == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	var key string
	fmt.Print("Masukkan jenis sampah yang ingin diubah: ")
	fmt.Scan(&key)

	var hasil [NMAX]int
	var jumlahHasil int

	// Cari semua indeks data yang jenisnya sama dengan key
	cariSemuaJenis(*data, n, key, &hasil, &jumlahHasil)

	if jumlahHasil == 0 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	fmt.Println("Data yang ditemukan:")
	for i := 0; i < jumlahHasil; i++ {
		idx := hasil[i]
		fmt.Printf("%d. Jenis: %s, Jumlah: %d, Berat: %.2f, Total: %.2f\n",
			i+1, data[idx].jenis, data[idx].jumlah, data[idx].berat, data[idx].total)
	}

	var pilihan int
	fmt.Print("Pilih nomor data yang ingin diubah: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > jumlahHasil {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	idxUbah := hasil[pilihan-1]

	var jenisBaru string
	var jumlahBaru int
	var beratBaru float64

	fmt.Print("Masukkan jenis sampah baru: ")
	fmt.Scan(&jenisBaru)
	fmt.Print("Masukkan jumlah baru: ")
	fmt.Scan(&jumlahBaru)
	fmt.Print("Masukkan berat baru (kg): ")
	fmt.Scan(&beratBaru)

	data[idxUbah].jenis = jenisBaru
	data[idxUbah].jumlah = jumlahBaru
	data[idxUbah].berat = beratBaru
	data[idxUbah].total = float64(jumlahBaru) * beratBaru

	fmt.Println("Data sampah berhasil diubah.")
}

func hapusData(data *dataSampah, n *int) {
	tampilkanData(*data, *n)
	if *n == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	var key string
	fmt.Print("Masukkan jenis sampah yang ingin dihapus: ")
	fmt.Scan(&key)

	var hasil [NMAX]int
	var jumlahHasil int

	cariSemuaJenis(*data, *n, key, &hasil, &jumlahHasil)

	if jumlahHasil == 0 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	fmt.Println("Data yang ditemukan:")
	for i := 0; i < jumlahHasil; i++ {
		idx := hasil[i]
		fmt.Printf("%d. Jenis: %s, Jumlah: %d, Berat: %.2f, Total: %.2f\n",
			i+1, data[idx].jenis, data[idx].jumlah, data[idx].berat, data[idx].total)
	}

	var pilihan int
	fmt.Print("Pilih nomor data yang ingin dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > jumlahHasil {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	idxHapus := hasil[pilihan-1]

	// Geser data ke kiri
	for i := idxHapus; i < *n-1; i++ {
		data[i] = data[i+1]
	}
	*n--
	fmt.Println("Data berhasil dihapus.")
}

func cariData(data dataSampah, n int, key string) int {
	if n == 0 {
		fmt.Println("Data kosong.")
		return -1
	}

	var metode int
	fmt.Print("Pilih metode pencarian (1=Sequential, 2=Binary): ")
	fmt.Scan(&metode)

	if metode == 2 {
		urutkanData(&data, n) // pastikan data terurut dulu
		return binarySearch(data, n, key)
	} else {
		return sequentialSearch(data, n, key)
	}
}

func sequentialSearch(data dataSampah, n int, key string) int {
	for i := 0; i < n; i++ {
		if data[i].jenis == key {
			return i
		}
	}
	return -1
}

func binarySearch(data dataSampah, n int, key string) int {
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		if data[mid].jenis == key {
			return mid
		} else if data[mid].jenis < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func urutkanData(data *dataSampah, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if data[i].jenis > data[j].jenis {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	fmt.Println("Data berhasil diurutkan berdasarkan jenis.")
	tampilkanData(*data, n)
}
func cariSemuaJenis(data dataSampah, n int, key string, hasil *[NMAX]int, jumlah *int) {
	*jumlah = 0
	for i := 0; i < n; i++ {
		if data[i].jenis == key {
			hasil[*jumlah] = i
			*jumlah++
		}
	}
}
