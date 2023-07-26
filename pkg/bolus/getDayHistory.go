package bolus

import "strings"

func GetDayHistory(fullHistory []string) []string {
	dayHistory := []string{}

	if len(fullHistory) <= 4 {
		dayHistory = append(dayHistory, strings.Join(fullHistory, "\n"))
	} else {
		dayHistory = append(dayHistory, strings.Join(fullHistory[0:4], "\n"))
		dayHistory = append(dayHistory, strings.Join(fullHistory, "\n"))
	}
	return dayHistory
}