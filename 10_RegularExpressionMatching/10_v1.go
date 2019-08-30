package main

import "fmt"

func isMatch(s string, p string) bool {
	sl, pl := len(s), len(p)

	// initial zero-filled 2d array, size: (len(s)+1) * (len(p)+1)
	m := make([][]bool, sl+1)
	for i := range m {
		m[i] = make([]bool, pl+1)
	}

	// fill true in 0th row and 0th col
	m[0][0] = true
	for y := 1; y <= pl; y++ {
		if p[y-1] == '*' {
			m[0][y] = m[0][y-2]
		}
	}

	// dp
	for i := 0; i < sl; i++ {
		for j := 0; j < pl; j++ {
			x, y := i+1, j+1

			if p[j] == s[i] || p[j] == '.' {
				// if match, inherit from previous result
				m[x][y] = m[x-1][y-1]
			} else if p[j] == '*' {
				// if *, see previous char
				if p[j-1] == s[i] || p[j-1] == '.' {
					// matching zero/one/multiple char
					if m[x-1][y] || m[x][y-1] || m[x][y-2] {
						m[x][y] = true
					}
				} else if p[j-1] != s[i] {
					// matching zero char
					m[x][y] = m[x][y-2]
				}
			}
		}
	}

	return m[sl][pl]
}

func main() {
	fmt.Println(isMatch("aaa", "a*a"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aaa", "a*ab"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}
