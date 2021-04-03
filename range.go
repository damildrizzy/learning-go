package main
import "fmt"

var pow = [] int {1,2,4,8,16,32,64,128}
var pow2 = []int{2,4,6,8}

func main() {
	for i,v := range pow{
		fmt.Printf("2**%d = %d\n", i,v)
	}

	for _,v := range pow2{
		fmt.Println(v/2)
	}

	cities := map[string]int{
		"New York": 833,
		"LA": 500,
		"Chicago": 200,
	}

	for key, value := range cities {
		fmt.Println(key,value)
	}
}
