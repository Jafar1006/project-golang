package main
import "fmt"

func seqSearch(T tabINT,N int, X int)bool{
	var ketemu bool
	var k int
	ketemu = false
	k = 0
	for !ketemu && k < n {
		ketemu = T[k] == x
		k = k + 1
	}		
	return ketemu
}