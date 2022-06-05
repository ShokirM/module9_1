package types

type Money int64

const (
	Active   = true
	Inactive = false
)

type Card struct {
	Number  string
	Balance Money
	Active  bool
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
