package term_deposit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTermDeposit_ValidateArguments(t *testing.T) {
	tests := []struct {
		name               string
		startAmount        float64
		interestRate       float64
		term               int
		interestPaymentArg string
		expectedError      error
	}{
		{
			name:               "does not return an error for valid arguments",
			startAmount:        1,
			interestRate:       1.1,
			term:               3,
			interestPaymentArg: "maturity",
			expectedError:      nil,
		},
		{
			name:               "returns an error for a single invalid argument",
			startAmount:        -1,
			interestRate:       1.1,
			term:               3,
			interestPaymentArg: "maturity",
			expectedError:      fmt.Errorf("startingBalance must be greater than 0"),
		},
		{
			name:               "returns an error for a multiple invalid argument",
			startAmount:        -1,
			interestRate:       -1.1,
			term:               3,
			interestPaymentArg: "maturity",
			expectedError:      fmt.Errorf("startingBalance must be greater than 0, got '-1'\ninterestRate must be greater than 0, got '-1.1'"),
		},
		{
			name:               "returns an error for an invalid interest period",
			startAmount:        1,
			interestRate:       1.1,
			term:               3,
			interestPaymentArg: "never",
			expectedError:      fmt.Errorf("interestPaid must be one of: monthly, quarterly, annually, maturity. Got 'never'"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := TermDeposit{
				StartingBalance: tt.startAmount,
				InterestRate:    tt.interestRate,
				TermYears:       tt.term,
				InterestPaid:    tt.interestPaymentArg,
			}
			err := td.validateArguments()
			if tt.expectedError != nil {
				assert.ErrorContains(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
