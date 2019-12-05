package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
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

	seeking := 19690720

	// Here starts the most advanced, optimal and efficient wat of doing things, keep running around until you find the result (needs to refactor this)
	for i := 1; i < 100; i++ {
		noun := i
		mutableInput := make([]int, len(input))

		for j := 1; j < 100; j++ {
			verb := j

			copy(mutableInput, input)
			mutableInput[1] = noun
			mutableInput[2] = verb

			for i := 0; i < len(mutableInput); i += 4 {
				//fmt.Printf("%d %d %d %d", )
				switch mutableInput[i] {
				case 1:
					mutableInput[mutableInput[i+3]] = mutableInput[mutableInput[i+1]] + mutableInput[mutableInput[i+2]]
				case 2:
					mutableInput[mutableInput[i+3]] = mutableInput[mutableInput[i+1]] * mutableInput[mutableInput[i+2]]
				case 99:
					//fmt.Println("HALTING...")
					//fmt.Printf("value at position 0: %d\n", mutableInput[0])

					if seeking == mutableInput[0] {
						fmt.Printf("Combination: noun %d verb %d\n", noun, verb)
						fmt.Printf("Result: %d\n", 100*noun+verb)
					}

					i = len(mutableInput)
					break
				default:
					fmt.Println("INVALID OPCODE")
					i = len(mutableInput)
					break
				}

			}

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
