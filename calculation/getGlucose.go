package calculation

import (
	"fmt"
)

const (
	idealGlucose            = 8.0 // Идеальный уровень глюкозы
	sensitivityCoefficient  = 4.0 // Чувствительность к инсулину
	carbohydrateCoefficient = 1.0 // Углеводный коэффициент
)

// Инсулин на коррекцию уровня глюкозы = (текущий - идеальный) / Чувствительность к инсулину
// Доза инсулина на еду = Углеводный коэффициент * Количество ХЕ

func GetGlucose() {
	glucose := 0.0
	xe := 0.0
	fmt.Print("Результат измерения: ")
	fmt.Scanln(&glucose)
	fmt.Print("XE: ")
	fmt.Scanln(&xe)
	bolus := (glucose - idealGlucose) / sensitivityCoefficient
	if xe != 0{
		bolus += carbohydrateCoefficient * xe
	}
	fmt.Printf("%.1f", bolus)
}