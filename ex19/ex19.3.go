package main

import (
	"fmt"
)

type myInt int

func (m myInt) Add (a int) myInt {
	rst := int(m) + a
	return myInt(rst)
}

func (a int) Del(b int) int { // int 는 외부 타입이니까 안됨 리시버는 지역 타입만 쓸 수 있다.
	return a + b
} //cannot define new methods on non-local type int

func addFunc(m myInt, a int) myInt{
	rst := int(m) + a
	return myInt(rst)
}

func main() {
	var a myInt
	a = a.Add(10)

	fmt.Println(a)
}

