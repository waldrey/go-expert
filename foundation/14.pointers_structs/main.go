package main

import "fmt"

type AccountBank struct {
	balance float64
}

func newAccountBank() *AccountBank {
	return &AccountBank{balance: 0.0}
}

func (ab AccountBank) simulate(amount float64) float64 {
	ab.balance += amount

	fmt.Printf("[SIMULAÇÃO] O saldo em sua conta é de %v\n", ab.balance)

	return ab.balance
}

func main() {
	accountBank := AccountBank{balance: 100}

	accountBank.simulate(200)
	fmt.Printf("[REAL] O saldo da sua conta é de %v", accountBank.balance)
}
