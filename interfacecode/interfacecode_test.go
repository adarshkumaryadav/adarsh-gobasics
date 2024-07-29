package interfacecode

import (
	"math"
	"testing"
	"testing/dsa/structcode"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	helper := func(t *testing.T, want float32, shape shapeArea) {
		t.Helper()
		got := shape.area()
		if !assert.Equal(t, want, got) {
			t.Fail()
		}

	}
	t.Run("rectangle area", func(t *testing.T) {
		r := extendedRectangle{
			structcode.Rectangle{
				Width:  10.0,
				Height: 20.0},
		}
		helper(t, 60.0, r)
	})
	t.Run("circle area", func(t *testing.T) {
		c := extendedCircle{
			structcode.Circle{
				Radius: 1,
			},
		}
		var want float32 = math.Pi * 1 * 1
		helper(t, want, c)
	})
}
