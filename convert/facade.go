package convert

import (
	"strconv"
)

func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func StrToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func FloatToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func StrToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func BoolToStr(b bool) string {
	return strconv.FormatBool(b)
}
