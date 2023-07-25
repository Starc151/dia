package bolus

import "time"

// Инсулин на коррекцию уровня глюкозы = (текущий - идеальный) / Чувствительность к инсулину
// Доза инсулина на еду = Углеводный коэффициент * Количество ХЕ
func getCoefficients() (float64, float64){
	loc, _ := time.LoadLocation("Europe/Moscow")
    time.Local = loc
	sensitivityCoeff  := 4.0 // Чувствительность к инсулину
	carbohydrateCoeff := 1.25 // Углеводный коэффициент (ед / 1хе)

	if "10:31" <= time.Now().Format("15:04") && time.Now().Format("15:04") <= "13:00"{
		sensitivityCoeff  = 3.5
		carbohydrateCoeff = 0.8
	} else if "13:01" <= time.Now().Format("15:04") && time.Now().Format("15:04") <= "18:00"{
		sensitivityCoeff  = 3.5
		carbohydrateCoeff = 1
	} else if "18:01" <= time.Now().Format("15:04") && time.Now().Format("15:04") <= "23:59"{
		sensitivityCoeff  = 3.0
		carbohydrateCoeff = 1
	}
	return sensitivityCoeff, carbohydrateCoeff
}
