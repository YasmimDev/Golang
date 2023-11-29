package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReturnValue(t *testing.T) {

	t.Run("Retorno de valores maiores que o limite atribuido", func(t *testing.T) {

		values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		limit := 3

		got := ReturnValue(values, limit)
		want := []int{4, 5, 6, 7, 8}

		if reflect.DeepEqual(got, want) {
			fmt.Println("Returned vectors", got, want)
		} else {
			t.Error("Returned not vectors", got, want)
		}

	})
}
