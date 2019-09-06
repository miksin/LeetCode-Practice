package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	var results [][]int = [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}

		l := i + 1
		r := len(nums) - 1
		for l < r {
			switch sum := nums[l] + nums[r] + nums[i]; {
			case sum == 0:
				results = append(results, []int{nums[i], nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			case sum > 0:
				r--
				for nums[r] == nums[r+1] && l < r {
					r--
				}
			case sum < 0:
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}

		for i < len(nums)-2 && nums[i] == nums[i+1] {
			i++
		}
	}

	return results
}

func main() {
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
