package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return float32(3.213)
	case 0 <= balance && balance < 1000:
		return float32(0.5)
	case 1000 <= balance && balance < 5000:
		return float32(1.621)
	default:
		return float32(2.475)
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	interestPerc := InterestRate(balance)
	return balance * float64(interestPerc)/100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	interest := Interest(balance)
	return balance + interest
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var yearsToBalance int
	if targetBalance <= balance {
		return yearsToBalance
	} else {
		for i:=0; i>=0; i++ {
			yearsToBalance += 1
			balance = AnnualBalanceUpdate(balance)
			if balance >= targetBalance {
				break
			}
		}
	}
	return yearsToBalance
}
