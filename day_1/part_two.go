package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := LoadInput("./input")
	totalFuel := 0.00
	for _, i := range strings.Split(strings.Trim(input, "\n"), "\n") {
		mass, _ := strconv.ParseFloat(strings.Trim(i, "\n"), 64)
		fuelRequired := CalcFuel(mass)
		fuelMass := fuelRequired
		for fuelMass > 0 {
			fuelMass = CalcFuel(fuelMass)
			if fuelMass <= 0 {
				continue
			}
			totalFuel += fuelMass
		}
		totalFuel += fuelRequired
	}
	fmt.Printf("\nTotal Fuel required: %.0f\n", totalFuel)
}

//LoadInput - function to read file into string
func LoadInput(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	return string(file)
}

//CalcFuel - function to calculate fuel required by mass
func CalcFuel(mass float64) float64 {
	return math.Floor(mass/3) - 2
}
