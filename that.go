package test

import "testing"

func That(t *testing.T, actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) {
	if result := assert(actual, expected...); result != "" {
		t.Fatal(result)
	}
}
