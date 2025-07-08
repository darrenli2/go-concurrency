package main

import (
	"fmt"
)

func main() {
	for msg := range []int{1, 2, 3} {
		done := make(chan bool)
		go run(msg, done)
		go executeMethod(msg, done)
		// <-done
	}

}

func run(msg int, done chan bool) {
	fmt.Println("run", <-done)
}

func executeMethod(r int, done chan bool) {
	arr := []string{"a", "b", "c"}
	for _, val := range arr {
		fmt.Printf("iteration is %d, value is %s\n", r, val)
	}
	done <- true
}
