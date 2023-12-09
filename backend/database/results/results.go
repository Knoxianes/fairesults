package results

import (
	"Knoxiaes/fairesults/helpers"
	"math"
	"time"
)



type Result struct{
	ResultID int
	UserID int
	CompetitionName string
	Category int
	NumberOfCompetitors int
	Place int
	CompetitionRank float64
	Date int64
	MassCoefficinet float64
	Medal int
	Record int
	Points float64
}

func (result *Result) CalculatePoints() {
	result.MassCoefficinet = calculateMassCoefficient(result.NumberOfCompetitors)	
	pointsReduction := calculateReduction(result.Date)
	firstPartOfFormula := float64(100.0/float64(result.NumberOfCompetitors))
	secondPartOfFormula := float64(result.NumberOfCompetitors + 1 - result.Place)
	result.Points = firstPartOfFormula * secondPartOfFormula * result.MassCoefficinet * result.CompetitionRank * pointsReduction
	result.Points = math.Round(result.Points*100)/100
	
}

func calculateMassCoefficient(numberOfCompetitors int) float64{
	if numberOfCompetitors <= 10{
		return 0.5
	}
	if numberOfCompetitors <= 30{
		return 0.6
	}
	if numberOfCompetitors <= 50{
		return 0.7
	}
	if numberOfCompetitors <= 70{
		return 0.8
	}
	if numberOfCompetitors <= 90{
		return 0.9
	}
	if numberOfCompetitors <= 110{
		return 1.0
	}
	if numberOfCompetitors <= 130{
		return 1.1
	}
	return 1.2
}

func calculateReduction(date int64) float64{
	resultSeason := time.Unix(date,0).Year()
	if resultSeason == helpers.CurrentSeason{
		return 1.0	
	}
	if resultSeason == helpers.CurrentSeason - 1{
		return 2/3
	}
	if resultSeason == helpers.CurrentSeason - 2{
		return 1/4
	}
	return 0
}
