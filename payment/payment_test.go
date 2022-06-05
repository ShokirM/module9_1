package payment

import (
	"fmt"
	"github.com/ShokirM/module9_1.git/types"
)
func ExamplePaymentSource() {
	Cards := []types.Card {
		{
			Number: "5058 **** **** 1001",
			Balance: 10_000,
			Active: types.Active,
		},
	}
	fmt.Println(PaymentSource(Cards))

	//Output:
	//[{ 5058 **** **** 1001 10000}] 

}