package test

import "fmt"

func HelloWorld(s string) string {
	return fmt.Sprintf("%s", s) // nolint:staticcheck
}
