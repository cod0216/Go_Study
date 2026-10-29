/*
	append()

	슬라이스 요소 추가

	슬라이스에 요소를 추가한 슬라이스 반환

*/

package main
import "fmt"

func main() {
	var slice = []int{1, 2, 3}
	slice2 := append(slice, 4)

	fmt.Println(slice) // 기존 슬리이스는 바뀌지 않는다.
	fmt.Println(slice2)

	slice = append(slice, 4) // 이렇게 하면 기존 슬라이스가 바뀜 append는 return을 슬라이스로 한다.
	fmt.Println(slice)
}