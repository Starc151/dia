package bolus

import (
	"fmt"

	ydb "github.com/Starc151/dia/pkg/ydb"
)

func GetFullHistory() [][]string {
	resYdb := ydb.Select()
	fullHistory := [][]string{}
	if len(resYdb) == 0{
		return fullHistory
	}
	lastResDay := 0
	day := []string{resYdb[lastResDay].Date}

	for k, v := range resYdb{
		if v.Date == resYdb[lastResDay].Date {
			day = append(day, fmt.Sprint("|", v.Time, "| Gl:", v.Glucose, ", Xe:", v.Xe, ", Bl:", v.Bolus))
		} else {
			lastResDay = k
			fullHistory = append(fullHistory, day)
			day = nil
			day = append(day, resYdb[lastResDay].Date)
			day = append(day, fmt.Sprint("|", v.Time, "| Gl:", v.Glucose, ", Xe:", v.Xe, ", Bl:", v.Bolus))
		}
	}
	fullHistory = append(fullHistory, day)

	return fullHistory
}