package ex01

import (
	"testing"
)

func TestWithdrawSuccess(t *testing.T) {
	//given
	Deposit(10)
	//when
	ok := Withdraw(5)
	//then
	if result := Balance(); ok && result != 5 {
		t.Errorf("Withdraw result is not 5.actual is %v ok is %v", result, ok)
	}
}

func TestWithdrawUnSuccess(t *testing.T) {
	//given
	//when
	ok := Withdraw(20)
	//then
	if result := Balance(); !ok && result != 5 {
		t.Errorf("Withdraw result is not 10.actual is $v ok is %v", result, ok)
	}
}
