package term_deposit

import (
	"errors"
)

type TermDeposit struct {
}

type Validator interface {
	validate() error
}

func (td TermDeposit) ValidateArguments(startAmount float64, interestRate float64, term int, interestPaymentArg string) error {
	validators := []Validator{}
	errorMessages := []error{}

	validators = append(validators,
		PositiveValueValidator{"startingBalance", startAmount},
		PositiveValueValidator{"termYears", float64(term)},
		PositiveValueValidator{"interestRate", interestRate},
		InterestPeriodValidator{interestPaymentArg},
	)

	for _, validator := range validators {
		err := validator.validate()
		if err != nil {
			errorMessages = append(errorMessages, err)
		}
	}

	return errors.Join(errorMessages...)
}

