package main
import "fmt"
func SeqSearch(T tabInt, N int, X int)int{
	var idx
	var k int
	idx = -1
	k = 0
	for idx == -1 && k < N {
		if T[k] == X {
			idx = k
		}
		k = k + 1
	}
	return idx
}