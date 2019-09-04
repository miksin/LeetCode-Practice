package main

import (
	"bytes"
	"fmt"
)

func digitRoman(num int, oneUnit, fiveUnit, tenUnit byte) string {
	var roman bytes.Buffer

	switch {
	case num <= 3:
		for i := 0; i < num; i++ {
			roman.WriteByte(oneUnit)
		}
	case num == 4:
		roman.WriteByte(oneUnit)
		roman.WriteByte(fiveUnit)
	case num >= 5 && num <= 8:
		roman.WriteByte(fiveUnit)
		for i := 0; i < num-5; i++ {
			roman.WriteByte(oneUnit)
		}
	case num == 9:
		roman.WriteByte(oneUnit)
		roman.WriteByte(tenUnit)
	}

	return roman.String()
}

func intToRoman(num int) string {
	var symbols []byte = []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M', '?', '?'}

	var roman string
	var symbolOffset int
	for num > 0 {
		digit := num % 10
		dr := digitRoman(digit, symbols[symbolOffset], symbols[symbolOffset+1], symbols[symbolOffset+2])
		roman = dr + roman
		symbolOffset += 2
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
