/*
 ==, != 같다 비교
 <, >, <=, >= 대소 비교

 사전식 비교 - 대문자가 더 작다
 'A'-'Z' : 65-90
 'a'-'z' : 97-122

*/

/*
type StringHeader struct {
	Data uintptr // row pointer go에서 기본적으로 쓸수 없는 포인터 내부에선 사용 가능  --> 첫 메모리 주소를 그대로 들고 있는 포인터
	Len  int // 길이
}

100 101 102
[ ] [ ] [ ]   -> Data는 101주소값, Len은 3



64비트는 8바이트 주소값
32비트는 4바이트 주소값


*/

package main

import (
		"unsafe"
		"fmt"
		"reflect"
)

func main() {
	str1 := "hello 월드"
	str2 := str1

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}


