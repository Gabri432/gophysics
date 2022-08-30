// A package that provides different functions for making some of the most common calculation in Physics.
//
// Also if you find any issue or any idea to improve the program you can use this link https://github.com/Gabri432/gophysics/issues/new.
//
// If you enjoyed using this program also consider putting a star in the github repository at  https://github.com/Gabri432/gophysics.
//
// Thank you!! :)
package gophysics

import "math"

func Power(base, power float64) float64 {
	return math.Pow(base, power)
}

func Power10(base float64, power int) float64 {
	return base * math.Pow10(power)
}
