package models

import (
	"testing"
	"time"
)

type MockWorkSimulatorWithDuration struct {
	DurationRecorded time.Duration
}

func (mws *MockWorkSimulatorWithDuration) SimulateWork(duration time.Duration) {
	mws.DurationRecorded = duration
}

func TestGrinder_Grind(t *testing.T) {
	tests := []struct {
		name             string
		gramsPerSecond   int
		beanWeightGrams  int
		expectedDuration time.Duration
	}{
		{
			name:             "Standard grind",
			gramsPerSecond:   10,
			beanWeightGrams:  20,
			expectedDuration: 2 * time.Second, // beanWeightGrams / gramsPerSecond
		},
		{
			name:             "Fast grind",
			gramsPerSecond:   20,
			beanWeightGrams:  20,
			expectedDuration: 1 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grinder := NewGrinder(tt.gramsPerSecond)
			mockSimulator := MockWorkSimulatorWithDuration{}

			beans := Beans{weightGrams: tt.beanWeightGrams}
			grinder.Grind(beans, &mockSimulator)

			if grinder.busy {
				t.Errorf("Grinder should not be busy after grinding")
			}

			if mockSimulator.DurationRecorded != tt.expectedDuration {
				t.Errorf("Expected work to be simulated for %v, got %v", tt.expectedDuration, mockSimulator.DurationRecorded)
			}
		})
	}
}
