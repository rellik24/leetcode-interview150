package top150

// NOTE: map 可以直接儲值不必檢測，也可以直接 for 取值

func MajorityElement(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	max := 0
	major := 0

	for i, v := range m {
		if v > max {
			max = v
			major = i
		}
	}
	return major
}
