// A package that provides different functions and constants for making some of the most common calculation in Physics.
//
// Gravity, Fluids, Termodinamics, Relativity are the most covered areas.
//
// As so this package does not claim to be the most extensive or the best one, but it is thought as a useful tool for the most common cases.
//
// Formulas were taken from this italian book: Title - "Fisica Volume 1", Authors - ["P. Mazzoldi", "M. Nigro", "C. Voci"].
//
// If you find any issue or any idea to improve the program you can use this link https://github.com/Gabri432/gophysics/issues/new.
//
// Also If you enjoyed using this program consider putting a star in the github repository at  https://github.com/Gabri432/gophysics.
//
// The package is under the MIT License https://github.com/Gabri432/gophysics/blob/master/license.
//
// Thank you!! :)
package gophysics

import "math"

//PowerAt calls the math.Pow method to return base**power
func PowerAt(base, power float64) float64 {
	return math.Pow(base, power)
}

//Power10 calls the math.Pow10 method to return base*10**power
func Power10(base float64, power int) float64 {
	return base * math.Pow10(power)
}

//Force is the product between mass and acceleration
func Force(mass, acceleration float64) float64 {
	return mass * acceleration
}

//Speed is the distance/time ratio
func Speed(distance, time float64) float64 {
	return distance / time
}

//Time is the distance/speed ratio
func Time(distance, speed float64) float64 {
	return distance / speed
}

//Work is the product between force and distance
func Work(force, distance float64) float64 {
	return force * distance
}

//Acceleration is the force/mass ratio
func Acceleration(force, mass float64) float64 {
	return force / mass
}

//Density is the weight/volume ratio
func Density(weight, volume float64) float64 {
	return weight / volume
}

//Intensity is the power/area ratio
func Intensity(power, area float64) float64 {
	return power / area
}

//Potential Energy is the product between mass, mass acceleration and height
func PotentialEner(mass, acceleration, height float64) float64 {
	return mass * acceleration * height
}

//Kinetic Energy is the product between half the mass and the square speed
func KineticEner(mass, speed float64) float64 {
	return (1 / 2) * mass * (speed * speed)
}

//Mechanic Energy formula
func MechanicalEner(potential, kinetic float64) float64 {
	return potential + kinetic
}

//Momentum is the product between
func Momentum(mass, speed float64) float64 {
	return mass * speed
}

//Power is the power/area ratio
func Power(work, time float64) float64 {
	return work / time
}

//Potential Gravitational Energy, check for more information https://en.wikipedia.org/wiki/Gravitational_energy
func PotentialGravEner(mass, height float64) float64 {
	return mass * GRAVITY * height
}

//Potential Elastic Energy, check for more information https://en.wikipedia.org/wiki/Elastic_energy
func PotentialElasticEner(elasticConstant, distance float64) float64 {
	return (1 / 2) * elasticConstant * (distance * distance)
}

//Hagen-Poiseuille law, check for more information https://en.wikipedia.org/wiki/Hagen%E2%80%93Poiseuille_equation
func LawHagenPoiseuille(fluidViscosity, pipeLength, flowRate, pipeRadius float64) (pressureDifference float64) {
	return ((8 * fluidViscosity * pipeLength * flowRate) / (3.14159 * math.Pow(pipeRadius, 4)))
}

//Stokes law, check for more information https://en.wikipedia.org/wiki/Stokes%27_law
func LawStokes(fluidViscosity, radius, speed float64) (force float64) {
	return 6 * 3.14159 * fluidViscosity * radius * speed
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
