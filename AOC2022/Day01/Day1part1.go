package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	// Read Input File
	input,_ := os.Open("input.txt")
	defer input.Close() // delays execution time of function and waits for nearby functions return
	sc := bufio.NewScanner(input) // Read input of big size ( buffer size )

	// Search for maximum sum of calories
	maxCalories := 0
	currentElfCalories := 0

	for sc.Scan(){
		snack, err := strconv.Atoi(sc.Text()) // converts string type to int and stores in snack
		currentElfCalories += snack

		if err != nil{
			if currentElfCalories > maxCalories{
				maxCalories = currentElfCalories
			}
			// start with new elf
			currentElfCalories=0
		}

	}
	fmt.Println(maxCalories)
}


