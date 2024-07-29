package mapcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	myMap := map[string]string{"Test": "this is just a test", "What": "what do you mean?"}

	got, err := Search(myMap, "Test")

	if got != "this is just a test" && assert.Nil(t, err) {
		t.Fail()
	}

}
