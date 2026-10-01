package main

import "fmt"

func main() {

	day := "thursday"

	switch day { // 여러개를 한번에 검사할 수 있음
	case "monday", "tuesday":
		fmt.Println("월, 화요일은 수업 가는 날입니다.")
	case "wednesday", "thursday", "friday":
		fmt.Println("수, 목, 금요일은 실습 가는 날입니다.")
	}
}