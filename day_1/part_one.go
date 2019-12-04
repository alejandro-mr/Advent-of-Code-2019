package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	file := LoadInput("./input")
	fuelCounter := 0.00
	for _, i := range strings.Split(strings.Trim(file, "\n"), "\n") {
		mass, _ := strconv.ParseFloat(strings.Trim(i, "\n"), 64)
		fuelCounter += (math.Floor(mass/3) - 2)
	}
	fmt.Printf("Fuel required: %.0f\n", fuelCounter)
}

//LoadInput - function to read file into string
func LoadInput(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	return string(file)
}
