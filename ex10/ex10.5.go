package main

import "fmt"

type ColorType int // 값의 의미를 나타내기 위한 코드 값 (값이 중요한게 아님 의미가 중요한거임)
const (
	Red ColorType = iota 
	Blue
	Green
	Yellow 
)

func colorToString(color ColorType) string {

	/*
		다른 언어에선 break 써야 되지만 Go에선 break 안써줘도 된다(써도 되는데 생략해도 된다는 뜻)
		fallthrough 사용해서 case가 중복이면 case 계속 진행
	*/
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Red
}

func main() {
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}
