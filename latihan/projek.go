package main

import "fmt"

const NMAX int = 100

type Subscription struct {
	namaLangganan string
	biaya         int
}
type finansialManage [NMAX]Subscription

func main() {
	var arrData finansialManage
	var nData,cariBiaya, pemasukan, pengeluaran, sisaUang, idx int
	var perintah, perintahSort,perintahCari, cariNama string
	cetakUI()
	fmt.Scan(&perintah)
	for {
		switch perintah {
		case "1": //menambahkan pemasukan uang
			masukUang(&pemasukan)
			cetakUI()
			fmt.Scan(&perintah)
		case "2": //menambahkan subscription apa saja
			baca(&arrData, &nData)
			fmt.Println(nData)
			cetakUI()
			fmt.Scan(&perintah)
		case "3": //menghapus langganan
			cetakArr(arrData, nData)
			hapusArray(&arrData, &nData)
			cetakUI()
			fmt.Scan(&perintah)
		case "4": //mencari Data berdasarkan nama atau biaya langganan
			fmt.Println("Pilih cari berdasarkan nama atau biaya: (nama/biaya)")
			fmt.Scan(&perintahCari)
			if perintahCari == "nama" {
				cetakArr(arrData, nData)
				fmt.Println("Pilih Nama Langganan yang ingin dicari: ")
				fmt.Scan(&cariNama)
				idx = findIdxName(arrData ,nData, cariNama)
				if idx != -1 {
					fmt.Printf("Data ditemukan\n%s %d",arrData[idx].namaLangganan, arrData[idx].biaya)			
				} else {
					fmt.Println("Data tidak ditemukan")
				}
			} else if perintahCari == "biaya"{
				cetakArr(arrData, nData)
				fmt.Println("Pilih Biaya Langganan yang ingin dicari: ")
				InsertionSort(&arrData,nData)
				fmt.Scan(&cariBiaya)
				idx = findIdxCost(arrData ,nData, cariBiaya)
				if idx != -1 {
					fmt.Printf("Data ditemukan\n%s %d",arrData[idx].namaLangganan, arrData[idx].biaya)
				} else {
					fmt.Println("Data tidak ditemukan")
				}
			}
			cetakUI()
			fmt.Scan(&perintah)
		case "5": //mengedit data
			cetakArr(arrData, nData)
			editDataNama(&arrData, nData)	
			cetakArr(arrData, nData)
			cetakUI()
			fmt.Scan(&perintah)
		case "6":	//detail semua langganan dan pengeluaran
			fmt.Println("Detail Total Pengeluaran: ")
			cetakArr(arrData, nData)
			pengeluaran = totalBiaya(arrData, nData)
			fmt.Printf("Total Biaya sebesar: Rp. %d\n", pengeluaran)
			if pengeluaran == 0 {
				fmt.Printf("Uang yang dimiliki: Rp. %d\n", pemasukan)
			}else if pemasukan >= pengeluaran{ 
				sisaUang = pemasukan - pengeluaran
				fmt.Printf("Uang yang dimiliki: Rp. %d\n", pemasukan)
				fmt.Printf("Sisa uang : Rp. %d\n", sisaUang)
			}else {
				sisaUang = -1 * (pemasukan - pengeluaran)
				fmt.Printf("Uang yang tersedia: Rp. %d\n", pemasukan)
				fmt.Printf("Uang anda kurang Rp. %d \n", sisaUang)
			}
			fmt.Printf("Banyaknya langganan: %d\n", nData)
			cetakUI()
			fmt.Scan(&perintah)
		
		case "7": // mnegurutkan data secara ascending atau descending
			cetakUISort()
			fmt.Scan(&perintahSort)
			if	perintahSort == "1" || perintahSort == "ascending"{
				InsertionSort(&arrData, nData)
				cetakArr(arrData, nData)
				cetakUI()
				fmt.Scan(&perintah)
			} else if perintahSort== "2" || perintahSort == "descending" {
				SelectionSort(&arrData, nData)
				cetakArr(arrData, nData)
				cetakUI()
				fmt.Scan(&perintah)
			}
		case "8":
			return
		default:
			fmt.Println("✖ Pilihan tidak valid.")
			cetakUI()
			fmt.Scan(&perintah)
		}
	}
}
func editDataNama(A *finansialManage, n int){
	//I.S Array A dan integer n sudah terdefinisi.
	//F.S Array A pada indeks tertentu telah diperbarui nilainya
	var idx, biaya int	
	var confirm, x, nama string
	fmt.Println("Pilih nama yang ingin diubah:")
	fmt.Scan(&x)
	idx = findIdxName(*A , n, x)
	if idx != -1 {
		fmt.Println("Masukan nama langganan dan biayanya!")
		fmt.Scan(&nama, &biaya)
		fmt.Println("Apakah anda yakin ingin mengedit nama dan biayanya? (ya/tidak)")
		fmt.Scan(&confirm)
		if confirm == "ya"{
			A[idx].namaLangganan = nama 
			A[idx].biaya = biaya
			fmt.Println("Data berhasil diubah!")
		}
	}else{
		fmt.Println("Data tidak ditemukan")
	}
}
func findIdxName(A finansialManage, n int, x string )int{
	//Array A, integer n, dan string x sudah terdefinisi,  akan mengembalikan nilai index, jika tidak ditemukan maka akan mengembalikan -1 (sequential search)
	
	var i, ketemu int
	ketemu = -1
	for i<n && ketemu == -1{
		if A[i].namaLangganan == x {
			ketemu = i
		}
		i++
	}
	return ketemu
}
func findIdxCost(A finansialManage, n int, x int )int{
	//Array A, integer n, dan integer x sudah terdefinisi, akan mengembalikan nilai index, jika tidak ditemukan maka akan mengembalikan -1 (binary search)

	var left, mid, right, ketemu int
	left = 0
	right = n-1
	ketemu = -1
	for left<=right && ketemu == -1 {
		mid = (left+right)/2
		if A[mid].biaya > x {
			right = mid - 1
		}else if A[mid].biaya < x{
			left = mid + 1
		}else {
			ketemu = mid
		}
	}
	return ketemu
}
func cetakUISort(){
	//I.S
	//F.S menampilkan menu sort
	fmt.Println("\n===== Aplikasi Manajemen Subskripsi & Keuangan =====")
	fmt.Println("Pilih urut berdasarkan :")
	fmt.Println("1. Urut dari kecil ke besar(ascending)")
	fmt.Println("2. urut dari besar ke kecil(descending)")
}
func cetakUI() {
	//I.S
	//F.S menampilkan menu utama
	fmt.Println("\n===== Aplikasi Manajemen Subskripsi & Keuangan =====")
	fmt.Println("1. Masukan Pemasukan")
	fmt.Println("2. Tambah Langganan")
	fmt.Println("3. Hapus Langganan")
	fmt.Println("4. Cari Data")
	fmt.Println("5. Edit Data")
	fmt.Println("6. Lihat Total Pengeluaran Langganan")
	fmt.Println("7. Pengurutan")
	fmt.Println("8. Keluar")
	fmt.Print("Pilih menu (1-8): ")
}
func masukUang(A *int){
	//I.S A belum terdifinisi
	//F.S A telah terdefinisi jumlah nominal uang yang dimasukkan
	var pemasukan int
	fmt.Println("Masukan Pemasukan Anda!:")
	fmt.Scan(&pemasukan)
	*A = *A + pemasukan
	fmt.Printf("Pemasukan sebesar Rp. %d berhasil dimasukan!\nTotal uang anda sekarang adalah Rp. %d\n", pemasukan, *A)
}
func baca(A *finansialManage, n *int) {
	//I.S A dan n belum terdefinisi
	//F.S A telah terdefinisi dan n telah terisi
	var nLangganan string
	var cost, i int
	i = 0
	fmt.Println("Jika ingin menghentikannya input berhenti/stop/none")
	fmt.Scan(&nLangganan)
	for i<NMAX && nLangganan != "berhenti" && nLangganan != "stop" && nLangganan != "none"{
		fmt.Scan(&cost)
		A[i].namaLangganan = nLangganan
		A[i].biaya = cost
		i++
		fmt.Scan(&nLangganan)
	}
	*n = i
}
func hapusArray(A *finansialManage, n *int){
	//I.S A dan n terdefinisi
	//F.S A telah dihapus pada indeks tertentu dan n telah diupdate
	var confirm, x string
	var i , j int
	fmt.Print("Cari data nama yang ingin dihapus: ")
	fmt.Scan(&x)
	j = findIdxName(*A, *n , x)
	if j != -1 {
		fmt.Println("Apakah anda yakin ingin menghapusnya? (ya/tidak)")
		fmt.Scan(&confirm)
		if confirm == "ya" {
			for i = j ; i<*n ; i++{
				A[i]=A[i+1]
				j++
			}	
			*n = i-1
			fmt.Println("Data berhasil dihapus")
			cetakArr(*A, *n)
		}
	}else {
		fmt.Println("Data yang ingin dihapus tidak ditemukan")
	}
}
func cetakArr(A finansialManage, n int){
	//I.S A dan n terdefinisi
	//F.S menampilkan semua elemen A sejumlah n  
	var i int
	for i = 0; i<n ; i++{
		fmt.Printf("%d. %s %d\n",i+1 , A[i].namaLangganan, A[i].biaya )
	}
}
func totalBiaya(A finansialManage, n int)int{
	//A dan n terdefinisi, kemudian akan mengembalikan nilai integer yaitu total hasil dari seluruh biaya langganan
	var i, hasil int
	hasil = 0
	for i = 0 ; i<n ; i++{
		hasil = hasil + A[i].biaya
	}
	return hasil
}
func InsertionSort(A *finansialManage, n int){
	//I.S A dan n terdefinisi
	//F.S A dan n telah disorting berdasarkan biaya langganan secara ascending menggunakan algoritma insertion sort 
	var pass, i int
	var temp Subscription
	pass = 1
	for pass < n {
		i = pass
		temp = A[pass]

		for i > 0 && A[i-1].biaya > temp.biaya{
			A[i]=A[i-1]
			i--
		}
		A[i]= temp
		pass++
	}
}
func SelectionSort(A *finansialManage, n int){
	//I.S A dan n terdefinisi
	//F.S A dan n telah disorting berdasarkan biaya langganan secara descending menggunakan algoritma selection sort 
	var i,idx, pass int
	var temp Subscription
	pass = 1
	for pass < n{
		idx= pass - 1
		for i = pass; i < n;i++{
			if A[i].biaya > A[idx].biaya{
				idx = i
			}
		} 
		temp = A[idx]
		A[idx] = A[pass-1]
		A[pass-1]= temp
		pass++
	}
}