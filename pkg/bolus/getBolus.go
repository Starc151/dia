package bolus

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	ydb "github.com/Starc151/dia/pkg/ydb"
)

const (
	lowerGlucose float64 = 7.0
	upperGlucose float64 = 9.0
	idealGlucose = (lowerGlucose + upperGlucose) / 2
)
func GetBolus(glucoseStr, xeStr string) string {
	sensitivity, carbohydrate := getCoefficients()

	indicators := make(map[string]float64)
	toFloat := func (in string) float64 {
				in = strings.ReplaceAll(in, ",", ".")
				out, _ := strconv.ParseFloat(in, 64)
				return out
			}
	indicators["glucose"] = toFloat(glucoseStr)
	indicators["xe"] = toFloat(xeStr)
	indicators["bolus"] = 0.0
	glucoseNill := false
	if indicators["glucose"] == 0.0 {
		indicators["glucose"] = idealGlucose
		glucoseNill = !glucoseNill
	}
	insert := func ()  {
		if glucoseNill {
			indicators["glucose"] = 0.0
		}
		ydb.Insert(indicators)
	}
	defer insert()
	corectGlucose := math.Abs(indicators["glucose"] - idealGlucose)
	corrctXe := corectGlucose / (sensitivity * carbohydrate)
	
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
			indicators["bolus"] = math.Round(indicators["bolus"]*10)/10
		}
		if indicators["xe"] != 0.0{
			indicators["bolus"] += carbohydrate * indicators["xe"]
		}
		return fmt.Sprintf("Болюс: %.1f", indicators["bolus"])
	}
}
