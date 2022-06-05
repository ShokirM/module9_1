package main

import (
	"fmt"
	"github.com/ShokirM/module9_1.git/payment"
	"github.com/ShokirM/module9_1.git/types"
)

func main() {
	Cards := []types.Card {
		{
			Number: "5058 **** **** 1001",
			Balance: 10_000,
			Active: types.Active,
		},
		{
			Number: "5058 **** **** 1002",
			Balance: 14_000,
			Active: types.Active,
		},
		{
			Number: "5058 **** **** 1003",
			Balance: 10_000,
			Active: types.Active,
		},
		{
			Number: "5058 **** **** 1001",
			Balance: 0,
			Active: types.Active,
		},
	}
	fmt.Println(payment.PaymentSource(Cards))
}