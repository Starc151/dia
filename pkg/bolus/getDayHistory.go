package bolus

import "strings"

func GetDayHistory(fullHistory []string) []string {
	dayHistory := []string{}

	if len(fullHistory) <= 3 {
		dayHistory = append(dayHistory, strings.Join(fullHistory, "\n"))
	} else {
		dayHistory = append(dayHistory, strings.Join(fullHistory[0:3], "\n"))
		dayHistory = append(dayHistory, strings.Join(fullHistory, "\n"))
	}
	return dayHistory
}