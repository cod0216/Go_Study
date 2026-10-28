package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 무엇을 랜덤 시드로 하면 좋은가? -> 프로그램 실행시마다 계속 변화하는 값
// func (t Time) UnixNano() int64 - Time 객체 t의 1970년 1월 1일부터 현재까지 경과된 시간을 나노 seconds 단위로 반환

func main(){
	// rand.Seed(50)
	t := time.Now()
	fmt.Println(t)
	rand.Seed(t.UnixNano())
	for i := 0; i < 10; i++{
		fmt.Println(rand.Intn(100))
	}
}

/*
	time 패키지
		시각을 나타내는 Time 객체
		시간을 나타내는 Duration 객체
		타임존을 나태나는 Location	

*/