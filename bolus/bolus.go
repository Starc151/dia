package bolus

import "fmt"

const (
	lowerGlucose float64 = 7.0
	upperGlucose float64 = 9.0
	idealGlucose = (lowerGlucose + upperGlucose) / 2 // Идеальный уровень глюкозы
)

func Bolus() {
	sensitivityCoeff, carbohydrateCoeff := getCoefficients()
	(idealGlucose - glucose) / (sensitivityCoeff * carbohydrateCoeff)
	glucose := getRes("Уровень глюкозы: ")
	xe := getRes("XE: ")

	if glucose <= lowerGlucose - 1.0 {
		fmt.Println("Рекомендуется перекусить на: ". )
	}

	bolus := (glucose - idealGlucose) / sensitivityCoeff
	if xe != 0.0{
		bolus += carbohydrateCoeff * xe
	}

	fmt.Printf("Рекомендуемый болюс: %.1f", bolus)
}
