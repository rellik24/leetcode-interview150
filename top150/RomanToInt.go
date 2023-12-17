package top150

func RomanToInt(s string) int {
	rim := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		v := string(s[i])
		if i+1 < len(s) && rim[v] < rim[string(s[i+1])] {
			sum += rim[string(s[i+1])] - rim[v]
			i++
			continue
		}
		sum += rim[v]
	}
	return sum
}
