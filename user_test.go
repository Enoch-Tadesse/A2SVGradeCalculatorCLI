package main

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {

	dummyData := User{
		Name: "someone",
		Subjects: []Subject{
			Subject{"Calculus", 95},
			Subject{"Philosophy", 88},
			Subject{"Database", 91},
		},
	}
	expected := "91.33"
	average := fmt.Sprintf("%.2f", dummyData.calculateAverage())
	if average != expected {
		t.Fatalf("Expected average %s, Found %s", expected, average)
	}
}
