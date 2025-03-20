// Write a program that starts with a single bank account with a starting balance of Rs.500. It should be possible
// to deposit and withdraw money concurrently. However, the balance update should happen only once at any point in
// time. The withdrawal action should fail if the balance is reaching negative.

// Expectations: Use Goroutines, Mutex

package main

import (
	"fmt"
	"sync"
)

type Account struct {
	mu      sync.Mutex
	balance float64
}

var InsufficientBalance = fmt.Errorf("insufficient balance")

func (account *Account) withdraw(amount float64, wg *sync.WaitGroup, results chan string) {
	defer wg.Done()

	account.mu.Lock()
	defer account.mu.Unlock()

	if account.balance < amount {
		results <- fmt.Sprintf("Withdraw %.2f failed: %v", amount, InsufficientBalance)
	} else {
		account.balance -= amount
		results <- fmt.Sprintf("Withdraw %.2f successful. Remaining Balance: %.2f", amount, account.balance)
	}
}

func (account *Account) deposit(amount float64, wg *sync.WaitGroup, results chan string) {
	defer wg.Done()

	account.mu.Lock()
	defer account.mu.Unlock()

	account.balance += amount
	results <- fmt.Sprintf("Deposit %.2f successful. New Balance: %.2f", amount, account.balance)
}

func main() {
	var wg sync.WaitGroup
	account := Account{
		balance: 500,
	}

	results := make(chan string, 4)

	wg.Add(4)

	go account.deposit(1, &wg, results)
	go account.withdraw(100, &wg, results)
	go account.withdraw(300, &wg, results)
	go account.withdraw(200, &wg, results)

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("Final balance:", account.balance)
}
