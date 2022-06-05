package payment

import "github.com/ShokirM/module9_1.git/types"

var s []types.PaymentSource

func PaymentSource(cards []types.Card) []types.PaymentSource {
	for _, card := range cards {
		if card.Balance <= 0 {
			continue
		}
		if !card.Active {
			continue
		}
		s = append(s, types.PaymentSource{Number: card.Number, Balance: card.Balance})
	}
	return s
}
