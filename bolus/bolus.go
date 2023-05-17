package bolus

import "fmt"

const idealGlucose float64 = 8.0 // Идеальный уровень глюкозы

func Bolus() {
	sensitivityCoeff, carbohydrateCoeff := getCoefficients()
	glucose := getRes("Уровень глюкозы: ")
	xe := getRes("XE: ")
	bolus := (glucose - idealGlucose) / sensitivityCoeff
	if xe != 0.0{
		bolus += carbohydrateCoeff * xe
	}

	fmt.Printf("Рекомендуемый болюс: %.1f", bolus)
}
