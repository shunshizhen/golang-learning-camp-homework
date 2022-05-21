package main

import (
	"fmt"

	"github.com/shunshizhen/golang-learning-camp-homework/homework-for-errors/biz"
)

func main() {
	err := biz.Find()
	if err != nil {
		fmt.Printf("%+v\n", err)
        return 
	}
	fmt.Print("Done!")
}
