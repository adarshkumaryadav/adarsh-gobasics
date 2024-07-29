package integer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNumber(t *testing.T) {

	assertIn := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Fail()
		}
	}

	t.Run("add first 3 int", func(t *testing.T) {
		want := 3
		got := addNumbers(3)
		assert.Equal(t, want, got)
	})
	t.Run("second test .. should fail", func(t *testing.T) {
		want := 10
		got := addNumbers(5)
		assertIn(t, got, want)
	})

	// check if the below line is reachable or not
	fmt.Println("reached here...")

}
