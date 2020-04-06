package main

import (
	"fmt"
	"math"
)

//Converter Struct
type Converter struct{}

// Centimeter delcaration
type Centimeter float64

// Feet delcaration
type Feet float64
type Minute float64
type Second float64
type Celcius float64
type Fahrenheit float64
type Radian float64
type Degree float64
type Kilogram float64
type Pound float64

// CentimeterToFeet method
func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	return Feet(c / 30.48)
}

// FeetToCentimeter method
func (cvr Converter) FeetToCentimeter(f Feet) Centimeter {
	return Centimeter(f * 30.48)
}

// MinuteToSecond method
func (cvr Converter) MinuteToSecond(m Minute) Second {
	return Second(m * 60)
}

// SecondToMinute method
func (cvr Converter) SecondToMinute(s Second) Minute {
	return Minute(s / 60)
}

// CelciusTofahrenheit method
func (cvr Converter) CelciusTofahrenheit(c Celcius) Fahrenheit {
	result := (c * 1.8) + 32
	return Fahrenheit(result)
}

// FahrenheitToCelcius method
func (cvr Converter) FahrenheitToCelcius(f Fahrenheit) Celcius {
	result := (f - 32) / 1.8
	return Celcius(result)
}

// RadianToDegree method
func (cvr Converter) RadianToDegree(rad Radian) Degree {
	result := rad * (180 / math.Pi)
	return Degree(result)
}

// DegreeToRadian method
func (cvr Converter) DegreeToRadian(deg Degree) Radian {
	result := deg * (math.Pi / 180)
	return Radian(result)
}

// KilogramToPound method
func (cvr Converter) KilogramToPound(kg Kilogram) Pound {
	return Pound(kg * 2.205)
}

// PoundToKilogram method
func (cvr Converter) PoundToKilogram(p Pound) Kilogram {
	return Kilogram(p / 2.205)
}

func main() {
	// Create instance of struct and store it in variable cvr
	cvr := Converter{}
	fmt.Println(cvr.CentimeterToFeet(10))
	fmt.Println(cvr.DegreeToRadian(1))
}
