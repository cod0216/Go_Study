package main
import (
	"fmt"
	"sort"
)


type Student struct {
		Name string
		Age int
}
type Students []Student

func (s Students) Len() int {return len(s)}
func (s Students) Less(i, j int) bool {return s[i].Age < s[j].Age}
func (s Students) Swap(i, j int) {s[i], s[j] = s[j], s[i]}


func main() {
	slice := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(slice) // int 슬라이스 정렬
	fmt.Println(slice)


	s := []Student{
			{"화랑", 31},
			{"백두산", 52},
			{"류", 32},
			{"켄", 42},
			{"송하나", 18},
	}

	sort.Sort(Students(s)) // Sort(interface)  -> len 등 메서드를 가지고 있는 것 만 사용 할 수 있음
	fmt.Println(s)
}

/*
	type Student struct {
		Name string
		Age int
	}

	구조체 타입 정렬
	type Students []Student // []Student의 별칭 타입 Students

*/