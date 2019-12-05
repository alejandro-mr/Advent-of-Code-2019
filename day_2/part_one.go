package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("part one")
	stringInput := strings.Split(strings.Trim(LoadInput("./input"), "\n"), ",")
	input := make([]int, len(stringInput))
	for k, i := range stringInput {
		v, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
		} else {
			input[k] = v
		}
	}

	input[1] = 12
	input[2] = 2

	for i := 0; i < len(input); i += 4 {
		switch input[i] {
		case 1:
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
		case 2:
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
		case 99:
			fmt.Println("HALTING...")
			fmt.Printf("value at position 0: %d\n", input[0])
			i = len(input)
			break
		default:
			fmt.Println("INVALID OPCODE")
			i = len(input)
			break
		}

	}

}

//LoadInput - function to read file into string
func LoadInput(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	return string(file)
}
