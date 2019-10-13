package simulations

import "fmt"

//InterrestRater gives the interrest rate at a given period
type InterrestRater interface {
	InterrestRate(periode int) float64
}

//Pricer gives the price at a given period using another period as reference
type Pricer interface {
	Price(periode int) float64
}

//Renter gives the rent amount at a given period using another period as reference
type Renter interface {
	Rent(periode int) float64
}

//MarketReturnRater gives the return rate at a given period using another period as reference
type MarketReturnRater interface {
	MarketReturnRate(periode int) float64
}

//EconomicalScenario a scenario describing economical parameters evolution
type EconomicalScenario interface {
	InterrestRater
	Pricer
	Renter
	MarketReturnRater
}

//TransferDutiesRater gives the transfer duties amount
type TransferDutiesRater interface {
	TransferDutiesRate() float64
}

//PropertyTaxer gives the property taxe rate
type PropertyTaxer interface {
	PropertyTaxe() float64
}

//LegalContext all the
type LegalContext interface {
	TransferDutiesRate() float64
	PropertyTaxe(capital float64, debt float64) float64
}

func PrintPlayScenario(initialPrice float64, scenario EconomicalScenario) {
	price := initialPrice
	fmt.Printf("Price is now %f\n", price*scenario.Price(0))

	for i := 1; i < 10; i++ {
		fmt.Printf("Price is now %f\n", scenario.Price(i))
	}
}

//PeriodicalPaymenter gives the
type PeriodicalPaymenter interface {
	PeriodicalPayment(period int) float64
}

//PeriodicalPaymenter gives the
// type PeriodicalPaymenter interface {
// 	PeriodicalPayment() float64
// }

func Simulate(economicalScenario EconomicalScenario, legalContext LegalContext, initialPrice PeriodicalPaymenter, toPeriod int) SimulationStepSummary {

}

//SimulationStepSummary summary at a step of simmulation
type SimulationStepSummary struct {
	Debt           float64
	Capital        float64
	InterestsSum   float64
	InsuranceSum   float64
	TaxeSum        float64
	MaintenanceSum float64
}
