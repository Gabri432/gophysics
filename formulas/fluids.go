// It is a sub-package providing all fluid dynamics related formulas of the gophysics library
package formulas

import "math"

//Hagen-Poiseuille law, check for more information https://en.wikipedia.org/wiki/Hagen%E2%80%93Poiseuille_equation
func LawHagenPoiseuille(fluidViscosity, pipeLength, flowRate, pipeRadius float64) (pressureDifference float64, measurementUnit string) {
	return ((8 * fluidViscosity * pipeLength * flowRate) / (3.14159 * math.Pow(pipeRadius, 4))), "Pascal"
}

//Stokes law, check for more information https://en.wikipedia.org/wiki/Stokes%27_law
func LawStokes(fluidViscosity, radius, speed float64) (force float64, measurementUnit string) {
	return 6 * 3.14159 * fluidViscosity * radius * speed, "N"
}
