package zipper

import (
	"testing"
)

func TestStringInSliceHasTrue(t *testing.T) {
	needle := "Hey-stack"
	want := true
	var haystack []string
	haystack = append(haystack, "a", "b", "Hey-stack")
	result := stringInSlice(needle, haystack)

	if result != want {
		t.Fatal()
	}
}

// func TestStringInSliceHasFalse(t *testing.T) {
// 	needle := "Hey-not-stack"
// 	want := false
// 	var haystack []string
// 	haystack = append(haystack, "a", "b", "hey", "Hey-stack")
// 	result := stringInSlice(needle, haystack)

// 	if result != want {
// 		t.Fatal()
// 	}
// }
