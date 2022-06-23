package main

import (
	"fmt"

	"github.com/buoyantair/learning-golang/hello"
)

func main()  {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}