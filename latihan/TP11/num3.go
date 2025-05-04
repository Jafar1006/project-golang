package main
import "fmt"
const NMAX int = 10
type tabInt[NMAX]int
func main(){
	var data tabInt
	var nData int
	var x1 int
	fmt.Scan(&x1)
	fmt.Scan(&nData)
	bacaData(&data, nData)
	cetakData(data, nData)
	fmt.Print("Hasil pencarian: ")
	if sequentialSearch(data, nData, x1) == true {
		fmt.Printf("Bilangan ditemukan. Terdapat %d bilangan %d",frekuensiBilangan(data,nData, x1), x1)
	}else {
		fmt.Print("Bilangan tidak ditemukan")
	}
	fmt.Println()
}	
func bacaData(A *tabInt, n int){
	var i int
	if n >= NMAX {
		n = NMAX
	}
	for i = 0; i<n ; i++{
		fmt.Scan(&A[i])
	}
}
func cetakData(A tabInt, n int){
	var i int
	if n >= NMAX{
		n = NMAX
	}
	fmt.Print("Data Bilangan: ")
	for i=0; i<n; i++{
		fmt.Printf("%d ",A[i])
	}
	fmt.Println()
	
}
func frekuensiBilangan(A tabInt, n, x int)int{
	var frekuensi, i int
	if n >= NMAX {
		n = NMAX
	}
	frekuensi = 0
	for i=0; i<n; i++{
		if A[i] == x{
			frekuensi = frekuensi + 1
		}
	}		
	return frekuensi
}
func sequentialSearch(A tabInt, n, x int)bool{
	var i int
	var ketemu bool
	ketemu = false
	i = 0
	if n >= NMAX {
		n = NMAX
	}
	for !ketemu && i<n{
		ketemu = A[i] == x
		i = i + 1
	}		
	return ketemu
}