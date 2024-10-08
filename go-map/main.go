package main

import "fmt"

func main() {
	// instanciate a map
	// using make function, map is not nil, it is empty
	// with this declaration, we can specified size and cap
	m := make(map[string][]string)
	// declaration like this is the same as above, with a abit of a difference,
	// you can't specified size and cap
	// m := map[string][]string{}

	// declare map without initilization, this will cause error when assign key:value to a map,
	// assignment, map can't be nil
	// var m map[string][]string

	// add item to a map
	m["fruits"] = []string{"apple", "banana", "grape"}
	m["vegs"] = []string{"carbage", "spinich"}

	// get item from a map key
	fmt.Println(m)

	// for get, if map is nil, the program won't panic
	fmt.Println(m["fruits"])
	fmt.Println(m["vegs"])

	// loop within the slice
	for _, fruit := range m["fruits"] {
		fmt.Println(fruit)
	}

	// delete item from a map
	// same with delete, nil, wont trigger panic
	delete(m, "fruits")
	fmt.Println(m)

	// use struct as key in map
	type AggregateData struct {
		Date, SensorType string
	}

	weatherData := map[AggregateData]float64{}

	weatherData[AggregateData{
		Date:       "1992-02-14",
		SensorType: "hum",
	}] = 558.55

	weatherData[AggregateData{
		Date:       "1993-06-28",
		SensorType: "temp",
	}] = 554.2

	fmt.Println(weatherData)
}
