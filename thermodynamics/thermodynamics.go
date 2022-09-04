// It is a sub-package providing all thermodynamics related formulas of the gophysics library
package thermodynamics

// Defining custom Gas type with its own volume, pressure and temperature in celsius degrees.
type GasBody struct {
	Volume, Pressure, TemperatureCelsius float64
}

// Gay-Lussac first law  (Volume variation), check for more information https://en.wikipedia.org/wiki/Gay-Lussac%27s_law
func (g GasBody) LawGayLussacVolume() (value float64, measurementUnit string) {
	return LawGayLussacVolume(g.Volume, g.TemperatureCelsius)
}

// Gay-Lussac second law  (Volume variation), check for more information https://en.wikipedia.org/wiki/Gay-Lussac%27s_law
func (g GasBody) LawGayLussacPressure() (value float64, measurementUnit string) {
	return LawGayLussacPressure(g.Pressure, g.TemperatureCelsius)
}

// Gay-Lussac first law  (Volume variation), check for more information https://en.wikipedia.org/wiki/Gay-Lussac%27s_law
func LawGayLussacVolume(volume, celsiusDeg_Temperature float64) (value float64, measurementUnit string) {
	return volume * (1 + (1/273)*(celsiusDeg_Temperature)), "m^3"
}

// Gay-Lussac second law (Pressure Variation), check for more information https://en.wikipedia.org/wiki/Gay-Lussac%27s_law
func LawGayLussacPressure(pressure, celsiusDeg_Temperature float64) (value float64, measurementUnit string) {
	return pressure * (1 + (1/273)*(celsiusDeg_Temperature)), "Pascal"
}

// Net Heat Energy Transfer, check for more information https://en.wikipedia.org/wiki/Rate_of_heat_flow
func NetHeatEnergyTransfer(thermalConductivityConstant, area, kelvinDeg_HeatVariation, time, width float64) float64 {
	return -1 * (thermalConductivityConstant * area * kelvinDeg_HeatVariation * time) / (width)
}

// Heat Flux, check for more information https://en.wikipedia.org/wiki/Thermal_conductivity
func HeatFlux(thermalConductivityConstant, startingTemperatureKelvin, finalTemperatureKelvin, distance float64) (value float64, measurementUnit string) {
	return -1 * thermalConductivityConstant * (finalTemperatureKelvin - startingTemperatureKelvin) / distance, "W"
}

// Joule Heating, or Resistance, check for more information https://en.wikipedia.org/wiki/Joule_heating
func JouleHeating(resistance, currentIntensity float64) (power float64, measurementUnit string) {
	return resistance * (currentIntensity * currentIntensity), "W"
}

// Resistance, or Joule Heating, check for more information https://en.wikipedia.org/wiki/Joule_heating
func Resistance(resistance, currentIntensity float64) (power float64, measurementUnit string) {
	return JouleHeating(resistance, currentIntensity)
}

//Pressure is the normal force/area ratio, check for more information https://en.wikipedia.org/wiki/Pressure
func Pressure(normalForce, area float64) (value float64, measurementUnit string) {
	return normalForce / area, "Pascal"
}
