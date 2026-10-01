package main

import ("fmt" 
		"time")

func main () {

	// 무한 루프
	for i := 0; true; i++ { // for true { 내용 }, for { 내용 } --> 무한 루프
		time.Sleep(time.Second)
		fmt.Print(i, ", ")
	}
}