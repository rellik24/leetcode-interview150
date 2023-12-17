package top150

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	str := strings.ToLower(removeNonLetters(s))
	l := len(str)
	if l == 0 {
		return true
	}
	for i, v := range str {
		if string(v) != string(str[l-1-i]) {
			return false
		}
	}
	return true
}

func removeNonLetters(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return -1
	}, str)
}
