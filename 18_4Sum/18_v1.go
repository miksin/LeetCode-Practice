package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	results := [][]int{}

	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			l, r := j+1, len(nums)-1
			for l < r {
				switch sum := nums[i] + nums[j] + nums[l] + nums[r]; {
				case sum < target:
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					l++
				case sum > target:
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					r--
				case sum == target:
					results = append(results, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					l++
				}
			}

			for j < len(nums)-2 && nums[j] == nums[j+1] {
				j++
			}
		}

		for i < len(nums)-3 && nums[i] == nums[i+1] {
			i++
		}
	}

	return results
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
