package main

import "fmt"
const Y int = 3

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2} // [5]int{10, 20, ...}
	for i, v := range t{
		fmt.Println(i, v)
	}
}

// range 범위 
// range t -> 자기 참조
