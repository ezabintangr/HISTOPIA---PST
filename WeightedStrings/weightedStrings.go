package main

import (
	"fmt"
	"strings"
)

func WeightedStrings(s string, queries []int) []string {
	weight := make(map[int]bool)
	s = strings.ToLower(s)
	var runeLetter rune
	var letterCount int

	for _, v := range s {
		if v != runeLetter {
			for i := 1; i <= letterCount; i++ {
				mass := int(runeLetter-'a'+1) * i
				weight[mass] = true
			}
			runeLetter = v
			letterCount = 1
		} else {
			letterCount++
		}
	}
	for i := 1; i <= letterCount; i++ {
		mass := int(runeLetter-'a'+1) * i
		weight[mass] = true
	}

	result := make([]string, len(queries))
	for i, q := range queries {
		if weight[q] {
			result[i] = "Yes"
		} else {
			result[i] = "No"
		}
	}
	return result
}

func main() {
	fmt.Println(WeightedStrings("abbcccd", []int{1, 3, 9, 8}))
}
