package term_deposit

import (
	"errors"
	"fmt"
)

type TermDeposit struct {
	StartingBalance float64
	InterestRate    float64
	TermYears       int
	InterestPaid    string
}

type Validator interface {
	validate() error
}

func (td TermDeposit) ValidateAndPrintBalance() error {
	err := td.validateArguments()
	if err != nil {
		return err
	}

	td.printTotalBalance()
	return nil
}

func (td TermDeposit) printTotalBalance() {
	resultingBalance := TotalBalance(td.StartingBalance, td.InterestRate, td.TermYears, td.InterestPaid)
	fmt.Printf("Total balance deposit maturity $%0.0f\n", resultingBalance)
}

func (td TermDeposit) validateArguments() error {
	validators := []Validator{}
	errorMessages := []error{}

	validators = append(validators,
		PositiveValueValidator{"startingBalance", td.StartingBalance},
		PositiveValueValidator{"termYears", float64(td.TermYears)},
		PositiveValueValidator{"interestRate", td.InterestRate},
		InterestPeriodValidator{td.InterestPaid},
	)

	for _, validator := range validators {
		err := validator.validate()
		if err != nil {
			errorMessages = append(errorMessages, err)
		}
	}

	return errors.Join(errorMessages...)
}

