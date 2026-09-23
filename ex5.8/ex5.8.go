package main
import ( // 여러개 패키지 임포트 할때 소괄호 사용
	"bufio" // 
	"fmt" //
	"os" //
)

func main() {
	stdin := bufio.NewReader(os.Stdin) // os 패키지 안에 있는 Stdin 함수 (표준 입력을 나타낸다.)
	// := 선언 대입문, 선언과 대입을 동시에 한다.

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil { //if(err != null) 문제가 생기면
		fmt.Println(err)
		stdin.ReadString('\n') // 표준 입력에서 어떤(\n : 개행 문자) 문자가 나올때 까지 읽어라
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil { //if(err != null) 문제가 생기면
		fmt.Println(err)
		stdin.ReadString('\n') // 표준 입력에서 어떤(\n : 개행 문자) 문자가 나올때 까지 읽어라
	} else {
		fmt.Println(n, a, b)
	}
}