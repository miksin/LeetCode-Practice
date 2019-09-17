package main

import (
	"fmt"
)

type Record struct {
	results []string
	stack   []byte
}

func combi(u, d, n int, r *Record) {
	if u == n && d == n {
		r.results = append(r.results, string(r.stack))
		return
	}

	if u == n {
		r.stack = append(r.stack, ')')
		combi(u, d+1, n, r)
	} else if u == d {
		r.stack = append(r.stack, '(')
		combi(u+1, d, n, r)
	} else {
		r.stack = append(r.stack, '(')
		combi(u+1, d, n, r)
		r.stack = r.stack[:len(r.stack)-1]
		r.stack = append(r.stack, ')')
		combi(u, d+1, n, r)
	}

	r.stack = r.stack[:len(r.stack)-1]
}

func generateParenthesis(n int) []string {
	r := Record{
		results: []string{},
		stack:   []byte{},
	}

	combi(0, 0, n, &r)

	return r.results
}

func main() {
	for _, s := range generateParenthesis(3) {
		fmt.Println(s)
	}
	for _, s := range generateParenthesis(1) {
		fmt.Println(s)
	}
}

/*

n = 3

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
