package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < float64(0) {
		return 3.213
	} else if balance >= float64(0) && balance < float64(1000) {
		return 0.5
	} else if balance >= float64(1000) && balance < float64(5000) {
		return 1.621
	} else {
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	interestRate := InterestRate(balance)

	return (float64(interestRate) * balance) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	interest := Interest(balance)

	return balance + interest
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0

	for balance < targetBalance {
		interest := Interest(balance)
		balance += interest
		years++
	}

	return years
}
