package main

import "fmt"

func calculateIntcode(originalIntcode []int) []int {
	for i := 0;  i < len(originalIntcode); i += 4 {
		optCode := originalIntcode[i]	

		if (optCode == 99) {
			return originalIntcode
		}

		inputOne := originalIntcode[originalIntcode[i + 1]]
		inputTwo := originalIntcode[originalIntcode[i + 2]]
		outputPosition := originalIntcode[i + 3]

		if (optCode == 1) {
			originalIntcode[outputPosition] = inputOne + inputTwo
		}

		if (optCode == 2) {
			originalIntcode[outputPosition] = inputOne * inputTwo
		}
	}

	return originalIntcode
}

func getOutput(desiredOutput int) int {
	originalIntcode := []int{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,5,19,23,2,10,23,27,1,27,5,31,2,9,31,35,1,35,5,39,2,6,39,43,1,43,5,47,2,47,10,51,2,51,6,55,1,5,55,59,2,10,59,63,1,63,6,67,2,67,6,71,1,71,5,75,1,13,75,79,1,6,79,83,2,83,13,87,1,87,6,91,1,10,91,95,1,95,9,99,2,99,13,103,1,103,6,107,2,107,6,111,1,111,2,115,1,115,13,0,99,2,0,14,0}
	noun := 0
	verb := 0
	for i := 0;  i <= 99; i++ {
		noun = i
		for j := 0; j <= 99; j++ {
			verb = j
			attemptedIntcode := make([]int, len(originalIntcode))
			copy(attemptedIntcode, originalIntcode)
			attemptedIntcode[1] = noun
			attemptedIntcode[2] = verb
			// fmt.Printf("noun: %d", noun)
			// fmt.Printf("verb: %d", verb)
			if (calculateIntcode(attemptedIntcode)[0] == desiredOutput) {
				return (100 * noun) + verb
			}
		}
	}

	return (100 * noun) + verb
}

func main() {
	output := getOutput(19690720)
	fmt.Printf("ouput: %d \n", output)
}