package main

import "fmt"

/*
특정 조건일때 for문을 종료하고 시픙ㄹ때

1. 플래그 활용
2. 레이블 활용
*/

func main() {
	a := 1
	b := 1
	found := false
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++{
			for c := 1; c <= 9; c++{
				if a * b == 45 {
					found = true
					break
				}
			}
			if found{
				break
			}
		}
		if found{
			break
		}
	}
	fmt.Printf("%d * %d = %d", a, b, a *b)
}