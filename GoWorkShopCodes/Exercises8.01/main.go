package main

import
	"github.com/Franlky01/Exercises/package1/shape"



func main() {
	t := shape.Triangle{Base: 15.5,Height: 20.1}
	r := shape.Rectangle{Length: 20,Width: 10}
	s := shape.Square{Side: 10}
	shape.PrintShapeDetails(t,r,s)

}


