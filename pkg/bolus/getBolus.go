package bolus

import (
	"fmt"
	"math"

	as "github.com/Starc151/dia/pkg/assistant"
	ydb "github.com/Starc151/dia/pkg/ydb"
)

const (
	lowerGlucose float32 = 7.0
	upperGlucose float32 = 9.0
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
	bolus := float32(0.0)
	corectGlucose := math.Abs(float64(glucose - idealGlucose))
	corrctXe := corectGlucose / (sensitivity * carbohydrate)
	
	m := map[string]float32{
		"glucose"	: glucose,
		"xe"		: xe,
		"bolus"		: bolus,
	}

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
