package bolus

import "time"

// Инсулин на коррекцию уровня глюкозы = (текущий - идеальный) / Чувствительность к инсулину
// Доза инсулина на еду = Углеводный коэффициент * Количество ХЕ
func getCoefficients() (float64, float64){
	nowTime := time.Now()
	sensitivityCoeff  := 4.0 // Чувствительность к инсулину
	carbohydrateCoeff := 1.25 // Углеводный коэффициент

	if "10:31" <= nowTime.Format("15:04") && nowTime.Format("15:04") <= "13:00"{
		sensitivityCoeff  = 3.5
		carbohydrateCoeff = 0.8
	} else if "13:01" <= nowTime.Format("15:04") && nowTime.Format("15:04") <= "18:00"{
		sensitivityCoeff  = 3.5
		carbohydrateCoeff = 1
	} else if "18:01" <= nowTime.Format("15:04") && nowTime.Format("15:04") < "00:00"{
		sensitivityCoeff  = 3.0
		carbohydrateCoeff = 1
	}
	return sensitivityCoeff, carbohydrateCoeff
}
