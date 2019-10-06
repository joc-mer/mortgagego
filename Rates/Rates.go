package rates

import "math"

//AnnualToMonthlyRate Converts an annual rate to a monthly one
func AnnualToMonthlyRate(annualRate float64) float64 {
	return SubRate(annualRate, 12)
}

//SubRate Converts a rate to a rate for smaller periods
func SubRate(longRate float64, divider int) float64 {
	return math.Pow((1.0+longRate), (1.0/float64(divider))) - 1.0
}

//CalcReimbursment Caclulate the necessary reimbursment for a rate, a given number of period and an implicit capital of 1
func CalcReimbursment(rate float64, periods int) float64 {
	if rate == 0.0 {
		return 1 / float64(periods)
	}

	powRate := math.Pow((rate + 1), float64(periods))
	return (rate * powRate) / (powRate - 1)
}

//CalcCapital Caclulate the accumulated capital for a given rate, a given number of period and an implicit reimbursment of 1 for each period
func CalcCapital(rate float64, periods int) float64 {
	if rate == 0.0 {
		return float64(periods)
	}

	powRate := math.Pow((rate + 1), float64(periods))
	return (powRate - 1) / (rate * powRate)
}
