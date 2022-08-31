package gophysics

//Gay-Lussac first law  (Volume variation), check for more information https://en.wikipedia.org/wiki/Gay-Lussac%27s_law
func LawGayLussacVolume(volume, celsiusDeg_Temperature float64) float64 {
	return volume * (1 + (1/273)*(celsiusDeg_Temperature))
}

//Gay-Lussac second law (Pressure Variation), check for more information https://en.wikipedia.org/wiki/Gay-Lussac%27s_law
func LawGayLussacPressure(pressure, celsiusDeg_Temperature float64) float64 {
	return pressure * (1 + (1/273)*(celsiusDeg_Temperature))
}

//Net Heat Energy Transfer, check for more information https://en.wikipedia.org/wiki/Rate_of_heat_flow
func NetHeatEnergyTransfer(thermalConductivityConstant, area, kelvinDeg_HeatVariation, time, width float64) float64 {
	return -1 * (thermalConductivityConstant * area * kelvinDeg_HeatVariation * time) / (width)
}

//Heat Flux, check for more information https://en.wikipedia.org/wiki/Thermal_conductivity
func HeatFlux(thermalConductivityConstant, startingTemperatureKelvin, finalTemperatureKelvin, distance float64) float64 {
	return -1 * thermalConductivityConstant * (finalTemperatureKelvin - startingTemperatureKelvin) / distance
}

//Joule Heating, check for more information https://en.wikipedia.org/wiki/Joule_heating
func JouleHeating(resistance, currentIntensity float64) (power float64) {
	return resistance * (currentIntensity * currentIntensity)
}
