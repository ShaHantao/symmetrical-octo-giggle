package main

import "fmt"

func main() {
	var a, b, c, d, e, f, g, h int
	d = 0
	fmt.Scanf("%d %d", &a, &b)
	var louti = make([]int, a*b, a*b)
	var fangjian = make([]int, a*b, a*b)
	for i := 0; i < (a * b); i++ {
		fmt.Scanf("%d %d", &louti[i], &fangjian[i])
	}
	fmt.Scanf("%d", &c)
	for j := 1; j <= a; j++ {
		d = d + fangjian[c]
		e = 0
		f = 0
		h = c
		for ; e < fangjian[c]; e++ {
			if h+e > g+b {
				h = h - b
			}
			if louti[h+e] == 1 {
				f++
			}
			if f == fangjian[c] {
				c = h + e + b
				break
			}
		}
	}
	fmt.Println(d)
}
