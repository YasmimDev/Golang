package main

import "fmt"

type Stringer interface {
	String() string
}

type Article struct {
	Title  string
	Author string
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("the %q book was written by %s.", b.Title, b.Author)
}

func (a Article) String() string {
	return fmt.Sprintf("the %q article was written by %s.", a.Title, a.Author)
}

func main() {
	a := Article{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	Print(a)

	b := Book{
		Title:  "All about Go",
		Author: "Jenny bolphin",
		Pages:  25,
	}
	Print(b)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}
