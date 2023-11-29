package main

import (
	"errors"
)

func GetLargerSliceElementWise(Values []float64, R []float64) ([]float64, error) {

	out := []float64{}

	if len(Values) != len(R) {
		err := errors.New("Invalid Size!")
		return nil, err
	}

	for i := 0; i < len(R); i++ {

		if Values[i] > R[i] {

			out = append(out, Values[i])

		} else if Values[i] == R[i] {
			out = append(out, Values[i])

		} else {
			out = append(out, R[i])
		}
	}
	return out, nil
}
