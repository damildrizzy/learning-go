package main 
import "fmt"


type Artist struct{
	Name, Genre string
	Songs int
}

func newRelease (a *Artist) int {// * passing an Artist by reference
	a.Songs ++
	return a.Songs
}

func main() {
	me := &Artist{Name: "Zedd", Genre: "EDM", Songs: 42}
	fmt.Printf("%s released their %dth song ", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}