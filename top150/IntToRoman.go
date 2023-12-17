package top150

// https://leetcode.com/problems/integer-to-roman/description/?envType=study-plan-v2&envId=top-interview-150
// 暴力窮舉法: 由於減號的判斷較麻煩再加上題目限制較小 <=3999，列出所有減號的再來處理

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
