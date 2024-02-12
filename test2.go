package main

import (
	"fmt"
	"sort"
	"unicode"
)

type CharCount struct {
	Char  rune
	Count int
}

type ByCountThenChar []CharCount

func (a ByCountThenChar) Len() int      { return len(a) }
func (a ByCountThenChar) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCountThenChar) Less(i, j int) bool {
	if a[i].Count == a[j].Count {
		return a[i].Char < a[j].Char
	}
	return a[i].Count > a[j].Count
}

func countCharacters(words []string) []CharCount {
	charCounts := make(map[rune]int)
	for _, word := range words {
		for _, char := range word {
			if unicode.IsLetter(char) {
				charCounts[char]++
			}
		}
	}

	var charCountSlice []CharCount
	for char, count := range charCounts {
		charCountSlice = append(charCountSlice, CharCount{char, count})
	}

	sort.Sort(ByCountThenChar(charCountSlice))

	return charCountSlice
}

func main() {
	wordsCase1 := []string{"Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti"}
	fmt.Println("======Case 1=====")
	printCharacterCountsAndSortedChars(wordsCase1)
	wordsCase2 := []string{"Abc", "bCd"}
	fmt.Println("======Case 2=====")
	printCharacterCountsAndSortedChars(wordsCase2)
	wordsCase3 := []string{"Oke", "One"}
	fmt.Println("======Case 3=====")
	printCharacterCountsAndSortedChars(wordsCase3)
}


func printCharacterCountsAndSortedChars(words []string) {
	charCountSlice := countCharacters(words)
	var sortedChars string

	fmt.Println("Kelompok setiap huruf beserta jumlahnya:")
	for _, charCount := range charCountSlice {
		fmt.Printf("%c=%d ", charCount.Char, charCount.Count)
		sortedChars += string(charCount.Char)
	}
	fmt.Println()
	fmt.Println("Huruf yang telah diurutkan :", sortedChars)
}