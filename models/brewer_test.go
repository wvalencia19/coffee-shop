package models

import (
	"testing"
)

func TestBrewer_Brew(t *testing.T) {
	ouncesWaterPerSecond := 2
	beans := Beans{weightGrams: 24} // 24 grams should produce 12 ounces with the given ratio and multiplier
	expectedOunces := 12

	brewer := NewBrewer(ouncesWaterPerSecond)
	mockSimulator := &MockWorkSimulator{}

	coffee := brewer.Brew(beans, mockSimulator)

	if coffee.sizeOunces != expectedOunces {
		t.Errorf("Brew() got = %v, want %v", coffee.sizeOunces, expectedOunces)
	}
}
