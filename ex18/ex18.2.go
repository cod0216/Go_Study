/*
	Make()를 이용한 초기화

	var slice = make([]int, 3) // 내장 함수 slice := []int{0, 0, 0}
	slice[1] = 5

*/

package main
import "fmt"

func main() {
	var slice = []int{1, 2, 3}

	for i := 0; i < len(slice); i++{
		slice[i] += 10
		fmt.Print(slice[i], " ")
	}

	fmt.Println()

	for i, v := range slice{
		slice[i] = v*2
		fmt.Print(slice[i], " ")
	}

	fmt.Println()
}
