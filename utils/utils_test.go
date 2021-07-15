package utils

import (
	"strings"
	"testing"
)

func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}

func TestRates(t *testing.T) {
	rate, _ := Rates("GHS", "NGN")

	if rate != 69.31 {
		t.Errorf("failed expected 69.31")
	}
}

func TestRatesFull(t *testing.T) {
	tests := []struct {
		currency string
		target   string
		expected float64
	}{
		{"GHS", "NGN", 69.31},
		{"GHS", "KSH", 18.24},

		{"NGN", "GHS", 0.014},
		{"NGN", "KSH", 0.26},

		{"KSH", "NGN", 3.80},
		{"KSH", "GHS", 0.055},
	}

	for i := range tests {
		rate, _ := Rates(tests[i].currency, tests[i].target)
		if rate != tests[i].expected {
			t.Errorf("failed expected {%f} but received {%f}", tests[i].expected, rate)
		}
	}
}

func TestRatesFail(t *testing.T) {
	_, err := Rates("GHS", "ECO")
	if !ErrorContains(err, "Invalid Currency") {
		t.Errorf("unexpected error: %v", err)
	}
}
