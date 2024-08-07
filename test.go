package main

import (
	"fmt"

	"github.com/Gausons/short-url-go/api"
)

func main() {
	f := "apple"
	fmt.Println(f)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	whatAmi := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
		fmt.Print("I am a function")
	}
	whatAmi(true)
	api.Test()
}
