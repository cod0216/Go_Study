package main

import "fmt"

func Divide(a, b int) (int, bool) { // 함수 이름이 대문자 시작인지 소문자 시작인지 의미가 있다.
	// Go 에선 타입이 같으면 a int, b int 이렇게 안적고 a, b int 이렇게 적는다.
	if b == 0 {
		return 0, false
	}

	return a/b, true
}

func main () {
	c, success := Divide(9,3)
	fmt.Println(c, success)
	d, success := Divide(9,0)
	fmt.Println(d, success)
}