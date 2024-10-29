package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)
	wordSet := make(map[string]bool)
	keyMap := make(map[string]string)

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		if wordSet[lowerWord] {
			continue
		}
		wordSet[lowerWord] = true
		sortedWord := sortString(lowerWord)
		_, exists := keyMap[sortedWord]
		if !exists {
			keyMap[sortedWord] = lowerWord
		}
		anagramMap[sortedWord] = append(anagramMap[sortedWord], lowerWord)
	}

	result := make(map[string][]string)
	for key, group := range anagramMap {
		if len(group) > 1 {
			sort.Strings(group)
			result[keyMap[key]] = group
		}
	}

	return result
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "слиток", "слиток", "слиток", "Листок"}
	anagrams := findAnagrams(words)
	for key, group := range anagrams {
		fmt.Printf("Key: %s, Group: %v\n", key, group)
	}
}
