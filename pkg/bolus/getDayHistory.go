package bolus

import "strings"

func GetDayHistory(fullHistory []string) []string {
	dayHistory := []string{}

	if len(fullHistory) <= 5 {
		dayHistory = append(dayHistory, strings.Join(fullHistory, "\n"))
	} else {
		dayHistory = append(dayHistory, strings.Join(fullHistory[0:5], "\n"))
		dayHistory = append(dayHistory, strings.Join(fullHistory, "\n"))
	}
	return dayHistory
}