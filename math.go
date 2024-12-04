package math

import "math"

func Max(x int, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func Power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

func RoundTo(number float64, places int) float64 {
	c := Power(10, float64(places))
	return math.Round(number/c) * c
}

func RoundHundreds(number float64) float64 {
	return RoundTo(number, 2)
}

func RoundTens(number float64) float64 {
	return RoundTo(number, 1)
}
