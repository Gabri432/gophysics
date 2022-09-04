// It is a sub-package providing all gravitational related formulas of the gophysics library
package formulas

import (
	"math"

	"github.com/Gabri432/gophysics/constants"
)

// Defining a custom planet with its own mass and radius
type PlanetBody struct {
	Mass, Radius float64
}

// Calculates gravity field
func (p PlanetBody) GravField() (value float64, measurementUnit string) {
	return GravField(p.Mass, p.Radius)
}

// Calculate escape velocity
func (p PlanetBody) EscapeSpeed() (value float64, measurementUnit string) {
	return EscapeSpeed(p.Mass, p.Radius)
}

// Potential Gravitational Energy, check for more information https://en.wikipedia.org/wiki/Gravitational_energy
func PotentialGravEner(mass, height float64) (value float64, measurementUnit string) {
	return mass * constants.GRAVITY * height, "J"
}

// Universal Gravitational Law (Gravitational Attraction Law), check for more information https://en.wikipedia.org/wiki/Gravity
func GravAttract(mass1, mass2, distance float64) (value float64, measurementUnit string) {
	return (constants.G * mass1 * mass2) / (distance * distance), "m/s^2"
}

// Gravitational Field of an object, check for more information https://en.wikipedia.org/wiki/Gravity
func GravField(mass, distance float64) (value float64, measurementUnit string) {
	return (constants.G * mass) / (distance * distance), "m/s^2"
}

// Potential Gravitatonal Energy between two masses, check for more information https://en.wikipedia.org/wiki/Gravitational_energy
func PotentialGravEner2(mass1, mass2, distance float64) (value float64, measurementUnit string) {
	return -1 * ((constants.G * mass1 * mass2) / (distance)), "J"
}

// Escape Speed, check for more information https://en.wikipedia.org/wiki/Escape_velocity
func EscapeSpeed(mass, radius float64) (value float64, measurementUnit string) {
	return math.Pow((2*constants.G*mass)/(radius), 0.5), "m/s"
}
