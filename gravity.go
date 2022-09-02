package gophysics

import "math"

// Defining a custom planet with its own mass and radius
type PlanetBody struct {
	Mass, Radius float64
}

// Calculates gravity field
func (p PlanetBody) GravField() float64 {
	return GravField(p.Mass, p.Radius)
}

// Calculate escape velocity
func (p PlanetBody) EscapeSpeed() float64 {
	return EscapeSpeed(p.Mass, p.Radius)
}

//Potential Gravitational Energy, check for more information https://en.wikipedia.org/wiki/Gravitational_energy
func PotentialGravEner(mass, height float64) float64 {
	return mass * GRAVITY * height
}

//Universal Gravitational Law (Gravitational Attraction Law), check for more information https://en.wikipedia.org/wiki/Gravity
func GravAttract(mass1, mass2, distance float64) float64 {
	return (G * mass1 * mass2) / (distance * distance)
}

//Gravitational Field of an object, check for more information https://en.wikipedia.org/wiki/Gravity
func GravField(mass, distance float64) float64 {
	return (G * mass) / (distance * distance)
}

//Potential Gravitatonal Energy between two masses, check for more information https://en.wikipedia.org/wiki/Gravitational_energy
func PotentialGravEner2(mass1, mass2, distance float64) float64 {
	return -1 * ((G * mass1 * mass2) / (distance))
}

//Escape Speed, check for more information https://en.wikipedia.org/wiki/Escape_velocity
func EscapeSpeed(mass, radius float64) float64 {
	return math.Pow((2*G*mass)/(radius), 0.5)
}
