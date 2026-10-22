/*
	range로 한 글자씩 순회
*/


package main

import "fmt"

func main() {
	str := "hello 월드"

	for _, v := range str{
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", v, v, v)
	}
}
