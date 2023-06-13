package bolus

import (
	"fmt"
	"math"

	as "github.com/Starc151/dia/pkg/assistant"
)

const (
	lowerGlucose float64 = 7.0
	upperGlucose float64 = 9.0
	// Идеальный уровень глюкозы
	idealGlucose = (lowerGlucose + upperGlucose) / 2
)
func GetBolus(glucoseStr, xeStr string) string {
	sensitivity, carbohydrate := getCoefficients()
	glucose := as.ToFloat(glucoseStr)
	xe := as.ToFloat(xeStr)
	
	if glucose == 0.0 {
		glucose = idealGlucose
	}
	bolus := 0.0
	corectGlucose := math.Abs(glucose - idealGlucose)
	corrctXe := corectGlucose / (sensitivity * carbohydrate)
	
	if glucose <= lowerGlucose - 1.0 {
		if xe != 0.0 {
			xe -= corrctXe
			bolus += carbohydrate * xe
			return fmt.Sprintf("Болюс: %.1f", bolus)
		} else {
			return fmt.Sprintf("ГИПА: %.1f XE", corrctXe)
		}
	} else {
		if glucose > idealGlucose + 0.5 {
			bolus = corectGlucose / sensitivity
		}
		if xe != 0.0{
			bolus += carbohydrate * xe
		}
		return fmt.Sprintf("Болюс: %.1f", bolus)
	}
}
