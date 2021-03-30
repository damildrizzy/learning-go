package main 

import (
	"fmt"
	"time"
)

type Bootcamp struct{
	Lat float64
	Long float64
	Date time.Time
}

type Point struct {
	X, Y int	
}

var(
	p = Point{1, 2}
	q = &Point{1, 2}
	r = Point{X:1}
	s = Point{}
)



func main() {
	fmt.Println(Bootcamp{
		Lat: 34.01284,
		Long: -118.6534,
		Date: time.Now(),

	})
	fmt.Println(p,*q,r,s)

	event := Bootcamp{
		Lat: 33.1,
		Long: 30,
	}
	
	event.Date = time.Now()

	fmt.Println(event.Date, event.Lat, event.Long)
}