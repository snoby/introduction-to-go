package sprintf

import (
	"fmt"
	"testing"
)

func TestTrainJourney(t *testing.T) {
	train, number, duration := "Pelham 123", 291, 2.2
	got := fmt.Sprintf(" FIXME ", train, number, duration)
	want := "Pelham 123 number 291 duration 2.2 hours"
	if got != want {
		t.Fatalf("got %q, expected %q", got, want)
	}
}

func TestCities(t *testing.T) {
	cities := []string{"Tampa", "Atlanta", "Dallas"}
	got := fmt.Sprintf("%s", cities)
	want := "Tampa, Atlanta, Dallas"
	if got != want {
		t.Fatalf("got %q, expected %q", got, want)
	}
}
