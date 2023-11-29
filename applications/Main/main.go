package main

func ReturnValue(values []int, limit int) []int {

	out := []int{}

	for i := 0; i < len(values); i++ {

		if values[i] > limit {

    out = append(out, values[i])

		}
	}
	return out

}
