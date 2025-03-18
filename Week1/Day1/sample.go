package main

import "fmt"
import "math"

type geometry interface {
	area() float32
	perim() float32
}

type circle struct {
	radius float32
}

type rect struct {
	width  float32
	length float32
}

func (r rect) area() float32 {
	return r.width * r.length
}

func (r rect) perim() float32 {
	return 2 * (r.width + r.length)
}

// //////
func (c circle) area() float32 {
	return float32(math.Pi * math.Pow(float64(c.radius), 2))
}

func (c circle) perim() float32 {
	return 2 * math.Pi * c.radius
}

func main() {
	var g1 geometry = rect{width: 1.0, length: 2.0}
	var g2 geometry = circle{radius: 1.0}

	fmt.Println(g1.area(), g1.perim())
	fmt.Println(g2.area(), g2.perim())
	//fmt.Println(g1.(circle))
	if c, ok := g1.(circle); ok {
		fmt.Println("g1 is circle Radius is", c.radius)
	} else {
		fmt.Println("g1 not a circle")
	}
	fmt.Println(g2.(circle))

}
