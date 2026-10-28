package main

import (
	"fmt"
	"Go_Study/usepkg/custompkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}



/*
	go mod tidy // <-- 해당 패키지에서 import 하는 모듈을 찾고 다운 받는다.


	go env에서 go package <-- 여기에 package다운받아서 사용함
*/