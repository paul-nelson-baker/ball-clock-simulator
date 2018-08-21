package support

import (
	"testing"
	"strings"
)

func TestGetUserInput(t *testing.T) {
	inputs, _ := GetValidUserInput(strings.NewReader("123 4"))
	testSlicesEqual(t, inputs, []int{123, 4})
}
