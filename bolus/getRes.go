package bolus

import "fmt"

func getRes(str string) float64{
	res := 0.0
	fmt.Print(str)
	fmt.Scanln(&res)
	return res
}