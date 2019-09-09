package main

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		sum := 0
		for _, v := range nums {
			sum += v
		}
		return sum
	}

	sort.Ints(nums)
	sum := nums[0] + nums[1] + nums[len(nums)-1]
	minDiff := sum - target
	if minDiff == 0 {
		return sum
	}

	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum = nums[i] + nums[l] + nums[r]
			diff := sum - target
			if diff > 0 {
				r--
			} else if diff < 0 {
				l++
			} else {
				return sum
			}

			if abs(diff) < abs(minDiff) {
				minDiff = diff
			}
		}
	}

	return target + minDiff
}

func main() {
	fmt.Println(threeSumClosest([]int{0, 2, 1, -3}, 1))
}
