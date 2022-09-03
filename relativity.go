package gophysics

import "math"

//Relativistic time (Time dilatation), check for more information https://en.wikipedia.org/wiki/Time_dilation
func RelativTime(travelerTime, speed float64) (movingObserverTime float64, measurementUnit string) {
	return travelerTime / (math.Pow(1-((speed*speed)/(C*C)), 0.5)), "s"
}

//Lorentz factor, check for more information https://en.wikipedia.org/wiki/Lorentz_factor
func LorentzFactor(speed float64) (lorentzFactor float64) {
	return 1 / (math.Pow(1-((speed*speed)/(C*C)), 0.5))
}

//Relativistic distance (Compression of distances) in meters, check for more information https://en.wikipedia.org/wiki/Proper_length
func RelativDist(nonTravelerDistance, speed float64) (contractedDistance float64, measurementUnit string) {
	return (math.Pow(1-((speed*speed)/(C*C)), 0.5)) * nonTravelerDistance, "m"
}

//Relativistic Mass (Increment of Mass), check for more information https://en.wikipedia.org/wiki/Mass_in_special_relativity
func RelativMass(travelerMass, speed float64) (relativisticMass float64, measurementUnit string) {
	return travelerMass / (math.Pow(1-((speed*speed)/(C*C)), 0.5)), "kg"
}

//Relativistic Momentum (Increment of Momentum), check for more information https://en.wikipedia.org/wiki/Energy%E2%80%93momentum_relation
func RelativMomentum(travelerMomentum, speed float64) (relativisticMomentum float64, measurementUnit string) {
	return (travelerMomentum * speed) / (math.Pow(1-((speed*speed)/(C*C)), 0.5)), "kg*m/s"
}

//Photoelettric effect, check for more information https://en.wikipedia.org/wiki/Photoelectric_effect
func PhotoElettricEffect(frequence float64) (energy float64, measurementUnit string) {
	return PLANCK * frequence, "J"
}
