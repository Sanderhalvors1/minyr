package conv

/*
  I denne pakken skal alle konverteringfunksjoner
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius

var Fahrenheit float64
var Celsius float64
var Kelvin float64

func FahrenheitToCelsius(value float64) float64 {
	return (value - 32.0) * 5.0 / 9.0
}

func CelsiusToFahrenheit(value float64) float64 {
	return value*(9.0/5.0) + 32.0
}

func CelsiusToKelvin(value float64) float64 {
	return value + 273.15
}

func KelvinToCelsius(value float64) float64 {
	return value - 273.15
}

func KelvinToFahrenheit(value float64) float64 {
	return (value * 9.0 / 5.0) - 460.0
}

func FahrenheitToKelvin(value float64) float64 {
	return (value + 460.0) * 5.0 / 9.0
}

// De andre konverteringsfunksjonene implementere her
// ...
