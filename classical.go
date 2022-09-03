package gophysics

// Defining custom body with its own mass, speed, volume
type Body struct {
	Mass, Speed, Volume float64
}

//Force is the product between mass and acceleration
func (b Body) Force(acceleration float64) (value float64, measurementUnit string) {
	return Force(b.Mass, acceleration)
}

//Acceleration is the force/mass ratio
func (b Body) Acceleration(force float64) (value float64, measurementUnit string) {
	return Acceleration(force, b.Mass)
}

//Density is the weight/volume ratio
func (b Body) Density() (value float64, measurementUnit string) {
	return Density(b.Mass, b.Volume)
}

//Potential Energy is the product between mass, mass acceleration and height
func (b Body) PotentialEner(acceleration, height float64) (value float64, measurementUnit string) {
	return PotentialEner(b.Mass, acceleration, height)
}

//Kinetic Energy is the product between half the mass and the square speed
func (b Body) KineticEner() (value float64, measurementUnit string) {
	return KineticEner(b.Mass, b.Speed)
}

//Force is the product between mass and acceleration
func Force(mass, acceleration float64) (value float64, measurementUnit string) {
	return mass * acceleration, "N"
}

//Speed is the distance/time ratio
func Speed(distance, time float64) (value float64, measurementUnit string) {
	return distance / time, "m/s"
}

//Time is the distance/speed ratio
func Time(distance, speed float64) (value float64, measurementUnit string) {
	return distance / speed, "s"
}

//Work is the product between force and distance
func Work(force, distance float64) (value float64, measurementUnit string) {
	return force * distance, "J"
}

//Acceleration is the force/mass ratio
func Acceleration(force, mass float64) (value float64, measurementUnit string) {
	return force / mass, "m/s^2"
}

//Density is the weight/volume ratio
func Density(weight, volume float64) (value float64, measurementUnit string) {
	return weight / volume, "kg/m^3"
}

//Intensity is the power/area ratio
func Intensity(power, area float64) (value float64, measurementUnit string) {
	return power / area, "W/m^2"
}

//Potential Energy is the product between mass, mass acceleration and height
func PotentialEner(mass, acceleration, height float64) (value float64, measurementUnit string) {
	return mass * acceleration * height, "J"
}

//Kinetic Energy is the product between half the mass and the square speed
func KineticEner(mass, speed float64) (value float64, measurementUnit string) {
	return (1 / 2) * mass * (speed * speed), "J"
}

//Mechanic Energy is the sum of Potential energy and kinetic energy, check for more information https://en.wikipedia.org/wiki/Mechanical_energy
func MechanicalEner(potential, kinetic float64) (value float64, measurementUnit string) {
	return potential + kinetic, "J"
}

//Momentum is the product between mass and speed
func Momentum(mass, speed float64) (value float64, measurementUnit string) {
	return mass * speed, "kg*m/s"
}

//Power is the work/time ratio
func Power(work, time float64) (value float64, measurementUnit string) {
	return work / time, "W"
}

//Potential Elastic Energy, check for more information https://en.wikipedia.org/wiki/Elastic_energy
func PotentialElasticEner(elasticConstant, distance float64) (value float64, measurementUnit string) {
	return (1 / 2) * elasticConstant * (distance * distance), "J"
}

//Frequency is the speed/distance ratio
func Frequency(speed, distance float64) (value float64, measurementUnit string) {
	return speed / distance, "hertz"
}

//Doppler effect (when listener is getting closer to the sound source), check for more information https://en.wikipedia.org/wiki/Doppler_effect
func DopplerCloser(speed, frequence float64) (frequency float64, measurementUnit string) {
	return (1 + (speed / 340)) * frequence, "hertz"
}

//Doppler effect (when listener is getting far away from sound source), check for more information https://en.wikipedia.org/wiki/Doppler_effect
func DopplerFarer(speed, frequence float64) (frequency float64, measurementUnit string) {
	return (1 - (speed / 340)) * frequence, "hertz"
}

//Angular Frequency is the radiant/time ratio
func AngularFreq(time float64) (value float64, measurementUnit string) {
	return (2 * 3.1416) / time, "rad/s"
}
