# gophysics
![GitHub](https://img.shields.io/github/license/Gabri432/gophysics)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Gabri432/gophysics)
[![Go Reference](https://pkg.go.dev/badge/github.com/Gabri432/gophysics.svg)](https://pkg.go.dev/github.com/Gabri432/gophysics)

A set of some of the most common physics formulas and constants.

## Features
- It provides Physics formulas and constants for calculation porpuses.
- It allows to create some objects to simplify some calculations (see example).
- It also provides some common Math formulas in Physics.

## Which formulas are contained?
### Classical mechanics
- Force, Speed, Time, Work, Acceleration, Density, Intensity, Power, Momentum
- Potential Energy, Mechanical Energy, Kinetic Energy, Frequency, Doppler Effect
- Centripetal Acceleration, Centripetal Force, Pendulum Period, Projectile motion equations

### Gravity
- Potential Gravitational Energy, Universal Gravitational Law, Gravitational Field
- Escape Speed

### Fluids
- Stokes's Law, Hagen-Poiseuille's Law

### ElectroMagnetism
- Ohm's Law, Capacitance, Voltage, Coulomb's Law, Energy Density
- Electric Field, Eletric Potential Energy Difference.

### Termodynamics
- Gay-Lussac's Law, Net Heat Energy, Heat Flux, Eletric Field Flux, Joule Heating

### Relativity
- Lorentz Factor, Relativistic Time, Relativistic Distance, Relativstic Mass, Relativistic Momentum
- Photoelettric Effect, Drift Speed.


## How to use it
- Download the 'formulas' package to get the Physics formulas:
```
go get -u github.com/Gabri432/gophysics/formulas
```

- Download the 'constants' package to get the Physics constants:
```
go get -u github.com/Gabri432/gophysics/constants
```

- Download the 'mathem' package to get the Mathematical functions and constants of the library
```
go get -u github.com/Gabri432/gophysics/mathem
```


- Example of usage
```go
package main

import (
    "fmt"
    "github.com/Gabri432/gophysics/constants"
    "github.com/Gabri432/gophysics/formulas"
)

func main() {
    fmt.Println(constants.C)   // Write a constant
    fmt.Println(formulas.Force(3, 4))   // Call a function
    myPlanet := formulas.PlanetBody{Mass: constants.EARTH_MASS, Radius: constants.EARTH_RADIUS} // Create a custom object
	fmt.Println(myPlanet.EscapeSpeed())  // Call object methods to ease some calculations
}

=== Output ===
299792453  
12 N       <<< When calling functions their output is the value and the measurement unit
11183.719071923773 m/s

```

- Using mathematical functions
```go
import (
    "fmt"
    "github.com/Gabri432/gophysics/mathem"
)

func main() {
    fmt.Println(mathem.Pi) // Writing a constant
}

=== Output ===
3.1415926535

```

## Project Structure
### Folders

#### gophysics (main)
- `gophysics.go`, doesn't provide formulas, but other functions to simplify some actions.
- `license`, `readme.md`, `readme.it.md`, `CHANGELOG.txt`.

#### formulas
- `classical.go`, where all the Classical physics formulas are located.
- `fluids.go`, where all the Fluid formulas are located.
- `gravity.go`, where all the Gravity formulas are located.
- `thermodynamics.go`, where all the Thermodynamics formulas are located.
- `electromagnetism.go`, where all the Electromagnetism formulas are located.
- `relativity.go`, where all the Relativity formulas are located.

#### constants
- `constants.go`, where all the constants are located.

#### mathem 
- `constants.go`, where all the mathematical constants are located.
- `conversions.go`, where all the main functions to make mathematical calculations are located.

## Contributing to this project
- If you want to add a feature or making a fix check out this page explaining how to do it: [Contributing to gophysics](https://github.com/Gabri432/gophysics/blob/master/.github/CONTRIBUTING.md)

## Notes
- Formulas were taken from this italian book: Title - "Fisica Volume 1", Authors - ["Paolo Mazzoldi", "Massimo Nigro", "Cesare Voci"].