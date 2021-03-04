package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Printf("Fuel left: %d gallons", fuel)
}

func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 5000000
	case "Mars":
		fuel = 7000000
	default:
		fuel = 0
	}
	return fuel
}

func greetPlanet(planet string) {
	fmt.Printf("Welcome! You're on %v", planet)
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {

	fuel := 10000000

	planetChoice := "Mars"
	var trip1 = flyToPlanet(planetChoice, fuel)
	fmt.Println("")
	fuel = trip1
	fuelGauge(fuel)
	fmt.Println("")
	fmt.Println("---1st Trip Ends---")

	planetChoice = "Venus"
	var trip2 = flyToPlanet(planetChoice, fuel)
	fmt.Println("")
	fuel = trip2
	fuelGauge(fuel)
	fmt.Println("")
	fmt.Println("---2nd Trip Ends---")
}
