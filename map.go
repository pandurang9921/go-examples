package main
import "fmt"
func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 2
	fmt.Println("map:",m)
	v := m["k1"]
	fmt.Println(v)
}