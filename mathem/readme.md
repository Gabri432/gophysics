# mathem
![GitHub](https://img.shields.io/github/license/Gabri432/gophysics)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Gabri432/gophysics)
[![Go Reference](https://pkg.go.dev/badge/github.com/Gabri432/gophysics.svg)](https://pkg.go.dev/github.com/Gabri432/gophysics/mathem)

A sub-package of the gophysics library providing some of the most common Mathematical constants and formulas used in Physics.

## List of constants
- GreekPi = 3.1415926535
- Radiant = 57.2958
- E = 2.71828
- SquareRoot_2 = 1.41421
- SquareRoot_3 = 1.73205

## List of formulas
- Radiant to Degrees and viceversa conversions
- Circle Area, Sphere Volume
- Sine Square, Cosine Square
- Powering at, Powering 10 at
- Summatory, Product of Sequence
- Arithmetic Mean, Variance, Standard Deviation

## Example
```go
import (
    "fmt"
    "github.com/Gabri432/gophysics/mathem"
)

func main() {
    fmt.Println(mathem.StandardDeviation(2, 4, 4, 4, 5, 5, 7, 9)) // Using the Standard Deviation formula
}

=== Output ===
2

```