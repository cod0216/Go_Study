/*
슬라이스
	var slice1 = []int{1, 2, 3} // 대괄호 안에 길이를 넣지 않은 것을 주의
	var slice2 = []int{1, 5:2, 10:3} // [1 0 0 0 0 2 0 0 0 0 3]


주의!
	var array = [...]int{1, 2, 3} // 배열 선언
	var slice = []int{1, 2, 3} // 슬라이스 선언
*/

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10
	fmt.Println(slice)
}
