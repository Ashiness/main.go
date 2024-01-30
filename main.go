package main

import (
	"awesomeProject3/Cities"
	"awesomeProject3/Digit"
	newcolor "awesomeProject3/color"
	"awesomeProject3/wordz"
	"fmt"
	"github.com/huandu/xstrings"
)

func main() {

	newcolor.Greet()
	fmt.Println("Hello world")
	for x := 0; x < 5; x++ {
		fmt.Println(wordz.Hello)
		fmt.Println(wordz.Random())
	}

	for y := 0; y < 5; y++ {
		fmt.Println(Cities.Hello)
		fmt.Println(Cities.City())
	}

	for z := 0; z < 5; z++ {
		fmt.Println(Digit.Hello)
		fmt.Println(Digit.Digit())
	}

	var shuff = xstrings.Shuffle(Cities.City())
	fmt.Print(shuff)

}
