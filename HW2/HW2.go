package main

import (
	"fmt"
)

func main() {
	input := "чоаодбодоржэ"
	fmt.Println("String is ", input)
	leters := make(map[rune]int)
	for _, char := range input {
		leters[char]++
	}
	for char, fric := range leters {
		ffric := (float32(fric) / float32(len(input))) * 200
		leter := string(char)
		fmt.Printf("%s -%d %.2f \n", leter, fric, ffric)
	}
}
