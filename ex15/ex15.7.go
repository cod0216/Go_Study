/*
	str[] --> 이거 바이트가 나옴 
	다른 언어는 글자가 나오겠지만 go는 바이트가 나옴
*/


package main

import "fmt"

func main(){
	str1 := "hello"
	str2 := "world"

	str3 := str1 + " " + str2
	fmt.Println(str3)

	str1 += " " + str2
	fmt.Println(str1)


}