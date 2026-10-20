package main

import "fmt"

func main() {
	str := "hello 월드"

	for i := 0; i < len(str); i++ { // len()은 바이트 길이를 반환
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i]) // str[i]로 순회 하면 제대로 된 순회가 되지 않음 '한글'은 3바이트 이기 때문
	}
}