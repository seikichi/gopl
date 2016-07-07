package bank

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	reset()
	Deposit(300)

	if Withdraw(500) {
		t.Errorf("Deposit(300); Withdraw(500) = true, want false")
	}

	if !Withdraw(200) {
		t.Errorf("Deposit(300); Withdraw(200) = false, want true")
	}

	if got, want := Balance(), 100; got != want {
		t.Errorf("Deposit(300); Withdraw(200); Balance() = %d, want %d", got, want)
	}
}

func TestBank(t *testing.T) {
	reset()
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
