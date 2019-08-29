package main

import "fmt"

func area(i, j, ai, aj int) int {
	if i > j {
		i, j = j, i
	}
	if ai > aj {
		ai, aj = aj, ai
	}
	return (j - i) * ai
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := area(left, right, height[left], height[right])
	for left < right {
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}

		area := area(left, right, height[left], height[right])
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("ans: %d\n", maxArea(arr))
}
