package ex01

var deposits = make(chan int) //入金額を送信
var balances = make(chan int) //残高を受信する
var withdraw = make(chan int) //出金額を送信

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}
func Withdraw(amount int) {
	withdraw <- amount
}

func teller() {
	var balance int //balanceはtellerゴルーチンに閉じ込められている
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance: //ここがよくわからない
			//<- balancesだけにするとdeadlockになる
		case amount := <-withdraw:
			balance -= amount
		}
	}

}
func init() {
	go teller() //モニターゴルーチンを開始する
}
