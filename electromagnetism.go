package gophysics

import "math"

//Doppler effect (when listener is getting closer to the sound source), check for more information https://en.wikipedia.org/wiki/Doppler_effect
func DopplerCloser(speed, frequence float64) float64 {
	return (1 + (speed / 340)) * frequence
}

//Doppler effect (when listener is getting far away from sound source), check for more information https://en.wikipedia.org/wiki/Doppler_effect
func DopplerFarer(speed, frequence float64) float64 {
	return (1 - (speed / 340)) * frequence
}

//Coulomb law formula, check for more information https://en.wikipedia.org/wiki/Coulomb%27s_law
func LawCoulomb(charge1, charge2, distance float64) (electroStaticForce float64) {
	return COULOMB * (math.Abs(charge1) * math.Abs(charge2)) / (distance * distance)
}

//Gauss theorem, Flux of electric field, check for more information https://en.wikipedia.org/wiki/Gauss%27s_law
func GaussFlux(charge float64) float64 {
	return charge / VACUUM_PERMITTIVITY
}

//Electric Potential on electric Field formula from a point-like charge, check for more information https://en.wikipedia.org/wiki/Electric_potential
func ElectricPotential(charge, distance float64) float64 {
	return (charge) / (4 * 3.1416 * VACUUM_PERMITTIVITY * distance)
}

//Electric Potential Difference, or Voltage, between two point from the charge position, check for more information https://en.wikipedia.org/wiki/Voltage
func ElectricPotentDiff(charge, distance1, distance2 float64) float64 {
	return (charge / (4 * 3.1416 * VACUUM_PERMITTIVITY)) * ((1 / distance1) - (1 / distance2))
}

//Voltage, or Electric Potential Difference, between two point from the charge position, check for more information https://en.wikipedia.org/wiki/Voltage
func Voltage(charge, distance1, distance2 float64) float64 {
	return ElectricPotentDiff(charge, distance1, distance2)
}

//Conductor capacitance, check for more information https://en.wikipedia.org/wiki/Capacitance
func SelfCapacitance(charge, potential float64) (coulombPotential float64) {
	return charge / potential
}

//Sphere Conductor capacitance, check for more information https://en.wikipedia.org/wiki/Capacitance
func SelfCapacitanceSphere(radius float64) (coulombPotential float64) {
	return 4 * 3.14159 * VACUUM_PERMITTIVITY * radius
}

//Density of energy of ElectroStatic Field, check for more information https://en.wikipedia.org/wiki/Electric_potential_energy
func EnergyDens(electricFieldModule float64) float64 {
	return (1 / 2) * VACUUM_PERMITTIVITY * (electricFieldModule * electricFieldModule)
}

//First Ohm law, check for more information https://en.wikipedia.org/wiki/Ohm%27s_law
func LawOhm1(resistance, currentIntensity float64) (voltage float64) {
	return resistance * currentIntensity
}

//Second Ohm law, check for more information https://en.wikipedia.org/wiki/Ohm%27s_law
func LawOhm2(resistivity, length, area float64) (electricalResistance float64) {
	return resistivity * length / area
}

//Drift Speed, check for more information https://en.wikipedia.org/wiki/Drift_velocity
func DriftSpeed(intensity, chargeCarrierDensity, area, particleCharge float64) (driftVelocity float64) {
	return intensity / (chargeCarrierDensity * area * particleCharge)
}

//Energetic Density Mean, check for more information https://en.wikipedia.org/wiki/Energy_density
func EnergyDensity(electricField, magneticField float64) float64 {
	return (1/2)*DIELECTRIC*(electricField*electricField) + (1/(2*VACUUM_PERMEABILITY))*(magneticField*magneticField)
}
