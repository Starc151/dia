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
			fmt.Println("Рекомендуется перекусить на: ", corrctXe, "XE")
		}
	}

	// bolus = corectGlucose / sensitivityCoeff
	// if xe != 0.0{
	// 	bolus += carbohydrateCoeff * xe
	// }

	// fmt.Printf("Рекомендуемый болюс: %.1f", bolus)
}
