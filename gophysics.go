// A package that provides different functions and constants for making some of the most common calculation in Physics.
//
// Gravity, Fluids, Termodinamics, Relativity are the most covered areas.
//
// As so this package does not claim to be the most extensive or the best one, but it is thought as a useful tool for the most common cases.
//
// Formulas were taken from this italian book: Title - "Fisica Volume 1", Authors - ["P. Mazzoldi", "M. Nigro", "C. Voci"].
//
// If you find any issue or any idea to improve the program you can use this link https://github.com/Gabri432/gophysics/issues/new.
//
// Also If you enjoyed using this program consider putting a star in the github repository at  https://github.com/Gabri432/gophysics.
//
// The package is under the MIT License https://github.com/Gabri432/gophysics/blob/master/license.
//
// Thank you!! :)
package gophysics

import "math"

//PowerAt calls the math.Pow method to return base**power
func PowerAt(base, power float64) float64 {
	return math.Pow(base, power)
}

//Power10 calls the math.Pow10 method to return base*10**power
func Power10(base float64, power int) float64 {
	return base * math.Pow10(power)
}
