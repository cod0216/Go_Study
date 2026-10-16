package main

import "fmt"

type User struct {
	Name string
	Age int
}

func NewUser(name string, age int) *User{
	var u = User{name, age}
	return &u
} // 여기서 u가 원래 gc가 되어야 됨(지역변수)  

func main() {
	userPointer := NewUser("AAA", 23)
	fmt.Println(userPointer) // 근데 여기서 실행됨
}

/*
	  escape anala 			--> 탈출 분석
		탈출	분석

	stack에 만들어지면 안됨 
	탈출 분석 시 stack에 만들어 지지 않고 heap에 만든다.
	heap은 프로그램 전체에 사용된다.


	포인터는 메모리 주소를 값으로 갖는 타입이다.
	어렵게 생각하면 어렵지만 가볍게 생각하면 쉽다.
*/