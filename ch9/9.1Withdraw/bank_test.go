package bank_test

import (
	"fmt"
	bank "studygolang/ch9/9.1Withdraw"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})
	// Alice
	go func() {
		bank.Deposit(200)
		done <- struct{}{}
		fmt.Println("=", bank.Balance())
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestBank2(t *testing.T) {
	done := make(chan struct{})
	// Alice
	go func() {
		bank.Deposit(200)
		done <- struct{}{}
		//fmt.Println("=", bank.Balance())
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Bob
	go func() {
		//bank.Withdraw(200)
		amount, err := bank.Withdraw(200)
		fmt.Printf("%d, %s \n", amount, err)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
	<-done
	// 多次运行test ，发现，并不是每一次test都能成功。因为3个 go routine 是不确定顺序的。如果刚好余额大于等于200
	// 那么就可以取款成功，否则就会失败
	if got, want := bank.Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
