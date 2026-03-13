package amortization

import (
	"math"
)

func CalcMonthlyPayment(origBalance float64, interestRate float64, durationInMonths int) float64 {
	i := interestRate / 12
	onePlusItoN := math.Pow((1 + i), float64(durationInMonths))
	return origBalance * (i * onePlusItoN) / (onePlusItoN - 1)
}
