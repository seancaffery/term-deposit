package term_deposit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_totalBalance(t *testing.T) {
	startAmount := 10000
	interestRate := 2.1
	term := 5
	tests := []struct {
		name            string
		interestPayment string
		expectedBalance float64
	}{
		{
			name:            "interest paid at maturity",
			interestPayment: "maturity",
			expectedBalance: 11050,
		},
		{
			name:            "interest paid annually",
			interestPayment: "annually",
			expectedBalance: 11095,
		},
		{
			name:            "interest paid monthly",
			interestPayment: "monthly",
			expectedBalance: 11106,
		},
		{
			name:            "interest paid quarterly",
			interestPayment: "quarterly",
			expectedBalance: 11104,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TotalBalance(float64(startAmount), interestRate, term, tt.interestPayment)

			assert.Equal(t, tt.expectedBalance, got)
		})
	}
}
