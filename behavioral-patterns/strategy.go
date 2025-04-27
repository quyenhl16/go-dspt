// Refer to: https://refactoring.guru/design-patterns/strategy
/*
Strategy is a behavioral design pattern that lets you define a family of algorithms,
put each of them into a separate class, and make their objects interchangeable.
*/
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// PaymentStrategy Strategy interface with error handling
type PaymentStrategy interface {
	Pay(amount float64) error
}

// CreditCardPayment Concrete Strategy: Credit Card
type CreditCardPayment struct {
	Name   string
	CardNo string
}

func (c *CreditCardPayment) Pay(amount float64) error {
	if amount > 1000 {
		return errors.New("credit card limit exceeded")
	}
	fmt.Printf("[CreditCard] %s paid $%.2f using card ending in %s\n",
		c.Name, amount, c.CardNo[len(c.CardNo)-4:])
	return nil
}

// PayPalPayment Concrete Strategy: PayPal
type PayPalPayment struct {
	Email string
}

func (p *PayPalPayment) Pay(amount float64) error {
	if rand.Float32() < 0.1 {
		return errors.New("PayPal service unavailable")
	}
	fmt.Printf("[PayPal] Payment of $%.2f completed via %s\n", amount, p.Email)
	return nil
}

// CryptoPayment Concrete Strategy: Crypto
type CryptoPayment struct {
	WalletAddress string
}

func (c *CryptoPayment) Pay(amount float64) error {
	if amount < 10 {
		return errors.New("crypto transaction too small")
	}
	fmt.Printf("[Crypto] Transferred $%.2f to wallet %s\n", amount, c.WalletAddress)
	return nil
}

// PaymentContext Context
type PaymentContext struct {
	strategy PaymentStrategy
}

func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	pc.strategy = strategy
}

func (pc *PaymentContext) Pay(amount float64) {
	if pc.strategy == nil {
		fmt.Println("[Error] Payment strategy not set")
		return
	}
	err := pc.strategy.Pay(amount)
	if err != nil {
		fmt.Printf("[Payment Failed] %v\n", err)
	} else {
		fmt.Println("[Payment Successful]")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	users := []struct {
		name     string
		strategy PaymentStrategy
		amount   float64
	}{
		{"Alice", &CreditCardPayment{"Alice", "1234-5678-9012-3456"}, 500},
		{"Bob", &PayPalPayment{"bob@example.com"}, 120},
		{"Charlie", &CryptoPayment{"0xABC123XYZ"}, 5},                   // Will fail
		{"Dave", &CryptoPayment{"0xDEF456LMN"}, 150},                    // Will pass
		{"Eve", &CreditCardPayment{"Eve", "0000-1111-2222-3333"}, 2000}, // Will fail
	}

	for _, u := range users {
		fmt.Printf("\nProcessing payment for %s...\n", u.name)
		ctx := &PaymentContext{}
		ctx.SetStrategy(u.strategy)
		ctx.Pay(u.amount)
	}
}
