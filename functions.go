package main

import(
	"fmt"
	
)

func add (x, y int )int {
	return x + y
}

//named results
func location(city string) (state, continent string){
	switch city{
	case "Ibadan":
		state, continent = "Oyo", "Africa"
	case "Lekki":
		state, continent = "Lagos", "Nigeria"
	default:
		state, continent = "Remote", "Remote"
	}
	return

}


func main() {
	fmt.Println(add(1,2))
	fmt.Println(location("Ibadan"))
}