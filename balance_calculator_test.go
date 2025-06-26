package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_totalBalance(t *testing.T) {
	tests := []struct {
		name            string
		startAmount     float64
		interestRate    float64
		term            int
		interestPayment interestPeriod
		expectedBalance float64
	}{
		{
			name:            "interest paid at maturity",
			startAmount:     10000,
			interestRate:    2.1,
			term:            5,
			interestPayment: maturity,
			expectedBalance: 11050,
		},
		{
			name:            "interest paid yearly",
			startAmount:     10000,
			interestRate:    2.1,
			term:            5,
			interestPayment: yearly,
			expectedBalance: 11095,
		},
		{
			name:            "interest paid monthly",
			startAmount:     10000,
			interestRate:    2.1,
			term:            5,
			interestPayment: monthly,
			expectedBalance: 11106,
		},
		{
			name:            "interest paid quarterly",
			startAmount:     10000,
			interestRate:    2.1,
			term:            5,
			interestPayment: quarterly,
			expectedBalance: 11104,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := totalBalance(tt.startAmount, tt.interestRate, tt.term, tt.interestPayment)

			assert.Equal(t, tt.expectedBalance, got)
		})
	}
}
