package utils

import (
	"fmt"
	"strconv"
)

func Float32DecimalTwo(value float32) float32 {
	return float32(Float64DecimalTwo(float64(value)))
}

func Float64DecimalTwo(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
