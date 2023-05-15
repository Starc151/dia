package calculation

import "fmt"

const idealGlucose float64 = 8.0 // Идеальный уровень глюкозы

func GetGlucose() {
	glucose := 0.0
	xe := 0.0
	sensitivityCoeff, carbohydrateCoeff := getCoefficients()

	fmt.Print("Введите результат измерения: ")
	fmt.Scanln(&glucose)
	fmt.Print("Введите XE: ")
	fmt.Scanln(&xe)
	bolus := (glucose - idealGlucose) / sensitivityCoeff
	if xe != 0.0{
		bolus += carbohydrateCoeff * xe
	}
	fmt.Printf("%.1f", bolus)
}