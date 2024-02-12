package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countCharacters(str string) map[rune]int {
	charCounts := make(map[rune]int)

	str = strings.ToLower(str)
	for _, char := range str {
		if unicode.IsLetter(char) {
			charCounts[char]++
		}
	}

	return charCounts
}

func main() {
	input := "We Always Mekar"
	charCounts := countCharacters(input)


	
	fmt.Println("Jumlah setiap huruf dalam string Case 1:")
	for char, count := range charCounts {
		fmt.Printf("%c=%d, ", char, count)
	}
	fmt.Println()

	input2 := "coding is fun"
	charCounts2 := countCharacters(input2)

	fmt.Println("Jumlah setiap huruf dalam string Case 2:")
	for char, count := range charCounts2 {
		fmt.Printf("%c=%d, ", char, count)
	}
	fmt.Println()



}
