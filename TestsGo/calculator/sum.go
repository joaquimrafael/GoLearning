package calculator

import "github.com/shopspring/decimal"

func Add(n1, n2 decimal.Decimal) decimal.Decimal {
	return n1.Add(n2)
}
