package main

import "fmt"

const NMAX int = 100

type Subscription struct {
	namaLangganan string
	biaya         int
}
type finansialManage [NMAX]Subscription

func main() {
	var arr finansialManage
	var n int
	var perintah string
	cetakUI()
	fmt.Scan(&perintah)
	for {
		switch perintah {
		case "1":
			baca(&arr, &n)
			fmt.Println(n)
			cetakUI()
			fmt.Scan(&perintah)
			// tambah subscription apa saja
		case "2":
			cetakarr(arr, n)
			hapusarray(&arr, &n)
			fmt.Println(n)
			cetakarr(arr, n)
			cetakUI()
			fmt.Scan(&perintah)
			//hapus langganan
		case "3":
			var x string
			cetakarr(arr, n)
			fmt.Println("Pilih Nama Langganan yang ada list: ")
			fmt.Scan(&x)
			renameData(&arr, n, x)
			cetakUI()
			fmt.Scan(&perintah)
		case "4":
			fmt.Println("Detail Total Pengeluaran: ")
			cetakarr(arr, n)
			fmt.Printf("Total Biaya sebesar: Rp. %d\n", totalBiaya(arr, n))
			fmt.Printf("Banyaknya langganan: %d\n", n)
			cetakUI()
			fmt.Scan(&perintah)
			//detail semua langganan
			// detail semua pengeluaran
		case "5":

		case "6":
			return
		default:
			fmt.Println("✖ Pilihan tidak valid.")
		}
	}
}
func renameData(A *finansialManage, n int, x string ){
	var ketemu int
	var perintah string
	ketemu = Caridata(*A, n, x)
	if ketemu == -1 {
		fmt.Print("Data tidak ada")
	}else {
		fmt.Println("Data ditemukan")
		fmt.Println("Apakah ingin dirubah? (ya/tidak) ")
		fmt.Scan(&perintah)
		if perintah == "ya"{
			fmt.Println("Masukan nama langganan dan biayanya!")
			fmt.Scan(&A[ketemu].namaLangganan, &A[ketemu].biaya)
		}
	}

}
func Caridata(A finansialManage, n int, x string)int{
	var ketemu int = -1
	var i int
	for i<n && ketemu == -1{
		if A[i].namaLangganan == x {
			ketemu = i
		}
		i++
	}
	return ketemu
}

func cetakUI() {
	fmt.Println("\n===== Aplikasi Manajemen Subskripsi & Keuangan =====")
	fmt.Println("1. Tambah Langganan")
	fmt.Println("2. Hapus Langganan")
	fmt.Println("3. Cari Data")
	fmt.Println("4. Lihat Total Pengeluaran Bulanan")
	fmt.Println("5. Keluar")
	fmt.Println("6. Keluar")
	fmt.Print("Pilih menu (1-6): ")
}
func baca(A *finansialManage, n *int) {
	var nLangganan string
	var cost int
	fmt.Scan(&nLangganan)
	for nLangganan != "none" {
		fmt.Scan(&cost)
		A[*n].namaLangganan = nLangganan
		A[*n].biaya = cost
		*n++
		fmt.Scan(&nLangganan)
	}
}
func hapusarray(A *finansialManage, n *int){
	var i , j int
	fmt.Print("Pilih nomor yang ingin dihapus: ")
	fmt.Scan(&j)
	for i = 0 ; i<*n ; i++{
		A[j-1]=A[j]
		j++
	}	
	*n = i-1
}
func cetakarr(A finansialManage, n int){
	var i int
	for i = 0; i<n ; i++{
		fmt.Printf("%d. %s %d\n",i+1 , A[i].namaLangganan, A[i].biaya )
	}
}
func totalBiaya(A finansialManage, n int)int{
	var i, hasil int
	hasil = 0
	for i = 0 ; i<n ; i++{
		hasil = hasil + A[i].biaya
	}
	return hasil
}
