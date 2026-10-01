package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main () {
	if true && IncreaseAndReturn() < 5 { // and면 앞에가 false 면 뒤엔 검사 안함 
		fmt.Println("1 증가")
	}
}