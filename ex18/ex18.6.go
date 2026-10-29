// package main

// import (
// 		"fmt"
// )

// func addNum(slice *[]int){ // <-- copy되는 크기가 작음, 메모리 주소가 작으니
// 	*slice = append(*slice, 4)
// }

// func main() {
// 	slice := []int{1, 2, 3}
// 	addNum(&slice)

// 	fmt.Println(slice)
// }
package main

import (
		"fmt"
)

func addNum(slice []int) []int{ // 일반적인 경우에 이게 더 좋음, 
	slice = append(slice, 4)
	return slice
}
// 복사 되고 리턴되고 싫어하는데 슬라이스는 주소 값이여서 24바이터만 복사되고 성능상 문제가 되지 않는다. 값 타임임(append는 값을 받아서 값을 반환하는 거임)
// 포인터 형태로 쓰지 말고 값 타입으로 써라 --> 처음부터 값 타입으로 쓰라고 만들어 놓았음
func main() {
	slice := []int{1, 2, 3}
	slice = addNum(slice)

	fmt.Println(slice)
}


// 값 타입으로 쓰는거에 대한 알러지가 없어도 된다.

// 메모리를 많이 쓰는게 string이 있는데 string도 마찬가지로 내부에 pointer를 가지고 있음
