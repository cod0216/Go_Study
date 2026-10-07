package main

import "fmt"
const Y int = 3

func main() {
	nums := [...]int{10, 20, 30, 40, 50} // [5]int{10, 20, ...}
	nums[2] = 300
	for i := 0; i < len(nums); i++{
		fmt.Println(nums[i])
	}
}

// len() -> 내장 함수, 패키지에 소속된 함수가 아니다.
// 여러 자료구조가 들어오면 그 길이를 반환하는 함수다.