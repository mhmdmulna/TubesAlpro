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
		fmt.Println("7. Tampilkan Statistik")
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
			menuCariData(data, jumlahData)
		case 5:
			menuUrutkanData(&data, jumlahData)
		case 6:
			tampilkanData(data, jumlahData)
		case 7:
			tampilkanStatistik(data, jumlahData)
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
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
	fmt.Print("Apakah sampah ini didaur ulang? (ya/tidak): ")
	var daur string
	fmt.Scan(&daur)
	A[*n].daurUlang = (daur == "ya")

	A[*n].total = float64(A[*n].jumlah) * A[*n].berat
	*n++
	fmt.Println("Data sampah berhasil ditambahkan.")
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
	cariSemuaJenis(*data, n, key, &hasil, &jumlahHasil)

	if jumlahHasil == 0 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	fmt.Println("Data yang ditemukan:")
	for i := 0; i < jumlahHasil; i++ {
		idx := hasil[i]
		fmt.Printf("%d. Jenis: %s, Jumlah: %d, Berat: %.2f, Total: %.2f, Daur Ulang: %t\n",
			i+1, data[idx].jenis, data[idx].jumlah, data[idx].berat, data[idx].total, data[idx].daurUlang)
	}

	var pilihan int
	fmt.Print("Pilih nomor data yang ingin diubah: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > jumlahHasil {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	idxUbah := hasil[pilihan-1]

	// Masukkan data baru
	var jenisBaru string
	var jumlahBaru int
	var beratBaru float64
	var daurUlangInput string

	fmt.Print("Masukkan jenis sampah baru: ")
	fmt.Scan(&jenisBaru)

	fmt.Print("Masukkan jumlah baru: ")
	fmt.Scan(&jumlahBaru)
	if jumlahBaru < 0 {
		fmt.Println("Jumlah tidak boleh negatif.")
		return
	}

	fmt.Print("Masukkan berat baru (kg): ")
	fmt.Scan(&beratBaru)
	if beratBaru < 0 {
		fmt.Println("Berat tidak boleh negatif.")
		return
	}

	fmt.Print("Apakah sampah ini didaur ulang? (ya/tidak): ")
	fmt.Scan(&daurUlangInput)

	data[idxUbah].jenis = jenisBaru
	data[idxUbah].jumlah = jumlahBaru
	data[idxUbah].berat = beratBaru
	data[idxUbah].total = float64(jumlahBaru) * beratBaru
	data[idxUbah].daurUlang = daurUlangInput == "ya"

	fmt.Println("Data sampah berhasil diubah.")
}

func hapusData(data *dataSampah, n *int) {
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
		fmt.Printf("%d. Jenis: %s, Jumlah: %d, Berat: %.2f, Total: %.2f, Daur Ulang: %t\n",
			i+1, data[idx].jenis, data[idx].jumlah, data[idx].berat, data[idx].total, data[idx].daurUlang)
	}

	var pilihan int
	fmt.Print("Pilih nomor data yang ingin dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > jumlahHasil {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	idxHapus := hasil[pilihan-1]

	// Geser elemen ke kiri untuk menghapus data
	for i := idxHapus; i < *n-1; i++ {
		data[i] = data[i+1]
	}
	*n--
	fmt.Println("Data berhasil dihapus.")
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

func tampilkanData(data dataSampah, n int) {
	if n == 0 {
		fmt.Println("Belum ada data.")
		return
	}
	fmt.Printf("\n%-4s %-15s %-10s %-10s %-10s %-10s\n", "No", "Jenis", "Jumlah", "Berat", "Total", "DaurUlang")
	for i := 0; i < n; i++ {
		daur := "Tidak"
		if data[i].daurUlang {
			daur = "Ya"
		}
		fmt.Printf("%-4d %-15s %-10d %-10.2f %-10.2f %-10s\n", i+1, data[i].jenis, data[i].jumlah, data[i].berat, data[i].total, daur)
	}
}

func tampilkanStatistik(data dataSampah, n int) {
	totalJumlah, totalDaurUlang := 0, 0
	for i := 0; i < n; i++ {
		totalJumlah += data[i].jumlah
		if data[i].daurUlang {
			totalDaurUlang += data[i].jumlah
		}
	}
	fmt.Println("\n--- Statistik ---")
	fmt.Println("Total sampah:", totalJumlah)
	fmt.Println("Total sampah yang didaur ulang:", totalDaurUlang)
}

func menuCariData(data dataSampah, n int) {
	var key string
	fmt.Print("Masukkan jenis sampah yang dicari: ")
	fmt.Scan(&key)
	fmt.Print("Pilih metode pencarian (1 = Sequential, 2 = Binary): ")
	var metode int
	fmt.Scan(&metode)

	var index int
	if metode == 1 {
		index = sequentialSearch(data, n, key)
	} else {
		urutkanDataByJenis(&data, n)
		index = binarySearch(data, n, key)
	}

	if index != -1 {
		fmt.Println("Data ditemukan:")
		fmt.Printf("Jenis: %s, Jumlah: %d, Berat: %.2f, Total: %.2f\n",
			data[index].jenis, data[index].jumlah, data[index].berat, data[index].total)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func menuUrutkanData(data *dataSampah, n int) {
	fmt.Print("Urut berdasarkan: 1 = Jenis (Selection Sort), 2 = Jumlah (Insertion Sort): ")
	var pilihan int
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		urutkanDataByJenis(data, n)
		fmt.Println("Data diurutkan berdasarkan jenis.")
	} else {
		urutkanDataByJumlah(data, n)
		fmt.Println("Data diurutkan berdasarkan jumlah.")
	}
	tampilkanData(*data, n)
}

func urutkanDataByJenis(data *dataSampah, n int) {
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if data[j].jenis < data[min].jenis {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
}

func urutkanDataByJumlah(data *dataSampah, n int) {
	for i := 1; i < n; i++ {
		temp := data[i]
		j := i - 1
		for j >= 0 && data[j].jumlah > temp.jumlah {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
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
	low, high := 0, n-1
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
