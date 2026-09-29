package main

import "fmt"

const PI = 3.14 // 타입 없는 상수
const FloatPI float64 = 3.14 // float64 타입 상수


func main() {
	var a int = PI * 100 // 오류가 발생하지 않는다.  -> 타입이 없는 상수는 리터럴, 문자 형태로 봐도 된다.
	var b int = FloatPI * 100 // 타입 오류 발생

	fmt.Println(a)
	fmt.Println(b)
}


// 상수는 좌변으로 사용할 수 없다.
// -> 메모리 공간이 없다. (동적 메모리 영역을 쓰지 않는다.)