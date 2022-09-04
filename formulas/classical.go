package formulas

import (
	"math"

	"github.com/Gabri432/gophysics/constants"
	"github.com/Gabri432/gophysics/mathem"
)

// Defining custom body with its own mass, speed, volume
type Body struct {
	Mass, Speed, Volume, Radius float64
}

// Force is the product between mass and acceleration
func (b Body) Force(acceleration float64) (value float64, measurementUnit string) {
	return Force(b.Mass, acceleration)
}

// Acceleration is the force/mass ratio
func (b Body) Acceleration(force float64) (value float64, measurementUnit string) {
	return Acceleration(force, b.Mass)
}

// Density is the weight/volume ratio
func (b Body) Density() (value float64, measurementUnit string) {
	return Density(b.Mass, b.Volume)
}

// Potential Energy is the product between mass, mass acceleration and height
func (b Body) PotentialEner(acceleration, height float64) (value float64, measurementUnit string) {
	return PotentialEner(b.Mass, acceleration, height)
}

// Kinetic Energy is the product between half the mass and the square speed
func (b Body) KineticEner() (value float64, measurementUnit string) {
	return KineticEner(b.Mass, b.Speed)
}

// Centripetal Force is the ratio between the mass and square speed product and the radius
func (b Body) CentripetalForce() (force float64, measurementUnit string) {
	return CentripetalForce(b.Mass, b.Speed, b.Radius)
}

// Centripetal Acceleration is the ratio between the square speed and the radius
func (b Body) CentripetalAccel() (force float64, measurementUnit string) {
	return CentripetalAccel(b.Speed, b.Radius)
}

// Maximum height of projectile, check for more info https://en.wikipedia.org/wiki/Projectile_motion
func (b Body) ProjectileFlightTime(angleInDeg float64) (value float64, measurementUnit string) {
	return ProjectileFlightTime(b.Speed, angleInDeg)
}

// Horizontal range of projectile, check for more info https://en.wikipedia.org/wiki/Projectile_motion
func (b Body) ProjectileMaxHeight(angleInDeg float64) (value float64, measurementUnit string) {
	return ProjectileMaxHeight(b.Speed, angleInDeg)
}

// Total time of flight of projectile, check for more info https://en.wikipedia.org/wiki/Projectile_motion
func (b Body) ProjectileMaxRange(angleInDeg float64) (value float64, measurementUnit string) {
	return ProjectileMaxRange(b.Speed, angleInDeg)
}

// Force is the product between mass and acceleration
func Force(mass, acceleration float64) (value float64, measurementUnit string) {
	return mass * acceleration, "N"
}

// Speed is the distance/time ratio
func Speed(distance, time float64) (value float64, measurementUnit string) {
	return distance / time, "m/s"
}

// Time is the distance/speed ratio
func Time(distance, speed float64) (value float64, measurementUnit string) {
	return distance / speed, "s"
}

// Work is the product between force and distance
func Work(force, distance float64) (value float64, measurementUnit string) {
	return force * distance, "J"
}

// Acceleration is the force/mass ratio
func Acceleration(force, mass float64) (value float64, measurementUnit string) {
	return force / mass, "m/s^2"
}

// Density is the weight/volume ratio
func Density(weight, volume float64) (value float64, measurementUnit string) {
	return weight / volume, "kg/m^3"
}

// Intensity is the power/area ratio
func Intensity(power, area float64) (value float64, measurementUnit string) {
	return power / area, "W/m^2"
}

// Potential Energy is the product between mass, mass acceleration and height
func PotentialEner(mass, acceleration, height float64) (value float64, measurementUnit string) {
	return mass * acceleration * height, "J"
}

// Kinetic Energy is the product between half the mass and the square speed
func KineticEner(mass, speed float64) (value float64, measurementUnit string) {
	return (1 / 2) * mass * (speed * speed), "J"
}

// Mechanic Energy is the sum of Potential energy and kinetic energy, check for more information https://en.wikipedia.org/wiki/Mechanical_energy
func MechanicalEner(potential, kinetic float64) (value float64, measurementUnit string) {
	return potential + kinetic, "J"
}

// Momentum is the product between mass and speed
func Momentum(mass, speed float64) (value float64, measurementUnit string) {
	return mass * speed, "kg*m/s"
}

// Power is the work/time ratio
func Power(work, time float64) (value float64, measurementUnit string) {
	return work / time, "W"
}

// Potential Elastic Energy, check for more information https://en.wikipedia.org/wiki/Elastic_energy
func PotentialElasticEner(elasticConstant, distance float64) (value float64, measurementUnit string) {
	return (1 / 2) * elasticConstant * (distance * distance), "J"
}

// Frequency is the speed/distance ratio
func Frequency(speed, distance float64) (value float64, measurementUnit string) {
	return speed / distance, "hertz"
}

// Doppler effect (when listener is getting closer to the sound source), check for more information https://en.wikipedia.org/wiki/Doppler_effect
func DopplerCloser(speed, frequence float64) (frequency float64, measurementUnit string) {
	return (1 + (speed / 340)) * frequence, "hertz"
}

// Doppler effect (when listener is getting far away from sound source), check for more information https://en.wikipedia.org/wiki/Doppler_effect
func DopplerFarer(speed, frequence float64) (frequency float64, measurementUnit string) {
	return (1 - (speed / 340)) * frequence, "hertz"
}

// Angular Frequency is the radiant/time ratio
func AngularFreq(time float64) (value float64, measurementUnit string) {
	return (2 * 3.1416) / time, "rad/s"
}

// Normal Force is the product between the mass, gravitational attracion field, angle cosine of the object.
func NormalForce(mass, angleInDeg float64) (force float64, measurementUnit string) {
	return mass * math.Cos(angleInDeg/57.2958), "N"
}

// Centripetal Force is the ratio between the mass and square speed product and the radius
func CentripetalForce(mass, speed, radius float64) (force float64, measurementUnit string) {
	return (mass * speed * speed) / radius, "N"
}

// Centripetal Acceleration is the ratio between the square speed and the radius
func CentripetalAccel(speed, radius float64) (acceleration float64, measurementUnit string) {
	return (speed * speed) / radius, "m/s^2"
}

// Pendulum Period, check for more information https://en.wikipedia.org/wiki/Pendulum
func PendulumPeriod(pendulumLength float64) (time float64, measurementUnit string) {
	return (2 * 3.1416) * math.Sqrt(pendulumLength/constants.G), "s"
}

// Maximum height of projectile, check for more info https://en.wikipedia.org/wiki/Projectile_motion
func ProjectileMaxHeight(initialVelocity, angleInDeg float64) (height float64, measurementUnit string) {
	return (initialVelocity * initialVelocity * mathem.SineSquare(angleInDeg/mathem.Radiant)) / (2 * constants.G), "m"
}

// Horizontal range of projectile, check for more info https://en.wikipedia.org/wiki/Projectile_motion
func ProjectileMaxRange(initialVelocity, angleInDeg float64) (length float64, measurementUnit string) {
	return (initialVelocity * initialVelocity * mathem.SineSquare(angleInDeg/mathem.Radiant)), "m"
}

// Total time of flight of projectile, check for more info https://en.wikipedia.org/wiki/Projectile_motion
func ProjectileFlightTime(initialVelocity, angleInDeg float64) (time float64, measurementUnit string) {
	return 2 * initialVelocity * math.Sin(angleInDeg/mathem.Radiant), "s"
}
