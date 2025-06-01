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
	metodeDaur   string
}

type dataSampah [NMAX]sampah

func main() {
	var data dataSampah
	var jumlahData int
	var pilihan int

	// ----- Dummy Data -----
	data[0] = sampah{jenis: "Plastik", jumlah: 10, berat: 0.5, total: 5.0, daurUlang: true, metodeDaur: "Dicacah"}
	data[1] = sampah{jenis: "Kertas", jumlah: 20, berat: 0.1, total: 2.0, daurUlang: true, metodeDaur: "Dibuat bubur"}
	data[2] = sampah{jenis: "Plastik", jumlah: 15, berat: 1.0, total: 15.0, daurUlang: true, metodeDaur: "Dicacah"}
	data[3] = sampah{jenis: "Kardus", jumlah: 60, berat: 0.2, total: 60 * 0.2, daurUlang: true, metodeDaur: "BuburKertas"}
	data[4] = sampah{jenis: "Styrofoam", jumlah: 20, berat: 0.01, total: 20 * 0.01, daurUlang: false, metodeDaur: "-"}
	data[5] = sampah{jenis: "Plastik", jumlah: 50, berat: 1.5, total: 25.0, daurUlang: true, metodeDaur: "Dicacah"}
	data[6] = sampah{jenis: "Styrofoam", jumlah: 42, berat: 0.5, total: 20 * 0.01, daurUlang: false, metodeDaur: "-"}
	jumlahData = 7

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
			fmt.Println("Terima kasih telah menggunakan aplikasi kami.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahData(A *dataSampah, n *int) {
	var daur string

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
	fmt.Scan(&daur)
	A[*n].daurUlang = (daur == "ya")

	if A[*n].daurUlang {
		fmt.Print("Masukkan metode daur ulang: ")
		fmt.Scan(&A[*n].metodeDaur)
	} else {
		A[*n].metodeDaur = "-"
	}

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
		fmt.Printf("%d. Jenis: %s, Jumlah: %d, Berat: %.2f, Total: %.2f, Daur Ulang: %t, Metode Daur Ulang: %s\n", i+1, data[idx].jenis, data[idx].jumlah, data[idx].berat, data[idx].total, data[idx].daurUlang, data[idx].metodeDaur)
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
	var daurUlangInput string
	var metodeDaurUlangBaru string

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
	data[idxUbah].daurUlang = (daurUlangInput == "ya")

	if data[idxUbah].daurUlang {
		fmt.Print("Masukkan metode daur ulang baru: ")
		fmt.Scan(&metodeDaurUlangBaru)
		data[idxUbah].metodeDaur = metodeDaurUlangBaru
	} else {
		data[idxUbah].metodeDaur = "-"
	}

	data[idxUbah].jenis = jenisBaru
	data[idxUbah].jumlah = jumlahBaru
	data[idxUbah].berat = beratBaru
	data[idxUbah].total = float64(jumlahBaru) * beratBaru

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
		fmt.Printf("%d. Jenis: %s, Jumlah: %d, Berat: %.2f, Total: %.2f, Daur Ulang: %t, Metode Daur Ulang: %s\n", i+1, data[idx].jenis, data[idx].jumlah, data[idx].berat, data[idx].total, data[idx].daurUlang, data[idx].metodeDaur)
	}

	var pilihan int
	fmt.Print("Pilih nomor data yang ingin dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > jumlahHasil {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	idxHapus := hasil[pilihan-1]

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
	fmt.Printf("\n%-4s %-15s %-10s %-10s %-10s %-10s %-20s\n", "No", "Jenis", "Jumlah", "Berat", "Total", "DaurUlang", "Metode Daur Ulang")
	for i := 0; i < n; i++ {
		daur := "Tidak"
		if data[i].daurUlang {
			daur = "Ya"
		}
		fmt.Printf("%-4d %-15s %-10d %-10.2f %-10.2f %-10s %-20s\n", i+1, data[i].jenis, data[i].jumlah, data[i].berat, data[i].total, daur, data[i].metodeDaur)
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
	var hasil [NMAX]int
	var jumlahHasil int
	fmt.Print("Masukkan jenis sampah yang dicari: ")
	fmt.Scan(&key)
	fmt.Print("Pilih metode pencarian (1 = Sequential, 2 = Binary): ")
	var metode int
	fmt.Scan(&metode)

	switch metode {
	case 1:
		sequentialSearch(data, n, key, &hasil, &jumlahHasil)

		if jumlahHasil == 0 {
			fmt.Println("Data tidak ditemukan.")
			return
		}

		fmt.Println("Data yang ditemukan:")
		for i := 0; i < jumlahHasil; i++ {
			idx := hasil[i]
			fmt.Printf("%d. Jenis: %s, Jumlah: %d, Berat: %.2f, Total: %.2f, Daur Ulang: %t, Metode Daur Ulang: %s\n", i+1, data[idx].jenis, data[idx].jumlah, data[idx].berat, data[idx].total, data[idx].daurUlang, data[idx].metodeDaur)
		}
	case 2:
		selectionSortByJenis(&data, n, true)
		binarySearch(data, n, key, &hasil, &jumlahHasil)

		if jumlahHasil == 0 {
			fmt.Println("Data tidak ditemukan.")
			return
		}

		fmt.Println("Data yang ditemukan:")
		for i := 0; i < jumlahHasil; i++ {
			idx := hasil[i]
			fmt.Printf("%d. Jenis: %s, Jumlah: %d, Berat: %.2f, Total: %.2f, Daur Ulang: %t, Metode Daur Ulang: %s\n", i+1, data[idx].jenis, data[idx].jumlah, data[idx].berat, data[idx].total, data[idx].daurUlang, data[idx].metodeDaur)
		}
	}
}

func menuUrutkanData(data *dataSampah, n int) {
	var metode, kriteria, arah int

	fmt.Println("\n--- Menu Pengurutan Data ---")
	fmt.Println("Metode:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Pilih metode pengurutan (1/2): ")
	fmt.Scan(&metode)

	fmt.Println("\nKriteria:")
	fmt.Println("1. Jenis")
	fmt.Println("2. Jumlah")
	fmt.Print("Pilih kriteria pengurutan (1/2): ")
	fmt.Scan(&kriteria)

	fmt.Println("\nArah:")
	fmt.Println("1. Menaik (Ascending)")
	fmt.Println("2. Menurun (Descending)")
	fmt.Print("Pilih arah pengurutan (1/2): ")
	fmt.Scan(&arah)

	switch metode {
	case 1:
		if kriteria == 1 {
			selectionSortByJenis(data, n, arah == 1)
		} else if kriteria == 2 {
			selectionSortByJumlah(data, n, arah == 1)
		}
	case 2:
		if kriteria == 1 {
			insertionSortByJenis(data, n, arah == 1)
		} else if kriteria == 2 {
			insertionSortByJumlah(data, n, arah == 1)
		}
	default:
		fmt.Println("Metode tidak valid.")
		return
	}

	fmt.Println("Data berhasil diurutkan.")
	tampilkanData(*data, n)
}

func selectionSortByJenis(data *dataSampah, n int, ascending bool) {
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (ascending && data[j].jenis < data[idx].jenis) || (!ascending && data[j].jenis > data[idx].jenis) {
				idx = j
			}
		}
		data[i], data[idx] = data[idx], data[i]
	}
}

