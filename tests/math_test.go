package tests

import (
	"github.com/gouef/math"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		x, y     int
		expected int
	}{
		{1, 2, 2},
		{-1, 1, 1},
		{-5, -3, -3},
		{0, 0, 0},
	}

	for _, test := range tests {
		t.Run("Max", func(t *testing.T) {
			result := math.Max(test.x, test.y)
			assert.Equal(t, test.expected, result, "Max(%d, %d) should equal %d", test.x, test.y, test.expected)
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		base, exponent, expected float64
	}{
		{2, 3, 8},
		{5, 2, 25},
		{2, 0, 1},
		{-2, 3, -8},
		{0, 2, 0},
	}

	for _, test := range tests {
		t.Run("Power", func(t *testing.T) {
			result := math.Power(test.base, test.exponent)
			assert.Equal(t, test.expected, result, "Power(%f, %f) should equal %f", test.base, test.exponent, test.expected)
		})
	}
}

func TestRoundTo(t *testing.T) {
	tests := []struct {
		number   float64
		places   int
		expected float64
	}{
		{123.4567, 1, 120},
		{123.4567, 2, 100},
		{-123.4567, 2, -100},
		{123.4567, 3, 0},
		{500.1567, 3, 1000},
		{123.4567, 0, 123},

		{123.4567, -1, 123.5},
		{123.4567, -2, 123.46},
		{-123.4567, -2, -123.46},
		{123.4567, -3, 123.457},
	}

	for _, test := range tests {
		t.Run("RoundTo", func(t *testing.T) {
			result := math.RoundTo(test.number, test.places)
			assert.Equal(t, test.expected, result, "RoundTo(%f, %d) should equal %f", test.number, test.places, test.expected)
		})
	}
}

func TestRoundHundreds(t *testing.T) {
	tests := []struct {
		number, expected float64
	}{
		{123.4567, 100},
		{123.4001, 100},
		{-123.4567, -100},
		{123.4444, 100},
	}

	for _, test := range tests {
		t.Run("RoundHundreds", func(t *testing.T) {
			result := math.RoundHundreds(test.number)
			assert.Equal(t, test.expected, result, "RoundHundreds(%f) should equal %f", test.number, test.expected)
		})
	}
}

func TestRoundTens(t *testing.T) {
	tests := []struct {
		number, expected float64
	}{
		{123.4567, 120},
		{123.4001, 120},
		{-123.4567, -120},
		{123.4444, 120},
	}

	for _, test := range tests {
		t.Run("RoundTens", func(t *testing.T) {
			result := math.RoundTens(test.number)
			assert.Equal(t, test.expected, result, "RoundTens(%f) should equal %f", test.number, test.expected)
		})
	}
}
