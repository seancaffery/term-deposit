package term_deposit

import "fmt"

type PositiveValueValidator struct {
	name  string
	value float64
}

func (v PositiveValueValidator) validate() error {
	if v.value > 0 {
		return nil
	}
	return fmt.Errorf("%s must be greater than 0, got '%+v'", v.name, v.value)
}

type InterestPeriodValidator struct {
	value string
}

func (v InterestPeriodValidator) validate() error {
	if _, ok := interestPeriodMapping[v.value]; !ok {
		return fmt.Errorf("interestPaid must be one of: monthly, quarterly, annually, maturity. Got '%v'", v.value)
	}
	return nil
}
