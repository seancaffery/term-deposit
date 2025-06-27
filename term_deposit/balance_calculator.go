package term_deposit

import "math"

type interestPeriod int

const (
	monthly interestPeriod = iota
	quarterly
	annually
	maturity
)

var interestPeriodMapping = map[string]interestPeriod{
	"monthly":   monthly,
	"quarterly": quarterly,
	"annually":  annually,
	"maturity":  maturity,
}

var interestPaidToPeriods = map[interestPeriod]float64{
	monthly:   12,
	quarterly: 4,
	annually:  1,
}

func TotalBalance(startAmount float64, interestRate float64, term int, interestPaymentArg string) float64 {
	interestPayment := interestPeriodMapping[interestPaymentArg]
	interestPercentage := (interestRate / 100)
	if interestPayment == maturity {
		return startAmount + calculateSimpleInterest(startAmount, interestPercentage, term, interestPayment)
	}

	return calculateCompoundInterest(startAmount, interestPercentage, term, interestPayment)
}

func calculateSimpleInterest(startAmount float64, interestPercentage float64, term int, interestPayment interestPeriod) float64 {
	return (startAmount * interestPercentage * float64(term))
}

// A = P x (1 + r)n
//
// A = ending balance
// P = starting balance (or principal)
// r = interest rate per period as a decimal (for example, 2% becomes 0.02)
// n = the number of time periods
func calculateCompoundInterest(startAmount float64, interestPercentage float64, term int, interestPayment interestPeriod) float64 {
	periods := interestPaidToPeriods[interestPayment]
	periodMonths := periods * float64(term)
	interestPerPeriod := interestPercentage / periods
	absoluteTotal := startAmount * math.Pow(1+interestPerPeriod, periodMonths)
	return math.Round(absoluteTotal)
}
