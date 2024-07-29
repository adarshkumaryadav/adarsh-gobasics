package interfacecode

import (
	"math"
	"testing/dsa/structcode"
)

type shapeArea interface {
	area() float32
	//perimeter() float32
}
type extendedRectangle struct {
	structcode.Rectangle
}
type extendedCircle struct {
	structcode.Circle
}

func (r extendedRectangle) area() float32 {
	return 2 * (r.Height + r.Width)
}

func (c extendedCircle) area() float32 {
	return math.Pi * c.Radius * c.Radius
}
