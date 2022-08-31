// A package that provides different functions for making some of the most common calculation in Physics.
//
// As so this package does not claim to be the most extensive or the best one, but it is thought as a useful tool for the most common cases.
//
// If you find any issue or any idea to improve the program you can use this link https://github.com/Gabri432/gophysics/issues/new.
//
// Also If you enjoyed using this program consider putting a star in the github repository at  https://github.com/Gabri432/gophysics.
//
// The package is under the MIT License https://github.com/Gabri432/gophysics/blob/master/license
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

//Potential Energy is the product between
func PotentialEner(mass, acceleration, height float64) float64 {
	return mass * acceleration * height
}

//Kinetic Energy is the product between
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

//Potential Elastic Energy, check for more information
func PotentialElasticEner(elasticConstant, distance float64) float64 {
	return (1 / 2) * elasticConstant * (distance * distance)
}

//Hagen-Poiseuille law, check for more information https://en.wikipedia.org/wiki/Hagen%E2%80%93Poiseuille_equation
func LawHagenPoiseuille(tubeRadius, pressureGradient, fluidViscosity, tubeLength float64) float64 {
	return ((3.14159 * math.Pow(tubeRadius, 4)) / (8 * fluidViscosity * tubeLength)) * pressureGradient
}

//Stokes law, check for more information https://en.wikipedia.org/wiki/Stokes%27_law
func LawStokes(fluidViscosity, radius, speed float64) float64 {
	return 6 * 3.14159 * fluidViscosity * radius * speed
}

//Universal Gravitational Law (Gravitational Attraction Law), check for more information https://en.wikipedia.org/wiki/Gravity
func GravAttract(mass1, mass2, distance float64) float64 {
	return (G * mass1 * mass2) / (distance * distance)
}

//Gravitational Field of an object
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

//Gay-Lussac first law  (Volume variation), check for more information https://en.wikipedia.org/wiki/Gay-Lussac%27s_law
func LawGayLussacVolume(volume, celsiusDeg_Temperature float64) float64 {
	return volume * (1 + (1/273)*(celsiusDeg_Temperature))
}

//Gay-Lussac second law (Pressure Variation), check for more information https://en.wikipedia.org/wiki/Gay-Lussac%27s_law
func LawGayLussacPressure(pressure, celsiusDeg_Temperature float64) float64 {
	return pressure * (1 + (1/273)*(celsiusDeg_Temperature))
}

//Rate of Heat Flow, check for more information https://en.wikipedia.org/wiki/Rate_of_heat_flow
func HeatFLowRate(termicConducibility, area, kelvinDeg_HeatVariation, time, width float64) float64 {
	return -1 * (termicConducibility * area * kelvinDeg_HeatVariation * time) / (width)
}

//Heat Flux, check for more information https://en.wikipedia.org/wiki/Thermal_conductivity
func HeatFlux(thermalConductivityConstant, startingTemperatureKelvin, finalTemperatureKelvin, distance float64) float64 {
	return -1 * thermalConductivityConstant * (finalTemperatureKelvin - startingTemperatureKelvin) / distance
}

//Doppler effect (when listener is getting closer to the sound source), check for more information https://en.wikipedia.org/wiki/Doppler_effect
func DopplerCloser(speed, frequence float64) float64 {
	return (1 + (speed / 340)) * frequence
}

//Doppler effect (when listener is getting far away from sound source), check for more information https://en.wikipedia.org/wiki/Doppler_effect
func DopplerFarer(speed, frequence float64) float64 {
	return (1 - (speed / 340)) * frequence
}
