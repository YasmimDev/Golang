package main

import "fmt"

type error interface {
	Error() string
}

type networkProblem struct {
	message string
	code    int
}

func (np networkProblem) Error() string {
	return fmt.Sprintf("network error! message: %s, code: %v", np.message, np.code)
}

func handleErr(err error) {

	fmt.Println(err.Error())

}

func main() {

	np := networkProblem{
		message: "we received a problem",
		code:    404,
	}

	handleErr(np)
}
