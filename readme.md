# gophysics
A set of some of the most common physics formulas and constants.

## Which formulas are contained?
### Classical mechanics
- Force, Speed, Time, Work, Acceleration, Density, Intensity, Power, Momentum
- Potential Energy, Mechanical Energy, Kinetic Energy

### Gravity
- Potential Gravitational Energy, Universal Gravitational Law, Gravitational Field
- Escape Speed

### Fluids
- Stokes's Law, Hagen-Poiseuille's Law

### ElectroMagnetism
- Ohm's Law, Capacitance, Voltage, Coulomb's Law, Doppler Effect, Energy Density

### Termodinamics
- Gay-Lussac's Law, Net Heat Energy, Heat Flux, Eletric Field Flux, Joule Heating

### Relativity
- Lorentz Factor, Relativistic Time, Relativistic Distance, Relativstic Mass, Relativistic Momentum
- Photoelettric Effect, Drift Speed.


## How to use it
- Download 
```
go get github.com/Gabri432/gophysics
```

- Example of usage
```
package main

import (
    "fmt"
    "github.com/Gabri432/gophysics
)

func main() {
    fmt.Println(gophysics.C)
}

=== Output ===
299792453  

```