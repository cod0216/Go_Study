/*
	슬라이스 동작원리
	type SliceHeader struct {
		Data uintptr // 실제 배열을 가리키는 포인터
		Len int		// 요소 개수 -> 컨텐츠 개수
		Cap int		// 실제 배열의 길이 -> capacity
	}
	사실 슬라이스는 Go에서 제공하는 배열을 가리키는 포인터 타입이다.

*/

package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println(array)
	fmt.Println(slice)
}

/*
 Go의 특징 중 하나가 일관성이다.

 함수를 호출하면 
 lvalue로 동작하나 rvalue로 동작하나
changeArray(array) <-- array는 ravlue로 동작

즉 array2 = array <-- 대입 연산자가 된다. 

array2 에 array 값을 복사한다.

*/