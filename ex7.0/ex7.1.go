package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func main(){
	c := Add(3, 6)

	fmt.Println(c)
}

// 함수 왜 쓰나?
// 반복 작업이 싫어서 사용함
