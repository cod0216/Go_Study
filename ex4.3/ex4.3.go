package main

import "fmt"

func main() {
	var a int = 3 // 8바이트
	var b int // 기본값 0
	var c = 4 // 4가 정수니까 c는 int
	d := 5 // 초기값 5

	fmt.Println(a, b, c, d) // 3 0 4 5

	var e = "Hello"
	f :=3.14

	fmt.Println(a, b, c, d, e, f)

	
}
