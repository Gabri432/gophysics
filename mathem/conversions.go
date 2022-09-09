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

//PowerAt calls the math.Pow method to return base**power
func PowerAt(base, power float64) float64 {
	return math.Pow(base, power)
}

//Power10 calls the math.Pow10 method to return base*10**power
func Power10(base float64, power int) float64 {
	return base * math.Pow10(power)
}

/*
Summatory function will calculate the user defined function from the starting to the ending integer value.

Each output will be added to a total sum that will be returned.

Example:

import "github.com/Gabri432/gophysics/mathem"

func main() {
	var oddNumberFunction = func(value int) float64 {
	                       return float64(2*value+1) //2n+1
                           }
    fmt.Println(mathem.Summatory(1, 3, oddNumberFunction)) // [2*(1)+1] + [2*(2)+1] + [2*(3)+1] = 3 + 5 + 7 = 15
}

*/
func Summatory(startingValue, endingValue int, function func(value int) float64) (sum float64) {
	for i := startingValue; i <= endingValue; i++ {
		sum += function(i)
	}
	return sum
}
