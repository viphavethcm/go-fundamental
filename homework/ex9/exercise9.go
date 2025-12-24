package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	account := &BankAccount{balance: 1000}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go account.Withdraw(1, &wg)
	}
	wg.Wait()
	fmt.Printf("Số dư tài khoản: %d\n", account.balance)
	fmt.Printf("Độ dài logs: %d\n", len(account.transactionLogs))
}

type BankAccount struct {
	balance         int
	transactionLogs []string
	mu              sync.Mutex
}

func (account *BankAccount) Withdraw(amount int, wg *sync.WaitGroup) (error, *BankAccount) {
	defer wg.Done()
	account.mu.Lock()
	defer account.mu.Unlock()
	if account.balance < amount {
		return errors.New("Số dư tài khoản không đủ!!"), account
	}
	time.Sleep(1 * time.Millisecond)
	account.balance = account.balance - amount
	account.transactionLogs = append(account.transactionLogs, fmt.Sprintf("Rút %d thành công", amount))
	return nil, account
}
