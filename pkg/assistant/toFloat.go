package assistant

import (
	"strconv"
	"strings"
)

func ToFloat(in string) float64 {
	in = strings.ReplaceAll(in, ",", ".")
	out, _ := strconv.ParseFloat(in, 64)
	return out
}
