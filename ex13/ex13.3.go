package main

import "fmt"

type User struct{ // 일반 고객용 구조체
	Name string
	ID string
	Age int	
}

type VIPUser struct { // VIP 고객용 구조체
	UserInfo 	User
	VIPLevel 	int
	Price 		int
}


/*
type VIPUser struct { // VIP 고객용 구조체
	User // 포함된 필드 방식(embedded field)
	VIPLevel 	int
	Price 		int
}



fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP레벨: %d VIP가격: %d만원\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.VIPLevel,
		vip.Price,
	)
*/


func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser {
			User{ "화랑", "hwarang", 48},
			3,
			250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP레벨: %d VIP가격: %d만원\n",
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VIPLevel,
		vip.Price,
	)

}