package gophysics

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

//Power is the work/time ratio
func Power(work, time float64) float64 {
	return work / time
}

//Potential Elastic Energy, check for more information https://en.wikipedia.org/wiki/Elastic_energy
func PotentialElasticEner(elasticConstant, distance float64) float64 {
	return (1 / 2) * elasticConstant * (distance * distance)
}
