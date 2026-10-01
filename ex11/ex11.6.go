package main

import (
	"bufio"
	"fmt"
	"os"
)

func main () {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("숫자를 입력 하세요 : ")
		var number int
		_, err := fmt.Scanln(&number) //읽어온 값의 개수 _, (필요 없다는 뜻, 빈칸 지시자, 첫번째 리턴값을 안쓰겠다 라는 의미, Go에선 할당한 변수를 반드시 써야됨 안쓰면 에러남)
		if err != nil {
			fmt.Println("숫자로 입력해주세요")

			// 이후 키보드 버퍼를 지운다.
			stdin.ReadString('\n') // 개행 문자가 나올 때까지 읽어온다.
			continue
		}
		fmt.Println("입력하신 숫자는 %d입니다. \n", number)
		if number%2 == 0{
			break // 짝수면 break
		}
	}
	fmt.Println("for문이 종료되었습니다.")
	
}