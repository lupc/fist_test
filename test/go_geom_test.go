package test

import (
	"fmt"
	"math"
	"testing"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/xy"
)

const (
	Degree2Radian = math.Pi / 180
	Radian2Degree = 180 / math.Pi
)

func TestCoordEqual(t *testing.T) {
	var p1 = geom.Coord([]float64{1, 1})
	var p2 = geom.Coord([]float64{1, 2})
	// if !p1.Equal(geom.XY, p2) {
	// 	t.Fail()
	// }
	fmt.Printf("p1.Equal(p2): %v\n", p1.Equal(geom.XY, p2))

}

func TestCoordDist(t *testing.T) {
	var p1 = geom.Coord([]float64{1, 1})
	var p2 = geom.Coord([]float64{1, 5})
	fmt.Printf("xy.Distance(p1, p2): %v\n", xy.Distance(p1, p2))

}

func TestCoordAngle(t *testing.T) {
	var p1 = geom.Coord([]float64{0, 0})
	var p2 = geom.Coord([]float64{-1, 1})
	fmt.Printf("xy.Angle(p1, p2): %v\n", xy.Angle(p1, p2)*Radian2Degree)

}

func TestCoordAngleBetween(t *testing.T) {
	var p1 = geom.Coord([]float64{1, 2})
	var p2 = geom.Coord([]float64{1, 1})
	var p3 = geom.Coord([]float64{2, 1})
	fmt.Printf("xy.AngleBetween(p1, p2): %v\n", xy.AngleBetween(p1, p2, p3)*Radian2Degree)

}
