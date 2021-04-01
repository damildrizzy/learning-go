package main
import "fmt"

func main() {
	p := []int{2,3,5,7,9}
	fmt.Println(p)

	fmt.Println(p[0:3])

	names := [4]string{
		"sam",
		"dami",
		"ayo",
		"ade",
	}
	fmt.Println(names)

	a := names[0:2]

	fmt.Println(a)

	//making a slice
	cities := make([]string, 3)
	cities[0] = "Ibadan"
	cities[1] = "Lagos"
	cities[2] = "Abuja"
	cities = append(cities, "Abk")
	otherCities := []string{"akure", "ilesha"}
	cities = append(cities, otherCities...)

	fmt.Println(cities)
}