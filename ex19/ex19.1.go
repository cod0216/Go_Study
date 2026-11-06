package main

import (
	"fmt"
)

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int){
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{ 100 } // 100으로 초기화
	
	withdrawFunc(a, 30)

	a.withdrawMethod(30)
	
	fmt.Printf("%d\n", a.balance)

}