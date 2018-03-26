package ex01

var deposits = make(chan int)          //入金額を送信
var balances = make(chan int)          //残高を受信する
var withdraw = make(chan withdrawInfo) //出金額を送信
type withdrawInfo struct {
	amount    int
	isSuccess chan bool
}

var done = make(chan bool)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	success := make(chan bool)
	withdraw <- withdrawInfo{amount, success}
	return <-success
}

//バッファありだとtellerで読み出すまで次行かないけどバッファ1以上持ってると自分で書き込んで自分で読めるので
//func Withdraw(amount int)bool{
//	withdraws <-withdrawResult{}
//	result:=<-withdraws
//}
//二人以上同時にwithdraw動かした時に大丈夫か。
//自分の結果が返ってくる保障があるか
//その場でチャンネルを作ってこのチャンネルに返してねを指定しなくてはいけない

func teller() {
	var balance int //balanceはtellerゴルーチンに閉じ込められている
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance: //ここがよくわからない
			//<- balancesだけにするとdeadlockになる
		case info := <-withdraw:
			if info.amount <= balance {
				balance -= info.amount
				info.isSuccess <- true
			} else {
				info.isSuccess <- false
			}
		}
	}

}
func init() {
	go teller() //モニターゴルーチンを開始する
}
