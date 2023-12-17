package top150

func TwoSum(numbers []int, target int) []int {
	m := make(map[int]int)

	for i, v := range numbers {
		complement := target - v
		if idx, ok := m[complement]; ok {
			return []int{idx + 1, i + 1}
		}
		m[v] = i
	}
	return []int{}
}
