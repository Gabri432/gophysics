// Package mathem is a sub-package of the gophysics library.
//
// Its porpuse is to make simple conversion of values that are common for Physics calculations.
//
// Some of the most basic operations are radiants to degrees (and viceversa) conversions, area calculations, ecc...
package mathem

import "math"

// It converts radiants to degrees
func RadToDeg(angleInRad float64) (degrees float64, measurementUnit string) {
	return angleInRad * Radiant, "Deg"
}

// It converts degrees to radiants
func DegToRad(angleInDeg float64) (radiants float64, measurementUnit string) {
	return angleInDeg / Radiant, "Rad"
}

// It calculates the circle area
func CircleArea(radius float64) (area float64, measurementUnit string) {
	return Pi * radius * radius, "Squares"
}

// It calculates sphere volume
func SphereVolume(radius float64) (volume float64, measurementUnit string) {
	return (4 / 3) * Pi * (radius * radius * radius), "Cubes"
}

// It calculates sine^2 of the angle in degrees
func SineSquare(angleInDeg float64) float64 {
	return (1 - math.Cos(2*(angleInDeg/Radiant))) / 2
}

// It calculates cosine^2 of the angle in degrees
func CosineSquare(angleInDeg float64) float64 {
	return 1 - SineSquare(angleInDeg)
}
