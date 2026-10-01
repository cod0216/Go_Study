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

OuterFor: // Label 레이블, goto문도 있음
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++{
			if a * b == 45 {
				break OuterFor
			}
		}
	}
	fmt.Printf("%d * %d = %d", a, b, a *b)
}

// 레이블 되도록 안쓰는게 좋다. instructure 포인트를 강제로 바꿔서 상태가 꼬일 수 있다.
// goto 문도 위험한 문법, 강력하지만 위험
// 레이블은 되도록 안쓰는게 좋다!