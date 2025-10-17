package conversion

import (
	"fmt"
)

const (
	MeterToFoot 	= 3.28084
	MeterToInch		= 39.3701
	MeterToYard 	= 1.09361
	MeterToMile		= 0.000621371

	KilogramToPound	= 2.20462
	KilogramToGram 	= 1000.0
	KilogramToOunce = 35.274

	InvalidUnit	= "invalid_unit"
)

type UnitConversion struct {
	Value	float64
	Type	string
	Error	error
}

func convertTemperatureToCelsius(value float64, fromUnit string) (float64, error){
	switch fromUnit { 
	case "c", "celsius", "cel": 
		return value, nil

	case "f", "fahrenheit", "fah":
		return (value - 32.0) * 5.0 / 9.0, nil

	case "k", "kelvin", "kel": 
		return value - 273.15, nil

	default:
		return 0, fmt.Errorf(InvalidUnit)
	}
}

func convertCelsiusToTarget(celsius float64, toUnit string) (float64, error) {
	switch toUnit {
	case "c", "celsius", "cel":
		return celsius, nil
	case "f", "fahrenheit", "fah":
		// F = C * 9/5 + 32
		return celsius*9.0/5.0 + 32.0, nil
	case "k", "kelvin", "kel":
		// K = C + 273.15
		return celsius + 273.15, nil
	default:
		return 0, fmt.Errorf(InvalidUnit)
	}
}

// HandleTemperatureConversion performs temperature conversion via Celsius as base. (Exported)
func HandleTemperatureConversion(value float64, fromUnit, toUnit string) (float64, error) {
	celsius, err := convertTemperatureToCelsius(value, fromUnit)
	if err != nil && err.Error() == InvalidUnit {
		return 0, fmt.Errorf("invalid temperature unit '%s' or '%s'", fromUnit, toUnit)
	}
	
	targetValue, err := convertCelsiusToTarget(celsius, toUnit)
	if err != nil && err.Error() == InvalidUnit {
		return 0, fmt.Errorf("invalid temperature unit '%s' or '%s'", fromUnit, toUnit)
	}

	return targetValue, nil
}

// convertLengthToMeters converts any length unit to the base (Meter).
func convertLengthToMeters(value float64, fromUnit string) (float64, error) {
	switch fromUnit {
	case "m", "meter", "meters":
		return value, nil
	case "ft", "foot", "feet":
		return value / MeterToFoot, nil
	case "in", "inch", "inches":
		return value / MeterToInch, nil
	case "km", "kilometer", "kilometers":
		return value * 1000, nil // km to m
	case "cm", "centimeter", "centimeters":
		return value / 100, nil // cm to m
	case "mm", "millimeter", "millimeters":
		return value / 1000, nil // mm to m
	case "yd", "yard", "yards":
		return value / MeterToYard, nil
	case "mi", "mile", "miles":
		return value / MeterToMile, nil
	default:
		return 0, fmt.Errorf(InvalidUnit)
	}
}

// convertMetersToTarget converts the base unit (Meter) to the target length unit.
func convertMetersToTarget(meters float64, toUnit string) (float64, error) {
	switch toUnit {
	case "m", "meter", "meters":
		return meters, nil
	case "ft", "foot", "feet":
		return meters * MeterToFoot, nil
	case "in", "inch", "inches":
		return meters * MeterToInch, nil
	case "km", "kilometer", "kilometers":
		return meters / 1000, nil // m to km
	case "cm", "centimeter", "centimeters":
		return meters * 100, nil // m to cm
	case "mm", "millimeter", "millimeters":
		return meters * 1000, nil // m to mm
	case "yd", "yard", "yards":
		return meters * MeterToYard, nil
	case "mi", "mile", "miles":
		return meters * MeterToMile, nil
	default:
		return 0, fmt.Errorf(InvalidUnit)
	}
}

// HandleLengthConversion performs length conversion via Meters as base. (Exported)
func HandleLengthConversion(value float64, fromUnit, toUnit string) (float64, error) {
	meters, err := convertLengthToMeters(value, fromUnit)
	if err != nil && err.Error() == InvalidUnit {
		return 0, fmt.Errorf("invalid length unit '%s' or '%s'", fromUnit, toUnit)
	}
	
	targetValue, err := convertMetersToTarget(meters, toUnit)
	if err != nil && err.Error() == InvalidUnit {
		return 0, fmt.Errorf("invalid length unit '%s' or '%s'", fromUnit, toUnit)
	}

	return targetValue, nil
}


// convertWeightToKilograms converts any weight unit to the base (Kilogram).
func convertWeightToKilograms(value float64, fromUnit string) (float64, error) {
	switch fromUnit {
	case "kg", "kilogram", "kilograms":
		return value, nil
	case "lb", "pound", "pounds":
		return value / KilogramToPound, nil
	case "g", "gram", "grams":
		return value / KilogramToGram, nil
	case "mg", "milligram", "milligrams":
		return value / 1000000, nil // mg to kg
	case "oz", "ounce", "ounces":
		return value / KilogramToOunce, nil
	default:
		return 0, fmt.Errorf(InvalidUnit)
	}
}

// convertKilogramsToTarget converts the base unit (Kilogram) to the target weight unit.
func convertKilogramsToTarget(kilograms float64, toUnit string) (float64, error) {
	switch toUnit {
	case "kg", "kilogram", "kilograms":
		return kilograms, nil
	case "lb", "pound", "pounds":
		return kilograms * KilogramToPound, nil
	case "g", "gram", "grams":
		return kilograms * KilogramToGram, nil
	case "mg", "milligram", "milligrams":
		return kilograms * 1000000, nil // kg to mg
	case "oz", "ounce", "ounces":
		return kilograms * KilogramToOunce, nil
	default:
		return 0, fmt.Errorf(InvalidUnit)
	}
}

// HandleWeightConversion performs weight conversion via Kilograms as base. (Exported)
func HandleWeightConversion(value float64, fromUnit, toUnit string) (float64, error) {
	kilograms, err := convertWeightToKilograms(value, fromUnit)
	if err != nil && err.Error() == InvalidUnit {
		return 0, fmt.Errorf("invalid weight unit '%s' or '%s'", fromUnit, toUnit)
	}
	
	targetValue, err := convertKilogramsToTarget(kilograms, toUnit)
	if err != nil && err.Error() == InvalidUnit {
		return 0, fmt.Errorf("invalid weight unit '%s' or '%s'", fromUnit, toUnit)
	}

	return targetValue, nil
}