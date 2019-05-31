package domain

import (
	"fmt"

	"github.com/shopspring/decimal"
)

//LineMonthIncome is a line one month income
type LineMonthIncome struct {
	Domain      `json:",inline" bson:",inline"`
	Line        `json:"line" bson:"line"`
	MonthIncome `json:"income" bson:"income"`
}

//TwoMonthsLineIncome include last two month incomes
type TwoMonthsLineIncome struct {
	Line
	FinalIncome             LineMonthIncome `json:"-"`
	PreviousIncome          LineMonthIncome `json:"-"`
	FinalTotalIncome        decimal.Decimal `json:"final_total_income"`
	PreviousTotalIncome     decimal.Decimal `json:"previous_total_income"`
	FinalAverageIncome      decimal.Decimal `json:"final_average_income"`
	PreviousAverageIncome   decimal.Decimal `json:"previous_average_income"`
	TotalIncomeGrowthRate   decimal.Decimal `json:"total_growth_rate"`
	AverageIncomeGrowthRate decimal.Decimal `json:"average_growth_rate"`
	TotalIncomeRank         int             `json:"totalIncomeRank"`
	CapitaIncomeRank        int             `json:"capitaIncomeRank"`
	Score                   int             `json:"score"`
}

func (a TwoMonthsLineIncome) String() string {
	return fmt.Sprintf("\nLine: %d. Final Total Income: %s. Previous Total Income: %s. Total Income Growth Rate: %s. Average Income Growth Rate: %s. Score: %d", a.Line.No, a.FinalTotalIncome, a.PreviousTotalIncome, a.TotalIncomeGrowthRate, a.AverageIncomeGrowthRate, a.Score)
}

//CMLISortByScore is sortabel array for TwoMonthsLineIncome struct
type CMLISortByScore []TwoMonthsLineIncome

func (a CMLISortByScore) Len() int      { return len(a) }
func (a CMLISortByScore) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a CMLISortByScore) Less(i, j int) bool {

	return a[i].Score < a[j].Score
}

//CMLISortByTotalIncome is sortabel array for TwoMonthsLineIncome struct
type CMLISortByTotalIncome []TwoMonthsLineIncome

func (a CMLISortByTotalIncome) Len() int      { return len(a) }
func (a CMLISortByTotalIncome) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a CMLISortByTotalIncome) Less(i, j int) bool {

	return a[i].GetTotalIncomeGrowthRate().LessThan(a[j].GetTotalIncomeGrowthRate())
}

//CMLISortByAverageIncome is sortabel array for TwoMonthsLineIncome struct
type CMLISortByAverageIncome []TwoMonthsLineIncome

func (a CMLISortByAverageIncome) Len() int      { return len(a) }
func (a CMLISortByAverageIncome) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a CMLISortByAverageIncome) Less(i, j int) bool {

	return a[i].GetAverageIncomeGrowthRate().LessThan(a[j].GetAverageIncomeGrowthRate())
}

//GetAverageIncomeGrowthRate get average person per day income groth rate
func (a *TwoMonthsLineIncome) GetAverageIncomeGrowthRate() decimal.Decimal {
	if a.AverageIncomeGrowthRate.Equal(decimal.Zero) {
		FAI := a.GetFinalAverageIncome()
		PAI := a.GetPreviousAverageIncome()
		a.AverageIncomeGrowthRate = FAI.Sub(PAI).Div(PAI).Mul(decimal.New(1, 2))
	}
	return a.AverageIncomeGrowthRate
}

//GetTotalIncomeGrowthRate get a line total income growth rate between two months
func (a *TwoMonthsLineIncome) GetTotalIncomeGrowthRate() decimal.Decimal {
	if a.TotalIncomeGrowthRate.Equal(decimal.Zero) {
		finalIncome := a.GetFinalTotalIncome()
		previousIncome := a.GetPreviousTotalIncome()
		a.TotalIncomeGrowthRate = finalIncome.Sub(previousIncome).Div(previousIncome).Mul(decimal.New(100, 0))
	}
	return a.TotalIncomeGrowthRate
}

//GetFinalTotalIncome plus final card and cash income
func (a *TwoMonthsLineIncome) GetFinalTotalIncome() decimal.Decimal {
	if a.FinalTotalIncome.Equal(decimal.Zero) {
		a.FinalTotalIncome = decimal.New(a.FinalIncome.Cash+a.FinalIncome.Card, 0)
	}
	return a.FinalTotalIncome
}

//GetPreviousTotalIncome plus previous card and cash income
func (a *TwoMonthsLineIncome) GetPreviousTotalIncome() decimal.Decimal {
	if a.PreviousTotalIncome.Equal(decimal.Zero) {
		a.PreviousTotalIncome = decimal.New(a.PreviousIncome.Cash+a.PreviousIncome.Card, 0)
	}
	return a.PreviousTotalIncome
}

//GetFinalAverageIncome get final average income
func (a *TwoMonthsLineIncome) GetFinalAverageIncome() decimal.Decimal {
	if a.FinalAverageIncome.Equal(decimal.Zero) {
		a.FinalAverageIncome = a.GetFinalTotalIncome().Div(decimal.New(a.FinalIncome.Attendance, 0))
	}
	return a.FinalAverageIncome
}

//GetPreviousAverageIncome get previous average income
func (a *TwoMonthsLineIncome) GetPreviousAverageIncome() decimal.Decimal {
	if a.PreviousAverageIncome.Equal(decimal.Zero) {
		a.PreviousAverageIncome = a.GetPreviousTotalIncome().Div(decimal.New(a.PreviousIncome.Attendance, 0))
	}
	return a.PreviousAverageIncome
}
