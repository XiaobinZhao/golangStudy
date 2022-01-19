package bank

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
//var getDraw = make(chan int) // draw

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) (int, error) {
	currentBalance := <-balances
	if amount > currentBalance {
		return 0, fmt.Errorf("amount is greater than balance")
	} else {
		//getDraw <- amount
		Deposit(-amount)
		return amount, nil
	}
}
func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		//case amount := <-getDraw:
		//	balance -= amount
		}
	}
}

// 练习 9.1： 给gopl.io/ch9/bank1程序添加一个Withdraw(amount int)取款函数。
// 其返回结果应该要表明事务是成功了还是因为没有足够资金失败了。这条消息会被发送给monitor的goroutine，
// 且消息需要包含取款的额度和一个新的channel，这个新channel会被monitor goroutine来把boolean结果发回给Withdraw。
func init() {  // init函数初始化
	go teller() // start the monitor goroutine
}

//!-
