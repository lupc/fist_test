package test

import (
	"fmt"

	"github.com/connerdouglass/go-geo"
)

func RunGeoDist() {
	var p1 = geo.Coordinate{Latitude: 1, Longitude: 1}
	var p2 = geo.Coordinate{Latitude: 1, Longitude: 2}
	fmt.Printf("p1.DistanceTo(p2): %v\n", p1.DistanceTo(p2))

}
