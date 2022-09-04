package formulas

import (
	"math"

	"github.com/Gabri432/gophysics/constants"
)

// Defining custom ChargeBody with its own charge, radius, distance from another charge
type ChargeBody struct {
	Charge, Radius, Distance float64
}

func (c ChargeBody) ElectricField() (value float64, measurementUnit string) {
	return ElectricField(c.Charge, c.Radius)
}

func (c ChargeBody) ElectricPotential() (value float64, measurementUnit string) {
	return ElectricPotential(c.Charge, c.Radius)
}

func (c ChargeBody) ElectricPotentDiff(distance2 float64) (value float64, measurementUnit string) {
	return ElectricPotentDiff(c.Charge, c.Distance, distance2)
}

// Coulomb law formula, check for more information https://en.wikipedia.org/wiki/Coulomb%27s_law
func LawCoulomb(charge1, charge2, distance float64) (electroStaticForce float64, measurementUnit string) {
	return constants.COULOMB * (math.Abs(charge1) * math.Abs(charge2)) / (distance * distance), "N"
}

// Gauss theorem, Flux of electric field, check for more information https://en.wikipedia.org/wiki/Gauss%27s_law
func GaussFlux(charge float64) (value float64, measurementUnit string) {
	return charge / constants.VACUUM_PERMITTIVITY, "N*m^2/C"
}

// Electric Field, check for more information https://en.wikipedia.org/wiki/Electric_potential
// or https://phys.libretexts.org/Bookshelves/College_Physics/Book%3A_College_Physics_(OpenStax)/19%3A_Electric_Potential_and_Electric_Field/19.01%3A_Electric_Potential_Energy-_Potential_Difference
func ElectricField(charge, radius float64) (value float64, measurementUnit string) {
	return (charge) / (4 * 3.1416 * constants.VACUUM_PERMITTIVITY * radius * radius), "N/C"
}

// Electric Potential on electric Field formula from a point-like charge, check for more information https://en.wikipedia.org/wiki/Electric_potential
// or https://phys.libretexts.org/Bookshelves/College_Physics/Book%3A_College_Physics_(OpenStax)/19%3A_Electric_Potential_and_Electric_Field/19.01%3A_Electric_Potential_Energy-_Potential_Difference
func ElectricPotential(charge, distance float64) (value float64, measurementUnit string) {
	return (charge) / (4 * 3.1416 * constants.VACUUM_PERMITTIVITY * distance), "V"
}

// Electric Potential Energy Difference (do not confuse it with Voltage), check for more information https://en.wikipedia.org/wiki/Electric_potential_energy
// or https://phys.libretexts.org/Bookshelves/College_Physics/Book%3A_College_Physics_(OpenStax)/19%3A_Electric_Potential_and_Electric_Field/19.01%3A_Electric_Potential_Energy-_Potential_Difference
func ElectricPotentEnergyDiff(charge1, charge2, distance1, distance2 float64) (value float64, measurementUnit string) {
	return ((charge1 * charge2) / (4 * 3.1416 * constants.VACUUM_PERMITTIVITY)) * ((1 / distance1) - (1 / distance2)), "J"
}

// Electric Potential Difference, or Voltage, between two point from the charge position, check for more information https://en.wikipedia.org/wiki/Voltage
// or https://phys.libretexts.org/Bookshelves/College_Physics/Book%3A_College_Physics_(OpenStax)/19%3A_Electric_Potential_and_Electric_Field/19.01%3A_Electric_Potential_Energy-_Potential_Difference
func ElectricPotentDiff(charge, distance1, distance2 float64) (value float64, measurementUnit string) {
	return (charge / (4 * 3.1416 * constants.VACUUM_PERMITTIVITY)) * ((1 / distance1) - (1 / distance2)), "V"
}

// Voltage, or Electric Potential Difference, between two point from the charge position, check for more information https://en.wikipedia.org/wiki/Voltage
// or https://phys.libretexts.org/Bookshelves/College_Physics/Book%3A_College_Physics_(OpenStax)/19%3A_Electric_Potential_and_Electric_Field/19.01%3A_Electric_Potential_Energy-_Potential_Difference
func Voltage(charge, distance1, distance2 float64) (value float64, measurementUnit string) {
	return ElectricPotentDiff(charge, distance1, distance2)
}

// Conductor capacitance, check for more information https://en.wikipedia.org/wiki/Capacitance
func SelfCapacitance(charge, potential float64) (coulombPotential float64, measurementUnit string) {
	return charge / potential, "C/V"
}

// Sphere Conductor capacitance, check for more information https://en.wikipedia.org/wiki/Capacitance
func SelfCapacitanceSphere(radius float64) (coulombPotential float64, measurementUnit string) {
	return 4 * 3.14159 * constants.VACUUM_PERMITTIVITY * radius, "C/V"
}

// Density of energy of Electric Field, check for more information https://en.wikipedia.org/wiki/Electric_potential_energy
func EnergyDens(electricFieldModule float64) (value float64, measurementUnit string) {
	return (1 / 2) * constants.VACUUM_PERMITTIVITY * (electricFieldModule * electricFieldModule), "J/m^3"
}

// First Ohm law, check for more information https://en.wikipedia.org/wiki/Ohm%27s_law
func LawOhm1(resistance, currentIntensity float64) (voltage float64, measurementUnit string) {
	return resistance * currentIntensity, "V"
}

// Second Ohm law, check for more information https://en.wikipedia.org/wiki/Ohm%27s_law
func LawOhm2(resistivity, length, area float64) (electricalResistance float64, measurementUnit string) {
	return resistivity * length / area, "Ohm"
}

// Drift Speed, check for more information https://en.wikipedia.org/wiki/Drift_velocity
func DriftSpeed(intensity, chargeCarrierDensity, area, particleCharge float64) (driftVelocity float64, measurementUnit string) {
	return intensity / (chargeCarrierDensity * area * particleCharge), "m/s"
}

// Energy Density of electric and magnetic fields, check for more information https://en.wikipedia.org/wiki/Energy_density
func EnergyDensity(electricField, magneticField float64) (value float64, measurementUnit string) {
	return (1/2)*constants.DIELECTRIC*(electricField*electricField) + (1/(2*constants.VACUUM_PERMEABILITY))*(magneticField*magneticField), "J/m^3"
}
