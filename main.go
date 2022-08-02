package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Println("Fuel remaining:", fuel)
}

func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}
	return fuel
}

func main() {
	fmt.Println(calculateFuel("Venus"))

}
