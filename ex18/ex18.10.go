/*
	캡사이즈 조절 슬라이싱
	slice[시작인덱스 : 끝인덱스 : 최대인덱스]

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[1:3:4] // -> cap은 4까지


	유용한 슬라이싱 기능 --> 복사
		슬라이스를 그대로 복사해서 새로운 슬라이스를 만들면 두 슬라이스가 서로 영향을 주지 않는다.


*/

package main

import "fmt"

// func main() {
// 	slice1 := []int{1, 2, 3, 4, 5}
// 	slice2 := slice1[:]

// 	slice2[1] = 100
// 	fmt.Println(slice1)
// }

// func main() {
// 	slice1 := []int{1, 2, 3, 4, 5}
// 	slice2 := make([]int, len(slice1))

// 	for i, v := range slice1 {
// 		slice2[i] = v
// 	}

// 	slice2[1] = 100
// 	fmt.Println(slice2)
// 	fmt.Println(slice1)
// }

// func main() {
// 	slice1 := []int{1, 2, 3, 4, 5}
// 	slice2 := append([]int{}, slice1...) // 빈 슬라이스에 append한다 
// 		//  slice1 --> 가변인자 전체 요소가 다 들어간다.

// 	slice2[1] = 100
// 	fmt.Println(slice2)
// 	fmt.Println(slice1)
// }

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	copy(slice1, slice2) // copy(dst,src)

	slice2[1] = 100
	fmt.Println(slice2)
	fmt.Println(slice1)
}