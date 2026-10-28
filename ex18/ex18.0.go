/*
슬라이스
	슬라이스는 Go에서 제공하는 동적 배열 타입임

	정적 -> 컴파일 시점에 정해진 것
	동적 -> 동적인 시점에서 정해진 것


	동적 배열 <-- go에서 동적 배열

	c++은 Vector<int> <-- 이렇게 씀
	Java는 ArrayList

	python slice임

	javascript는 [] <-- 기본적으로 동적 배열임

--------------------
	슬라이스 선언

	var slice []int
*/

package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10
	fmt.Println(slice)
}


/*
slice is empty []
panic: runtime error: index out of range [1] with length 0

goroutine 1 [running]:
main.main()
        /Users/eun/Desktop/Go_Study/ex18/ex18.0.go:35 +0x84
exit status 2
	에러는 당황히지 말고 잘 읽어보기
*/