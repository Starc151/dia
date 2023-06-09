package bolus

import (
	"fmt"
	"math"

	as "github.com/Starc151/dia/pkg/assistant"
	ydb "github.com/Starc151/dia/pkg/ydb"
)

const (
	lowerGlucose float64 = 7.0
	upperGlucose float64 = 9.0
	// Идеальный уровень глюкозы
	idealGlucose = (lowerGlucose + upperGlucose) / 2
)
func GetBolus(glucoseStr, xeStr string) string {
	sensitivity, carbohydrate := getCoefficients()

	indicators := make(map[string]float64)
	indicators["glucose"] = as.ToFloat(glucoseStr)
	indicators["xe"] = as.ToFloat(xeStr)
	indicators["bolus"] = 0.0

	if indicators["glucose"] == 0.0 {
		indicators["glucose"] = idealGlucose
	}

	corectGlucose := math.Abs(indicators["glucose"] - idealGlucose)
	corrctXe := corectGlucose / (sensitivity * carbohydrate)
	
	defer ydb.Insert(indicators)
	if indicators["glucose"] <= lowerGlucose - 1.0 {
		if indicators["xe"] != 0.0 {
			indicators["xe"] -= corrctXe
			indicators["bolus"] += carbohydrate * indicators["xe"]
			return fmt.Sprintf("Болюс: %.1f", indicators["bolus"])
		} else {
			return fmt.Sprintf("ГИПА: %.1f XE", corrctXe)
		}
	} else {
		if indicators["glucose"] > idealGlucose + 0.5 {
			indicators["bolus"] = corectGlucose / sensitivity
		}
		if indicators["xe"] != 0.0{
			indicators["bolus"] += carbohydrate * indicators["xe"]
		}
		return fmt.Sprintf("Болюс: %.1f", indicators["bolus"])
	}
}
