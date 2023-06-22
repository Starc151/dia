package assistant

import (
	"strconv"
	"strings"
)

func ToFloat(in string) float32 {
	in = strings.ReplaceAll(in, ",", ".")
	out, _ := strconv.ParseFloat(in, 32)
	return out
}
