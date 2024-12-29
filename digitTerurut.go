package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&input)
	strconv.Atoi(input)
	fmt.Println(checkOrder(input))
}
func checkOrder(num string) string {
	ascending, descending := true, true

	for i := 0; i < len(num)-1; i++ {
		if num[i] > num[i+1] {
			ascending = false
		} else if num[i] < num[i+1] {
			descending = false
		}

		if !ascending && !descending {
			return "tidak terurut"
		}
	}

	if ascending {
		return "ascending"
	} else if descending {
		return "descending"
	} else {
		return "tidak terurut"
	}
}

// func main() {
// 	var input string
// 	fmt.Print("Masukkan angka: ")
// 	fmt.Scan(&input)
// 	strconv.Atoi(input)
// 	fmt.Println(checkOrder(input))

// }
