package boost

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Div is for divide
func Div(a, b interface{}) float64 {
	var numerator float64
	switch value := a.(type) {
	case int:
		numerator = float64(value)
	case int8:
		numerator = float64(value)
	case int16:
		numerator = float64(value)
	case int32:
		numerator = float64(value)
	case int64:
		numerator = float64(value)
	case uint8:
		numerator = float64(value)
	case uint16:
		numerator = float64(value)
	case uint32:
		numerator = float64(value)
	case uint64:
		numerator = float64(value)
	case float32:
		numerator = float64(value)
	case float64:
		numerator = value
	default:
		numerator = 0
	}

	var denominator float64
	switch value := b.(type) {
	case int:
		denominator = float64(value)
	case int8:
		denominator = float64(value)
	case int16:
		denominator = float64(value)
	case int32:
		denominator = float64(value)
	case int64:
		denominator = float64(value)
	case uint8:
		denominator = float64(value)
	case uint16:
		denominator = float64(value)
	case uint32:
		denominator = float64(value)
	case uint64:
		denominator = float64(value)
	case float32:
		denominator = float64(value)
	case float64:
		denominator = value
	default:
		denominator = 0
	}

	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func SeriousDiv(a, b interface{}) float64 {
	var numerator float64
	switch value := a.(type) {
	case int:
		numerator = float64(value)
	case int8:
		numerator = float64(value)
	case int16:
		numerator = float64(value)
	case int32:
		numerator = float64(value)
	case int64:
		numerator = float64(value)
	case uint8:
		numerator = float64(value)
	case uint16:
		numerator = float64(value)
	case uint32:
		numerator = float64(value)
	case uint64:
		numerator = float64(value)
	case float32:
		numerator = float64(value)
	case float64:
		numerator = value
	default:
		numerator = 0
	}

	var denominator float64
	switch value := b.(type) {
	case int:
		denominator = float64(value)
	case int8:
		denominator = float64(value)
	case int16:
		denominator = float64(value)
	case int32:
		denominator = float64(value)
	case int64:
		denominator = float64(value)
	case uint8:
		denominator = float64(value)
	case uint16:
		denominator = float64(value)
	case uint32:
		denominator = float64(value)
	case uint64:
		denominator = float64(value)
	case float32:
		denominator = float64(value)
	case float64:
		denominator = value
	default:
		denominator = 0
	}

	if denominator == 0 {
		panic(errors.New("denominator is 0"))
	}
	return numerator / denominator
}

// Precision set target's precision
func Precision(target float64, precision int64) float64 {
	fmtStr := "%." + strconv.FormatInt(precision, 10) + "f"
	result, err := strconv.ParseFloat(fmt.Sprintf(fmtStr, target), 64)
	if err != nil {
		panic(err)
	}
	return result
}

// Round converts float64 to float64 with n decimal places
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

// Floor converts float64 to float64 with n decimal places
func Floor(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc(f*n10) / n10
}

// GetStringMod convert string to mod
func GetStringMod(s string, mod uint) uint {
	var sum uint
	for _, b := range []byte(s) {
		sum += uint(b)
	}
	return sum % mod
}

// FloatEqual guarantees n of effective figure
func FloatEqual(f1, f2 float64, n int) bool {
	min := math.Pow10(-1 * n)
	for math.Abs(f1) > 1 {
		f1 /= 10.0
		f2 /= 10.0
	}
	if f1 > f2 {
		return math.Dim(f1, f2) < min
	} else {
		return math.Dim(f2, f1) < min
	}
}

// FloatEqual2 guarantees diff of two number smaller equal than the defined small number
func FloatEqual2(f1, f2 float64, min float64) bool {
	if f1 > f2 {
		return math.Dim(f1, f2) < min
	} else {
		return math.Dim(f2, f1) < min
	}
}
