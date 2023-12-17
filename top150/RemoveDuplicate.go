package top150

func RemoveDuplicates(nums []int) int {
	m := make(map[int]int)

	var tmp []int
	for _, v := range nums {
		if m[v] < 2 { // 保留最多兩次
			m[v]++
			tmp = append(tmp, v)
		}
	}

	copy(nums, tmp)
	return len(tmp)
}
