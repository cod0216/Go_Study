// 고는 최강 타입 언어
// 타입이 안맞으면 연산이 안된다.


package main

import "fmt"

func main(){
	a := 3 // int  | 64 bit -> int64
	var b float64 = 3.5 // 실수
	var c int = b // b를 정수로 -> 내 예상 : 소숫점이 없어지지 않을까?  | 타입이 달라서 안된다.
	d := a * b // 3 * 3.5 = 7.5가 되지 않을까 예상	| 타입이 달라서 안된다.
	var e int64 = 7 // e 는 7인 int64형 정수. | a도 int64로 메모리 공간에 있지만 다른 타입이므로 연산이 안된다 
	f := a*e // 3*7 = 21 예상


	// Go는 최강 타입 언어 이다. 서로 같은 형태의 정수더라도 타입이 다르면 안된다.
	// 타입 변환을 해야된다.

	// var c int = int(b)
	// d := float(a) * b
	// f := a*int(e)


	


	fmt.Println(a, b, c, d, e, f)
}