package test

import "testing"

func TestHelloWorld(t *testing.T) {
	want := "hello world"
	got := HelloWorld("hello world")
	if want != got {
		t.Error("failed")
	}
}

func TestHelloWorld2(t *testing.T) {
	t.Run("First test", func(t *testing.T) {
		want := "aadi"
		got := HelloWorld("aadi")
		assert(got, want, t)
	})
	t.Run("Second test", func(t *testing.T) {
		want := "aadi"
		got := HelloWorld("aadi")
		assert(got, want, t)
	})
}

func assert(got, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf(t.Name(), "failed\n")
	}
}
