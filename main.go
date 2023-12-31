package main

import "fmt"

func main() {
	fmt.Println(IntToRoman(3))
}

func IntToRoman(num int) string {
	s := ""
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	str := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i, v := range nums {
		for num >= v {
			num -= v
			s += str[i]
		}
	}
	return s
}
