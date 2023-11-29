package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFloatingVectors(t *testing.T) {

	t.Run("Return of Values", func(t *testing.T) {

		Values := []float64{2.0, 7.0, 8.0, 9.0}
		R := []float64{2.0, 7.0, 5.0, 9.0, 5.8}

		got, err := GetLargerSliceElementWise(Values, R)
		want := []float64{2.0, 7.0, 8.0, 9.0}

		if reflect.DeepEqual(got, want) {
			fmt.Println("Returned Elements", got, want)
		} else {
			t.Error("Error on Return", err, got, want)
		}
	})
}
