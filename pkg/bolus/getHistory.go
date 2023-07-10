package bolus

import (
	"fmt"
	"strconv"

	ydb "github.com/Starc151/dia/pkg/ydb"
)

func GetHistory() string {
	resultYdb := ydb.Select()
	if len(resultYdb) == 0{
		return "Результатов нет"
	}
	history := fmt.Sprint(resultYdb[0].Date.Format("02 Jan 06"), "\n")
	history += strconv.Itoa(len(resultYdb))
	// history += fmt.Sprint(resultYdb[0].Date.Format("15:04"), ", Gl:", resultYdb[0].Glucose, ", Xe:", resultYdb[0].Xe, ", Bl:", resultYdb[0].Bolus, "\n")
	// if resultYdb[1].Date.Format("02/Jan") != resultYdb[0].Date.Format("02/Jan"){
	// 	history += fmt.Sprint(resultYdb[1].Date.Format("15:04"), ", Gl:", resultYdb[1].Glucose, ", Xe:", resultYdb[1].Xe, ", BL:", resultYdb[1].Bolus, "\n")
	// }
	// history += "...\n \n"
	// for i := 2; true; i++ {
	// 	if resultYdb[i-1].Date.Format("02/Jan") == resultYdb[i].Date.Format("02/Jan") {
	// 		history += fmt.Sprint(resultYdb[i].Date.Format("02/Jan"), "\n")
	// 		history += fmt.Sprint(resultYdb[i].Date.Format("15:04"), resultYdb[i].Glucose, resultYdb[i].Xe, resultYdb[i].Bolus, "\n")
	// 		history += fmt.Sprint(resultYdb[i+1].Date.Format("15:04"), resultYdb[i+1].Glucose, resultYdb[i+1].Xe, resultYdb[i+1].Bolus, "\n")
	// 		history += "..."
	// 		break
	// 	}
	// }
	return history
}