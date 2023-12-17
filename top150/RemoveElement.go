package top150

import "fmt"

func RemoveElement(nums []int, val int) int {
	var count int
	l := len(nums)
	i := 0
	for i < l {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			l--
		} else {
			i++
		}
	}
	fmt.Println(nums)
	return count
}
