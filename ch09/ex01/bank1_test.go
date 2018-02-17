package ex01

import (
	"testing"
	"fmt"
)

func TestWithdraw(t *testing.T) {
	//given
	Deposit(10)
	fmt.Println(Balance())
	//when
	ok := Withdraw(5)
	//then
	if result := Balance(); ok && result != 5 {
		t.Errorf("Withdraw result is not 5.actual is ", result)
	}
}
