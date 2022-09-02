package main

import (
	"fmt"
	"math"
)

func TwoSum(nums []int, target int) []int {
	if target > int(math.Pow10(9)) || target < int(math.Pow(-10, 9)) {
		return []int{}
	}
	c := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] > int(math.Pow10(9)) || nums[i] < int(math.Pow(-10, 9)) {
			return []int{}
		}
		cont := c
		for cont < len(nums) {

			val := nums[i] + nums[cont]
			if val == target {
				return []int{i, cont}
			}
			cont++
		}
		c++
	}
	return []int{}
}

func main() {
	arrT := []int{3, 2, 3}
	tar := 6
	out := TwoSum(arrT, tar)
	fmt.Println(out)
}
