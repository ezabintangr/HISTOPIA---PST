package main

import "fmt"

func balanceBracket(s string) string {
	matchBracket := []rune{}
	bracketPair := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			matchBracket = append(matchBracket, v)
		} else if v == ')' || v == '}' || v == ']' {
			if len(matchBracket) == 0 || matchBracket[len(matchBracket)-1] != bracketPair[v] {
				return "NO"
			}
			matchBracket = matchBracket[:len(matchBracket)-1]
		}
	}
	return "YES"
}

func main() {
	fmt.Println(balanceBracket("( [ { ( ( [ ] ) [ ] ) [ ] } ] )"))
}
