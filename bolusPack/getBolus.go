package bolusPack

import (
	"fmt"
	"math"

	assist "github.com/Starc151/dia/assistantPack"
)

const (
	lowerGlucose float64 = 7.0
	upperGlucose float64 = 9.0
	// Идеальный уровень глюкозы
	idealGlucose = (lowerGlucose + upperGlucose) / 2
)
func GetBolus(glucoseStr, xeStr string) string {
	sensitivityCoeff, carbohydrateCoeff := getCoefficients()
	glucose := assist.ToFloat(glucoseStr)
	xe := assist.ToFloat(xeStr)
	
	if glucose == 0.0 {
		glucose = idealGlucose
	}
	bolus := 0.0
	corectGlucose := math.Abs(glucose - idealGlucose)
	corrctXe := corectGlucose / (sensitivityCoeff * carbohydrateCoeff)
	
	if glucose <= lowerGlucose - 1.0 {
		if xe != 0.0 {
			xe -= corrctXe
			bolus += carbohydrateCoeff * xe
			return fmt.Sprintf("Рекомендуемый болюс: %.1f", bolus)
		} else {
			return fmt.Sprintf("Рекомендуется перекусить на: %.1f XE", corrctXe)
		}
	} else {
		if glucose > idealGlucose + 0.5 {
			bolus = corectGlucose / sensitivityCoeff
		}
		if xe != 0.0{
			bolus += carbohydrateCoeff * xe
		}
		return fmt.Sprintf("Рекомендуемый болюс: %.1f", bolus)
	}
}
