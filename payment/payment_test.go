package payment

import (
	"fmt"

	"github.com/ShokirM/module9_1.git/payment"
	"github.com/ShokirM/module9_1.git/types"
)

func ExamplePaymentSource() {
	for _, v := range payment.S {
		fmt.Println(v.Number)
	}

	// Output:
	//
	//



}
