package test

import (
	"fmt"

	"github.com/golang/geo/s2"
)

func RunS2Dist() {
	var p1 = s2.Point{}
	p1.X = 1
	p1.Y = 2
	p1.Z = 0

	var p2 = s2.Point{}
	p2.X = 3
	p2.Y = 2
	p2.Z = 0

	fmt.Printf("p1.Distance(p2): %v\n", p1.Distance(p2))
	fmt.Printf("p1.Angle(p2): %v\n", p1.Angle(p2.Vector))

}
