package bolus

import (
	"fmt"
	"math"
)

const (
	lowerGlucose float64 = 7.0
	upperGlucose float64 = 9.0
	idealGlucose = (lowerGlucose + upperGlucose) / 2 // Идеальный уровень глюкозы
)

func Bolus() {
	sensitivityCoeff, carbohydrateCoeff := getCoefficients()
	glucose := getRes("Уровень глюкозы: ")
	if glucose == 0.0 {
		glucose = idealGlucose
	}
	xe := getRes("XE: ")
	bolus := 0.0
	corectGlucose := math.Abs(glucose - idealGlucose)
	corrctXe := corectGlucose / (sensitivityCoeff * carbohydrateCoeff)
	
	if glucose <= lowerGlucose - 1.0 {
		if xe != 0.0 {
			xe -= corrctXe
			bolus += carbohydrateCoeff * xe
			fmt.Printf("Рекомендуемый болюс: %.1f", bolus)
		} else {
			fmt.Printf("Рекомендуется перекусить на: %.1f XE", corrctXe)
		}
	} else {
		if glucose > idealGlucose + 0.5 {
			bolus = corectGlucose / sensitivityCoeff
		}
		if xe != 0.0{
			bolus += carbohydrateCoeff * xe
		}
		fmt.Printf("Рекомендуемый болюс: %.1f", bolus)
	}
}
