package main

import "fmt"

func main() {

	//variable lenght arrays
	empty_map := map[string]string{}

	fmt.Println(empty_map)

	location_map := map[string]string{
		"city":    "Frankfurt",
		"state":   "Hessen",
		"country": "Germany",
	}

	fmt.Println(location_map)

	coordinates_map := map[string]float32{
		"latitude":  56.7,
		"longitude": 45.6,
	}

	fmt.Println(coordinates_map)

	//print specific values
	fmt.Println(coordinates_map["latitude"])
}
