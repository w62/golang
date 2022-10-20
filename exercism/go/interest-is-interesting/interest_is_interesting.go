package interest


// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var rate float32

	switch (true) {
	case balance < 0.0:
		rate = 3.213
	case balance >= 0.0 && balance < 1000.00:
		rate = 0.5
	case balance >= 1000.00 && balance < 5000:
		rate = 1.621
	case balance >= 5000.00:
		rate = 2.475
	default:
		rate = 0.0
	}
	return rate
	panic("Please implement the InterestRate function")
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) * 0.01
	panic("Please implement the Interest function")
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
	panic("Please implement the AnnualBalanceUpdate function")
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	start := balance
	
	var year int
	for start <= targetBalance {
		start += Interest(start) 
		year ++
	}

	return year
	panic("Please implement the YearsBeforeDesiredBalance function")
}
