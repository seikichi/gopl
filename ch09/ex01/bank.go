package bank

type withdrawRequest struct {
	amount int
	c      chan bool
}

var deposits = make(chan int)              // send amount to deposit
var balances = make(chan int)              // receive balance
var withdraws = make(chan withdrawRequest) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	c := make(chan bool)
	withdraws <- withdrawRequest{amount: amount, c: c}
	return <-c
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case req := <-withdraws:
			if balance < req.amount {
				req.c <- false
				continue
			}
			balance -= req.amount
			req.c <- false
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
