// Copyright 2016 budougumi0617 All Rights Reserved.

// Package bank provides a concurrency-safe bank with one account.
package bank

// Withdrawal incluse amount and bool channel.
type Withdrawal struct {
	amount  int
	success chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdrawals = make(chan Withdrawal)

// Deposit sends amount to deposit.
func Deposit(amount int) { deposits <- amount }

// Balance receives balance.
func Balance() int { return <-balances }

// Withdraw tries operation.
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawals <- Withdrawal{amount, ch}
	return <-ch
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdrawals:
			if w.amount > balance {
				w.success <- false
				continue
			}
			balance -= w.amount
			w.success <- true
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
