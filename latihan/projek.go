package main

import "fmt"

const NMAX int = 100

type Subscription struct {
	namaLangganan string
	biaya         int
}
type finansialManage [NMAX]Subscription

func main() {
	var arrData, arrData2 finansialManage
	var nData, nData2 ,cariBiaya, pemasukan, pengeluaran, sisaUang int
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
			fmt.Println(nData)
			cetakArr(arrData, nData)
			cetakUI()
			fmt.Scan(&perintah)
		case "4": //mencari Data berdasarkan nama atau biaya langganan
			fmt.Println("Pilih cari berdasarkan nama atau biaya: (nama/biaya)")
			fmt.Scan(&perintahCari)
			if perintahCari == "nama" {
				cetakArr(arrData, nData)
				fmt.Println("Pilih Nama Langganan yang ingin dicari: ")
				fmt.Scan(&cariNama)
				findName(arrData,&arrData2 ,nData, &nData2, cariNama)
				if nData2 == 0 {
					fmt.Println("Data tidak ditemukan!")
				}else {
					fmt.Println("Data ditemukan!")
					cetakArr(arrData2, nData2)
				}
			} else if perintahCari == "biaya"{
				cetakArr(arrData, nData)
				fmt.Println("Pilih Biaya Langganan yang ingin dicari: ")
				fmt.Scan(&cariBiaya)
				findCost(arrData,&arrData2 ,nData, &nData2, cariBiaya)
				if nData2 == 0 {
					fmt.Println("Data tidak ditemukan!")
				}else {
					fmt.Println("Data ditemukan!")
					cetakArr(arrData2, nData2)
				}
			}
			cetakUI()
			fmt.Scan(&perintah)
		case "5": //mengedit data
			editData(&arrData, nData)	
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
			if	perintahSort == "1"{
				InsertionSort(&arrData, nData)
				cetakArr(arrData, nData)
				cetakUI()
				fmt.Scan(&perintah)
			} else if perintahSort== "2" {
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
func editData(A *finansialManage, n int){
	//I.S Array A dan integer n sudah terdefinisi.
	//F.S Array A pada indeks tertentu telah diperbarui nilainya
	var pilih int	
	var confirm string	
	cetakArr(*A, n)
	fmt.Println("Pilih nomor yang ingin dirubah! ")
	fmt.Scan(&pilih)
	if pilih <= n && pilih >= 1 {
		fmt.Println("Apakah anda yakin ingin menghapusnya? (ya/tidak)")
		fmt.Scan(&confirm)
		if confirm == "ya"{
			fmt.Println("Masukan nama langganan dan biayanya!")
			fmt.Scan(&A[pilih-1].namaLangganan, &A[pilih-1].biaya)
			fmt.Println("Data berhasil diubah!")
			cetakArr(*A, n)
		}
	}
}
func findName(A finansialManage, B *finansialManage, n int, N *int, x string ){
	//I.S Array A dan B, integer n, N, dan string x sudah terdefinisi
	//F.S Array B telah terisi dengan data yang sesuai dengan namaLangganan yang dicari sejumlah N(sequential search)
	*N = 0
	for i:= 0;i<n;i++{
		if A[i].namaLangganan == x {
			B[*N]= A[i]
			*N++
		}
	}
}
func findCost(A finansialManage, B *finansialManage, n int, N *int, x int ){
	//I.S Array A dan B, integer n, N, dan string x sudah terdefinisi
	//F.S Array B telah terisi dengan data yang sesuai dengan biaya Langganan yang dicari sejumlah N(sequential search)
	*N = 0
	for i:= 0;i<n;i++{
		if A[i].biaya == x {
			B[*N]= A[i]
			*N++
		}
	}

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
	var cost int
	fmt.Println("Jika ingin menghentikannya input berhenti/stop/none")
	fmt.Scan(&nLangganan)
	for nLangganan != "berhenti" && nLangganan != "stop" && nLangganan != "none"{
		fmt.Scan(&cost)
		A[*n].namaLangganan = nLangganan
		A[*n].biaya = cost
		*n++
		fmt.Scan(&nLangganan)
	}
}
func hapusArray(A *finansialManage, n *int){
	//I.S A dan n terdefinisi
	//F.S A telah dihapus pada indeks tertentu dan n telah diupdate
	var confirm string
	var i , j int
	fmt.Print("Pilih nomor yang ingin dihapus: ")
	fmt.Scan(&j)
	if j>=1 && j<=*n {
		fmt.Println("Apakah anda yakin ingin menghapusnya? (ya/tidak)")
		fmt.Scan(&confirm)
		if confirm == "ya" {
			for i = 0 ; i<*n ; i++{
				A[j-1]=A[j]
				j++
			}	
			*n = i-1
		}
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
		A[pass -1]= temp
		pass++
	}
}