package main

import (
	"fmt"
)

func intToRoman(num int) string {
	var symbolDict [][]string = make([][]string, 4)
	symbolDict[0] = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	symbolDict[1] = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	symbolDict[2] = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	symbolDict[3] = []string{"", "M", "MM", "MMM", "", "", "", "", "", ""}

	var roman string
	var symbolOffset int
	for num > 0 {
		digit := num % 10
		dr := symbolDict[symbolOffset][digit]
		roman = dr + roman
		symbolOffset += 1
		num /= 10
	}

	return roman
}

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}

/*
Example 1:
Input: 3
Output: "III"

Example 2:
Input: 4
Output: "IV"

Example 3:
Input: 9
Output: "IX"

Example 4:
Input: 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.

Example 5:
Input: 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/
