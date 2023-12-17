package top150

import "fmt"

// 1, 2 都可以，差別在 Rotate1 這種 assign 值的方式，只會改變 func 內的 nums，不會改變外部的
// 要改變外部傳入的 nums，就必須使用 copy，遇到沒有回傳值可能是用指標取值的 func 就要這樣做

// func main() {
// 	nums := []int{-1, -100, 3, 99}
// 	Rotate1(nums, 2)
// 	fmt.Println(nums)
// 	Rotate2(nums, 2)
// 	fmt.Println(nums)
// }

func Rotate1(nums []int, k int) { // 只改變 func 內的 nums
	l := len(nums)
	k = k % l
	nums = append(nums[l-k:], nums[:l-k]...)
	fmt.Println(nums)
}

func Rotate2(nums []int, k int) { // 用 COPY 改變 func 外的 nums
	l := len(nums)
	k = k % l
	tmp := append(nums[l-k:], nums[:l-k]...)
	copy(nums, tmp)
}
