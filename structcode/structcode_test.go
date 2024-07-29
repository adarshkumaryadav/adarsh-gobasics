package structcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	t.Run("perimeter", func(t *testing.T) {
		var want float32 = 60.0
		got := perimeter(Rectangle{10.0, 20.0})
		if want != got {
			t.Errorf("was expecting %f got %f", want, got)
			t.Fail()
		}
	})

	t.Run("area", func(t *testing.T) {
		var want float32 = 10.0
		got := area(Rectangle{2, 5.0})
		if !assert.Equal(t, want, got, "want ..%f, got %f", want, got) {
			t.Fail()
		}
	})

	t.Run("circle area", func(t *testing.T) {
		var want float32 = 2 * 3.14 * 3.0
		c := Circle{3.0}
		got := c.area()
		if !assert.Equal(t, want, got, "want %f got %f", want, got) {
			t.Fail()
		}
	})

}
