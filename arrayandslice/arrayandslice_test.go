package arrayandslice

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("Array sum", func(t *testing.T) {
		array := [5]int{1, 2, 3, 4, 5}
		want := 15
		got := Sum(array)
		assert.Equal(t, want, got)
	})

	t.Run("slices sum", func(t *testing.T) {
		got := sumAll([]int{1, 2, 3}, []int{4, 5})
		want := []int{6, 9}
		// slices can only be compared to nil
		// if got != want {
		if !assert.Equal(t, got, want) {
			t.Errorf("test fail, want %d, got %d", want, got)
		}
		fmt.Println("I am here...")
		if !reflect.DeepEqual(got, want) {
			t.Error("failed here too")
			t.Fail()
		}
	})
}
