/*
	처음부터 슬라이싱
	
	  slice1 := []int{1, 2, 3, 4, 5}
	  slice2 := slice1[0:3] // slice2는 [1, 2, 3]

	  slice2 := slice1[0:3]
	  slice2 := slice1[:3] // 생략 가능

	끝까지 슬라이싱
	  slice1 := []int{1,2,3,4,5}
	  slice2 := slice1[2:len(slice1)] // slice2는 [3, 4 5]

	  slice2 := slice[2:len(slice1)]
	  slice2 := slice[2:] //생략 가능

	전체 슬라이싱
		array := [5]{1, 2, 3, 4, 5}
		slice := array[:] // 전체 슬라이스
*/

package main
import "fmt"

func main() {
	array := [100]int{1:1, 2:2, 99:100}
	slice1 := array[1:10]

	slice2 := slice1[2:99]
	fmt.Println(slice1)
	fmt.Println(slice2) //slice는 slice를 가리키는 포인터의 원본 값을 본다.
}