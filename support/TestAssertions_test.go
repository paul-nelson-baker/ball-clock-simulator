package support

import (
	"testing"
)

func testSlicesEqual(t *testing.T, actual, expected []int) {
	if (actual == nil) != (expected == nil) {
		t.FailNow()
	}
	if len(actual) != len(expected) {
		t.FailNow()
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.FailNow()
		}
	}
}
