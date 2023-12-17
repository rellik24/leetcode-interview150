package top150

// func main() {
// 	fmt.Println(isSubsequence("axc", "ahbgdc"))
// }

func IsSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	i := 0
	for _, v := range t {
		if string(v) == string(s[i]) {
			i++
		}
		if i == len(s) {
			return true
		}
	}

	return false
}
