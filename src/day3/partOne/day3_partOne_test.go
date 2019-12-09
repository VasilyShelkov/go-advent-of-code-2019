package main

import "testing"

func TestIntcodeOptCode1(t *testing.T) {
	wireOneDirections := []string{"R8", "U5", "L5", "D3"}
	wireTwoDirections := []string{"U7", "R6", "D4", "L4"}
	expectedOutput := 6
	result := calculateMinimumDistance([]([]string){wireOneDirections, wireTwoDirections})
	if (!(result == expectedOutput)) {
		t.Errorf(`calculateMinimumDistance is not equal to %v, got %v`, expectedOutput, result)
	}	
}

func TestIntcodeOptCode2(t *testing.T) {
	wireOneDirections := []string{"R75","D30","R83","U83","L12","D49","R71","U7","L72"}
	wireTwoDirections := []string{"U62","R66","U55","R34","D71","R55","D58","R83"}
	expectedOutput := 159
	result := calculateMinimumDistance([]([]string){wireOneDirections, wireTwoDirections})
	if (!(result == expectedOutput)) {
		t.Errorf(`calculateMinimumDistance is not equal to %v, got %v`, expectedOutput, result)
	}	
}

func TestIntcodeOptCodeReplaceHalt(t *testing.T) {
	wireOneDirections := []string{"R98","U47","R26","D63","R33","U87","L62","D20","R33","U53","R51"}
	wireTwoDirections := []string{"U98","R91","D20","R16","D67","R40","U7","R15","U6","R7"}
	expectedOutput := 135
	result := calculateMinimumDistance([]([]string){wireOneDirections, wireTwoDirections})
	if (!(result == expectedOutput)) {
		t.Errorf(`calculateMinimumDistance is not equal to %v, got %v`, expectedOutput, result)
	}	
}
