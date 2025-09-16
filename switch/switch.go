package main

import (
	"fmt"
	"time"
)

func main() {
	i := 4
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("unknown input")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's a weekend")
	default:
		fmt.Println("it's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() > 12:
		fmt.Println("it's after noon")

	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("invalid time")
	}

	whatAmI := func(i interface{}) {
        switch q := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
		case string:
			fmt.Println("i'm a string")
        default:
            fmt.Printf("Don't know type %T\n", q)
        }
    }

	whatAmI(false)
	whatAmI(23)
	whatAmI("hello")
	whatAmI(23.5)

}