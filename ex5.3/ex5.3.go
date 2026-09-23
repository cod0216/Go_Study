package main
import "fmt"

func main(){
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b) // func Scanln(a ...interface{}) (n int, err error) | ...interface{} -> 많이 받는다 라는 뜻 | (n int, err error) -> 출력이 2개
	if err != nil {
				fmt.Println(err)
	} else {
				fmt.Println(n, a, b) // 2 10 20
	}
}

