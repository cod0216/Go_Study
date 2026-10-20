package main

import "fmt"

func main() {
	str := "Hello 월드"
	arr := []rune(str) // rune -> int32의 별칭 타입 4바이트 타입의 정수 타입
	// []rune으로 타입 변환 후 한 글자씩 순회 [4][4][4] 한 글자가 4바이트
	// []int -> 동적 배열

	for i := 0; i < len(arr); i++{
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i]) // 한 칸이 4바이트
	}
}