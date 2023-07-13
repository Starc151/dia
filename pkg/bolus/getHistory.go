package bolus

import (
	"fmt"
	"strings"

	ydb "github.com/Starc151/dia/pkg/ydb"
)

func GetHistory() string {
	resYdb := ydb.Select()
	if len(resYdb) == 0{
		return "Результатов нет"
	}

	day1 := []string{}
	day2 := []string{}
	day1 = append(day1, resYdb[0].Date + "\n")
	lastResDay1 := 0
	for k, v := range resYdb{
		if v.Date == resYdb[0].Date{
			if k < 2 {
				day1 = append(day1, fmt.Sprint(v.Time, ", Gl:", v.Glucose, ", Xe:", v.Xe, ", Bl:", v.Bolus, "\n"))
				lastResDay1 = k
			} else if k == 2{
				day1 = append(day1, "...\n")
				lastResDay1 = k
			} else {
				lastResDay1 = k
			}
		}
	}

	day2 = append(day2, resYdb[lastResDay1+1].Date + "\n")
	for k, v := range resYdb{
		if v.Date == resYdb[lastResDay1+1].Date{
			if k < (lastResDay1 + 1 + 2)  {
				day2 = append(day2, fmt.Sprint(v.Time, ", Gl:", v.Glucose, ", Xe:", v.Xe, ", Bl:", v.Bolus, "\n"))
			} else {
				day2 = append(day2, "...\n")
				break
			}
		}
	}
	history := fmt.Sprint(day1, "\n", day2)
	history = strings.ReplaceAll(history, "[", "")
	history = strings.ReplaceAll(history, "]", "")
	return history
}