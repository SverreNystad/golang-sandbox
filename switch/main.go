package main

import (
	"fmt"
	"runtime"
	"time"
)

func whatOs() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func howManyDaysToSaturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// This is an examople of a switch statement with no condition.
// Switch true it is very good for many if else statements.
func greetMe() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Go-od morning!")
	case t.Hour() < 17:
		fmt.Println("Go-od afternoon.")
	default:
		fmt.Println("Go-od evening")
	}
}

func main() {
	// Defer is used to ensure that a function call is performed later in a program’s execution,
	defer fmt.Println("Bye!")
	whatOs()
	howManyDaysToSaturday()
	greetMe()

	fmt.Println("Look at how defers work")
	defer whatOs()
	defer howManyDaysToSaturday()
	defer greetMe()
	// Defers are put on a stack and executed in LIFO order.}
}
