package main

import "fmt"

var g int = 10 // 패키지 전역변수

func main(){
	var m int = 20

	{
		var s int = 50 // 지역 변수
		fmt.Println(m, s, g)
	}

	// m = s + 20 | undefined: s


}