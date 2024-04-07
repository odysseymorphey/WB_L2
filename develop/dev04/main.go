package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	str := []rune(s)
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})

	return string(str)
}

func SearchAnagrams(data []string) map[string][]string {
	anagrams := make(map[string][]string)

	for _, v := range data {
		v := strings.ToLower(v)
		sortedWord := sortString(v)

		anagrams[sortedWord] = append(anagrams[sortedWord], v)
	}

	for k, v := range anagrams {
		if len(v) == 1 {
			delete(anagrams, k)
		}
	}

	return anagrams
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стиль"}
	fmt.Println(SearchAnagrams(words))
}