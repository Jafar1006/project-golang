package main

import (
	"fmt"
)

func LowToUpper(char string) string {
	if char == "a" {
		char = "A"
	} else if char == "b" {
		char = "B"
	} else if char == "c" {
		char = "C"
	} else if char == "d" {
		char = "D"
	} else if char == "e" {
		char = "E"
	} else if char == "f" {
		char = "F"
	} else if char == "g" {
		char = "G"
	} else if char == "h" {
		char = "H"
	} else if char == "i" {
		char = "I"
	} else if char == "j" {
		char = "J"
	} else if char == "k" {
		char = "K"
	} else if char == "l" {
		char = "L"
	} else if char == "m" {
		char = "M"
	} else if char == "n" {
		char = "N"
	} else if char == "o" {
		char = "O"
	} else if char == "p" {
		char = "P"
	} else if char == "q" {
		char = "Q"
	} else if char == "r" {
		char = "R"
	} else if char == "s" {
		char = "S"
	} else if char == "t" {
		char = "T"
	} else if char == "u" {
		char = "U"
	} else if char == "v" {
		char = "V"
	} else if char == "w" {
		char = "W"
	} else if char == "x" {
		char = "X"
	} else if char == "y" {
		char = "Y"
	} else if char == "z" {
		char = "Z"
	} else {
		return char
	}
	return char
}

func main() {
	var karakter string
	fmt.Scan(&karakter)
	fmt.Print(LowToUpper(karakter))
}
