package main

import "fmt"
import "math"
import "container/list"
import "./rates"

func main() {
	monthlyRate := rates.AnnualToMonthlyRate(0.01)
	monthlyInsuranceRate := 0.0036 / 12.0

	capital := rates.CalcNetCapital(monthlyRate, monthlyInsuranceRate, 0.0, 15*12) * 1400.0

	fmt.Printf("Capital %f\n", capital)

	/*scenario := ConstantScenario{
		interrestRate:    0.008,
		priceGrowthRate:  0.02,
		rent:             1000,
		marketReturnRate: 0.01,
	}

	simulations.PrintPlayScenario(200000.0, scenario)*/
}

type ConstantScenario struct {
	interrestRate    float64
	priceGrowthRate  float64
	rent             float64
	marketReturnRate float64
}

func (c ConstantScenario) InterrestRate(period int) float64 {
	return c.interrestRate
}

func (c ConstantScenario) Price(period int) float64 {
	return math.Pow((1.0 + c.priceGrowthRate), float64(period))
}

func (c ConstantScenario) Rent(period int) float64 {
	return c.rent
}

func (c ConstantScenario) MarketReturnRate(period int) float64 {
	return c.marketReturnRate
}

func main1() {
	monthlyRate := rates.AnnualToMonthlyRate(0.01)

	capital := rates.CalcCapital(monthlyRate, 15*12)

	fmt.Printf("Capital %f\n", capital)
	fmt.Printf("Reimbursment %f\n", rates.CalcReimbursment(monthlyRate, 15*12)*capital)
}

func simulateStepsScenario1() {
	printfPattern := "%9.2f"

	capital := 200000.0
	year := 12
	periods := year * 12

	annualRate := 0.007
	monthlyRate := rates.AnnualToMonthlyRate(annualRate)
	annualInsuranceRate := 0.0036
	monthlyInsuranceRate := annualInsuranceRate / 12

	reimbursment := rates.CalcReimbursment(monthlyRate, periods) * (capital * (1 + float64(year)*annualInsuranceRate))

	fmt.Printf("Reimbursment = %f\n", reimbursment)

	steps := computeSteps(capital, reimbursment, monthlyRate, monthlyInsuranceRate)

	fmt.Printf("           |   Amount  |    Debt   |  Capital  | Interests | Int. Sum  | Insurance | Total Cost\n")

	period := 1
	for e := steps.Front(); e != nil; e = e.Next() {
		s := (e.Value).(step)
		fmt.Printf(
			"Period %-3d | "+printfPattern+" | "+printfPattern+" | "+printfPattern+" | "+printfPattern+" | "+printfPattern+" | "+printfPattern+" | "+printfPattern+" \n",
			period, s.Amount, s.Debt, s.Capital, s.Interests, s.InterestsSum, s.InsuranceSum, s.costSum(),
		)
		period++
	}
}

type step struct {
	Amount       float64
	Interests    float64
	InterestsSum float64
	Capital      float64
	Debt         float64
	Insurance    float64
	InsuranceSum float64
}

func (s step) costSum() float64 {
	return s.InterestsSum + s.InsuranceSum
}

func (s step) cost() float64 {
	return s.Interests + s.Insurance
}

func computeSteps(initialDebt float64, reimbursement float64, interestRate float64, insuranceRate float64) *list.List {
	ret := list.New()

	remainingDebt := initialDebt
	interestsSum := 0.0
	insuranceSum := 0.0

	insurance := initialDebt * insuranceRate

	for remainingDebt > 0 {
		interests := remainingDebt * interestRate
		interestsSum += interests
		insuranceSum += insurance

		capital := math.Min(reimbursement-interests-insurance, remainingDebt)
		remainingDebt -= capital

		ret.PushBack(step{
			Amount:       capital + interests + insurance,
			Interests:    interests,
			InterestsSum: interestsSum,
			Capital:      capital,
			Debt:         remainingDebt,
			Insurance:    insurance,
			InsuranceSum: insuranceSum,
		})
	}

	return ret
}
