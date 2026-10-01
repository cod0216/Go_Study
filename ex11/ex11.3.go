package main

import "fmt"

func main () {
	i := 0
	for ; i < 10; { // 후처리 i++을 생략할 수 있다.
		fmt.Print(i, ", ")
	}

	/*
		for i < 10 { // 이렇게만 적어도 된다. 단, for i < 10; { 내용 } 이렇게는 안된다.
			내용
		}
	*/
	fmt.Println()
	fmt.Print(i)
}