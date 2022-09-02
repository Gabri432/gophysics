package gophysics

import "math"

// Defining custom ChargeBody with its own charge, radius, distance from another charge
type ChargeBody struct {
	Charge, Radius, Distance float64
}

func (c ChargeBody) ElectricField() float64 {
	return ElectricField(c.Charge, c.Radius)
}

func (c ChargeBody) ElectricPotential() float64 {
	return ElectricPotential(c.Charge, c.Radius)
}

func (c ChargeBody) ElectricPotentDiff(distance2 float64) float64 {
	return ElectricPotentDiff(c.Charge, c.Distance, distance2)
}

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

//Electric Field, check for more information https://en.wikipedia.org/wiki/Electric_potential
// or https://phys.libretexts.org/Bookshelves/College_Physics/Book%3A_College_Physics_(OpenStax)/19%3A_Electric_Potential_and_Electric_Field/19.01%3A_Electric_Potential_Energy-_Potential_Difference
func ElectricField(charge, radius float64) float64 {
	return (charge) / (4 * 3.1416 * VACUUM_PERMITTIVITY * radius * radius)
}

//Electric Potential on electric Field formula from a point-like charge, check for more information https://en.wikipedia.org/wiki/Electric_potential
// or https://phys.libretexts.org/Bookshelves/College_Physics/Book%3A_College_Physics_(OpenStax)/19%3A_Electric_Potential_and_Electric_Field/19.01%3A_Electric_Potential_Energy-_Potential_Difference
func ElectricPotential(charge, distance float64) float64 {
	return (charge) / (4 * 3.1416 * VACUUM_PERMITTIVITY * distance)
}

//Electric Potential Energy Difference (do not confuse it with Voltage), check for more information https://en.wikipedia.org/wiki/Electric_potential_energy
// or https://phys.libretexts.org/Bookshelves/College_Physics/Book%3A_College_Physics_(OpenStax)/19%3A_Electric_Potential_and_Electric_Field/19.01%3A_Electric_Potential_Energy-_Potential_Difference
func ElectricPotentEnergyDiff(charge1, charge2, distance1, distance2 float64) float64 {
	return ((charge1 * charge2) / (4 * 3.1416 * VACUUM_PERMITTIVITY)) * ((1 / distance1) - (1 / distance2))
}

//Electric Potential Difference, or Voltage, between two point from the charge position, check for more information https://en.wikipedia.org/wiki/Voltage
// or https://phys.libretexts.org/Bookshelves/College_Physics/Book%3A_College_Physics_(OpenStax)/19%3A_Electric_Potential_and_Electric_Field/19.01%3A_Electric_Potential_Energy-_Potential_Difference
func ElectricPotentDiff(charge, distance1, distance2 float64) float64 {
	return (charge / (4 * 3.1416 * VACUUM_PERMITTIVITY)) * ((1 / distance1) - (1 / distance2))
}

//Voltage, or Electric Potential Difference, between two point from the charge position, check for more information https://en.wikipedia.org/wiki/Voltage
// or https://phys.libretexts.org/Bookshelves/College_Physics/Book%3A_College_Physics_(OpenStax)/19%3A_Electric_Potential_and_Electric_Field/19.01%3A_Electric_Potential_Energy-_Potential_Difference
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
