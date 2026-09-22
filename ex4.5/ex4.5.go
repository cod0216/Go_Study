package main

import "fmt"

func main(){
	var a int16 = 3456 // a 는 int16타입 -> 2바이트 정수 (양수 음수 가능)
	var b int8 = int8(a) // int16을 int8로 변환

	fmt.Println(a, b)
	
	/**
		00001101 10000000 a int16 = 3456
		00000000 10000000 b int8  = -128
	*/

	// 큰 타입에서 작은 타입으로 변환할때 값이 짤릴 수 있다.
}