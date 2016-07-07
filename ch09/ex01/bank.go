package bank

type withdrawRequest struct {
	amount int
	c      chan<- bool
}

var resets = make(chan struct{})
var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan withdrawRequest)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	c := make(chan bool)
	withdraws <- withdrawRequest{amount: amount, c: c}
	return <-c
}

func reset() { resets <- struct{}{} }

func teller() {
	var balance int
	for {
		select {
		case <-resets:
			balance = 0
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case req := <-withdraws:
			if balance < req.amount {
				req.c <- false
				continue
			}
			balance -= req.amount
			req.c <- true
		}
	}
}

func init() {
	go teller()
}