func selectionSortByJumlah(data *dataSampah, n int, ascending bool) {
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (ascending && data[j].jumlah < data[idx].jumlah) || (!ascending && data[j].jumlah > data[idx].jumlah) {
				idx = j
			}
		}
		data[i], data[idx] = data[idx], data[i]
	}
}

func insertionSortByJenis(data *dataSampah, n int, ascending bool) {
	for i := 1; i < n; i++ {
		temp := data[i]
		j := i - 1
		for j >= 0 && ((ascending && data[j].jenis > temp.jenis) || (!ascending && data[j].jenis < temp.jenis)) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func insertionSortByJumlah(data *dataSampah, n int, ascending bool) {
	for i := 1; i < n; i++ {
		temp := data[i]
		j := i - 1
		for j >= 0 && ((ascending && data[j].jumlah > temp.jumlah) || (!ascending && data[j].jumlah < temp.jumlah)) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func sequentialSearch(data dataSampah, n int, key string, hasil *[NMAX]int, jumlah *int) {
	*jumlah = 0
	for i := 0; i < n; i++ {
		if data[i].jenis == key {
			hasil[*jumlah] = i
			*jumlah++
		}
	}
}

func binarySearch(data dataSampah, n int, key string, hasil *[NMAX]int, jumlah *int) {
	*jumlah = 0
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if data[mid].jenis == key {
			hasil[*jumlah] = mid
			*jumlah++
			// Cari ke kiri
			left := mid - 1
			for left >= 0 && data[left].jenis == key {
				hasil[*jumlah] = left
				*jumlah++
				left--
			}
			// Cari ke kanan
			right := mid + 1
			for right < n && data[right].jenis == key {
				hasil[*jumlah] = right
				*jumlah++
				right++
			}
			return
		} else if data[mid].jenis < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
}
