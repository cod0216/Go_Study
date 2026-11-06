/*
언제 값타입을 쓱 언제 포인터 타입을 쓰는가?

time 패키지

Time와 Timer의 차이

Time 은 값 타입으로 반환  <- 시각
Timer는 포인터 타입으로 반환 <- 20분 알람 : 19분 남은 같은 알람임 (개념적 실체가 같음)

  -> 필드 값이 바뀔때 개념이 바뀌면 값 타입으로 써라


Go는 생성자, 소멸자가 없다.
java는 생성자면 있고 소멸자가 없다.


type Student  -> struct{
	Age int
	name string
}

import (
	"time"
	"fmt"
)

func main() {
	var t1 = time.Timer{}
	var t2 = time.NewTimer(time.second)
	fmt.Println(t1)
}

*/