package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is: ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The count is: ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	fmt.Println("Log from stringer wrapper", s.String())
}

func main() {
	b := book{
		title: "DSA",
	}

	logInfo(b)

	var c count = 10

	logInfo(c)
}
