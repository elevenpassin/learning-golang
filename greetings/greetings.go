package greetings

import "fmt"

func Hello(name string) {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
}