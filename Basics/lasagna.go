package main

import "fmt"

func main() {
	fmt.Println(ElapsedTime(1, RemainingOvenTime(10)))
}

/*
Define the OvenTime constant with how many minutes the lasagna should be in the oven. 
According to the cooking book, the expected oven time in minutes is 40:
*/
const OvenTime = 40

/*
Define the RemainingOvenTime() function that takes the actual minutes the lasagna has been in the oven as a parameter and returns how many minutes the lasagna still has to remain in the oven, based on the expected oven time in minutes from the previous task.
*/
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

/*
Define the PreparationTime function that takes the number of layers you added to the lasagna as a parameter and returns how many minutes you spent preparing the lasagna, assuming each layer takes you 2 minutes to prepare.
*/
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

/*
Define the ElapsedTime function that takes two parameters: 
- the first parameter is the number of layers you added to the lasagna, 
- and the second parameter is the number of minutes the lasagna has been in the oven. 

The function should return how many minutes in total you've worked on cooking the lasagna, which is the sum of the preparation time in minutes, and the time in minutes the lasagna has spent in the oven at the moment.
*/
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	prepTime := PreparationTime(numberOfLayers)

	total := prepTime + actualMinutesInOven

	return total
}