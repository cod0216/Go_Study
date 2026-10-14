package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 	//1바이트?
	B int	//8바이트?
	C int8	//1바이트?
	D int	//8바이트?
	E int8	//1바이트?
			// 19바이트?
}



type UserA struct {
	B int	//8바이트?
	D int	//8바이트?
	A int8 	//1바이트?
	C int8	//1바이트?
	E int8	//1바이트?
			// 19바이트 + 5바이트
}

func main() {
	user := User{1, 2, 3, 4, 5}
	userA := UserA{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(userA)) // 40바이트 why? pedding때문!  성능때문에! 대신 메모리 낭비
	fmt.Println(unsafe.Sizeof(user)) // 24바이트 --> 8 + 8 + 8 
	// 큰 PC에서는 반드시 신경 써야 되는건 아님 임베디드에선 필요할지도
}