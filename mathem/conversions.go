// Package mathem is a sub-package of the gophysics library.
//
// Its porpuse is to make the most common Math calculations used for Physics.
//
// Some of the most basic operations are radiants to degrees (and viceversa) conversions, area calculations, summatory, ecc...
//
// Here is the documentation Here is the documentation https://pkg.go.dev/github.com/Gabri432/gophysics/mathem
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

Each output will be added to the following one and then store into the resulting sum that will be returned.

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

/*
ProductOfSequence function will calculate the user defined function from the starting to the ending integer value.

Each output will be multiplied to the following one and then store into the resulting product that will be returned.

Example:

import "github.com/Gabri432/gophysics/mathem"

func main() {
	var squareFunction = func(value int) float64 {
	                       return float64(value*value) //n*n
                           }
    fmt.Println(mathem.ProductOfSequence(1, 3, squareFunction)) // [1*1] * [2*2] * [3*3] = 1 * 4 * 9 = 36
}
*/
func ProductOfSequence(startingValue, endingValue int, function func(value int) float64) (product float64) {
	product = 1
	for i := startingValue; i <= endingValue; i++ {
		product *= function(i)
	}
	return product
}

/*
Arithmetic mean return the average value of the len(values) user inserted numbers.

Example:

func main() {
	mathem.ArithemticMean(1,2,3,4,6) // (1+2+3+4+6)/5 => 16/5 == 3.2
}
*/
func ArithmeticMean(values ...float64) float64 {
	sum := 0.0
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}
	return sum / float64(len(values))
}

/*
Variance is the ratio between the square deviations of each value from the mean and the number of values.

Example:

func main() {
	fmt.Println(mathem.Variance(2, 4, 4, 4, 5, 5, 7, 9)) // returns 4

	// STEP 1: Arithmetic mean = 5;
	// STEP 2: Sum of the square deviations =
	// = [(2-5)^2 + (4-5)^2 + (4-5)^2 + (4-5)^2 + (5-5)^2 + (5-5)^2 + (7-5)^2 + (9-5)^2] = 40;
	// STEP 3: Mean of those values = [(9 + 1 + 1 + 1 + 0 + 0 + 4 + 16) / 8] == 4
}
*/
func Variance(values ...float64) float64 {
	arithMean := ArithmeticMean(values...)
	var squareDifference = func(index int) float64 {
		return (values[index] - arithMean) * (values[index] - arithMean)
	}
	return Summatory(0, len(values)-1, squareDifference) / float64(len(values))
}

/*
StandardDeviation is the square root of the Variance.

Example:

func main() {
	fmt.Println(mathem.StandardDeviation(2, 4, 4, 4, 5, 5, 7, 9)) // returns 2

	// STEP 1: Arithmetic mean = 5;
	// STEP 2: Sum of the square deviations =
	// = [(2-5)^2 + (4-5)^2 + (4-5)^2 + (4-5)^2 + (5-5)^2 + (5-5)^2 + (7-5)^2 + (9-5)^2] = 40;
	// STEP 3: Mean of those values = [(9 + 1 + 1 + 1 + 0 + 0 + 4 + 16) / 8] == 4;
	// STEP 4: Square Root of Variance = 2;
}
*/
func StandardDeviation(values ...float64) float64 {
	return math.Sqrt(Variance(values...))
}
